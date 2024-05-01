// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IForwarder} from "./interfaces/IForwarder.sol";
import {ConfirmedOwner} from "../shared/access/ConfirmedOwner.sol";
import {TypeAndVersionInterface} from "../interfaces/TypeAndVersionInterface.sol";
import {Utils} from "./libraries/Utils.sol";

// solhint-disable gas-custom-errors, no-unused-vars
contract KeystoneForwarder is IForwarder, ConfirmedOwner, TypeAndVersionInterface {
  error ReentrantCall();

  /// @notice This error is returned when the data with report is invalid.
  /// This can happen if the data is shorter than SELECTOR_LENGTH + REPORT_HEADER_LENGTH.
  /// @param data the data that was received
  error InvalidData(bytes data);

  /// @notice This error is thrown whenever trying to set a config
  /// with a fault tolerance of 0
  error FaultToleranceMustBePositive();

  /// @notice This error is thrown whenever a report is signed
  /// with more than the max number of signers
  /// @param numSigners The number of signers who have signed the report
  /// @param maxSigners The maximum number of signers that can sign a report
  error ExcessSigners(uint256 numSigners, uint256 maxSigners);

  /// @notice This error is thrown whenever a report is signed
  /// with less than the minimum number of signers
  /// @param numSigners The number of signers who have signed the report
  /// @param minSigners The minimum number of signers that need to sign a report
  error InsufficientSigners(uint256 numSigners, uint256 minSigners);

  error RepeatedSigner();

  error WrongNumberOfSignatures();

  error InvalidDonId();
  error InvalidSigner();

  uint256 private constant SELECTOR_LENGTH = 4;

  bool internal s_reentrancyGuard; // guard against reentrancy

  /// @notice Contains the signing address of each oracle
  struct OracleSet {
    uint8 f; // Number of faulty nodes allowed
    address[] signers;
    mapping(address => uint256) _positions; // 1-indexed to detect unset values
  }
  mapping(bytes4 donId => OracleSet) internal s_configs;

  // reportId = keccak256(bytes20(receiver) | selector | workflowExecutionId)
  mapping(bytes32 reportId => address transmitter) internal s_reports;

  constructor() ConfirmedOwner(msg.sender) {}

  uint256 internal constant MAX_ORACLES = 31;

  // NOTE: we don't inherit OCR2Base since unlike aggregator we only care about signers, not transmitters
  // and the signers don't fetch their config from the forwarder
  function setConfig(bytes4 donId, uint8 f, address[] calldata signers) external nonReentrant {
    if (f == 0) revert FaultToleranceMustBePositive();
    if (signers.length > MAX_ORACLES) revert ExcessSigners(signers.length, MAX_ORACLES);
    if (signers.length <= 3 * f) revert InsufficientSigners(signers.length, 3 * f + 1);

    // TODO: how does setConfig handle expiration? e.g. if the signer set changes

    // remove any old signer addresses
    for (uint256 i = 0; i < s_configs[donId].signers.length; ++i) {
      address signer = s_configs[donId].signers[i];
      delete s_configs[donId]._positions[signer];
    }
    delete s_configs[donId];
    s_configs[donId].signers = signers;

    // add new signer addresses
    for (uint256 i = 0; i < signers.length; ++i) {
      // assign indices, detect duplicates
      address signer = signers[i];
      if(s_configs[donId]._positions[signer] != 0) revert RepeatedSigner();
      s_configs[donId]._positions[signer] = uint8(i) + 1;
      s_configs[donId].signers.push(signer);
    }
    s_configs[donId].f = f;
  }

  function getSelector(bytes memory _data) private pure returns(bytes4 sig) {
      assembly {
          sig := mload(add(_data, 32))
      }
  }

  // send a report to receiver
  function report(
    address receiver,
    bytes calldata data,
    bytes[] calldata signatures
  ) external nonReentrant returns (bool) {
    if (data.length < SELECTOR_LENGTH + Utils.REPORT_HEADER_LENGTH) {
      revert InvalidData(data);
    }

    // data is an encoded call with the selector prefixed: (bytes4 selector, bytes report, ...)
    // we are able to partially decode just the first param, since we don't know the rest
    bytes4 selector = getSelector(data);
    bytes memory rawReport = abi.decode(data[4:], (bytes));

    (/* bytes32 workflowId */, bytes4 donId, bytes32 workflowExecutionId) = Utils._splitReport(rawReport);

    if (signatures.length != s_configs[donId].f + 1) {
      revert WrongNumberOfSignatures();
    }

    // f can never be 0, so this means the config doesn't actually exist
    if (s_configs[donId].f == 0) revert InvalidDonId();

    bytes32 hash = keccak256(rawReport);

    // validate signatures
    address[MAX_ORACLES] memory signed;
    uint8 index = 0;
    for (uint256 i = 0; i < signatures.length; i++) {
      // TODO: is libocr-style multiple bytes32 arrays more optimal, gas-wise?
      (bytes32 r, bytes32 s, uint8 v) = Utils._splitSignature(signatures[i]);
      address signer = ecrecover(hash, v, r, s);

      // validate signer is trusted and signature is unique
      index = uint8(s_configs[donId]._positions[signer]);
      if (index == 0) revert InvalidSigner(); // index is 1-indexed so we can detect unset signers
      index -= 1;
      if(signed[index] != address(0)) revert RepeatedSigner();
      signed[index] = signer;
    }

    bytes32 reportId = _reportId(receiver, selector, workflowExecutionId);

    if (s_reports[reportId] != address(0)) {
      // report was already processed
      return false;
    }

    // solhint-disable-next-line avoid-low-level-calls
    (bool success, bytes memory result) = receiver.call(data);

    s_reports[reportId] = msg.sender;
    return true;
  }

  function _reportId(address receiver, bytes4 selector, bytes32 workflowExecutionId) internal pure returns (bytes32) {
    // TODO: gas savings: could we just use a bytes key and avoid another keccak256 call
    return keccak256(bytes.concat(bytes20(uint160(receiver)), selector, workflowExecutionId));
  }

  // get transmitter of a given report or 0x0 if it wasn't transmitted yet
  function getTransmitter(address receiver, bytes4 selector, bytes32 workflowExecutionId) external view returns (address) {
    bytes32 reportId = _reportId(receiver, selector, workflowExecutionId);
    return s_reports[reportId];
  }

  /// @inheritdoc TypeAndVersionInterface
  function typeAndVersion() external pure override returns (string memory) {
    return "KeystoneForwarder 1.0.0";
  }

  /**
   * @dev replicates Open Zeppelin's ReentrancyGuard but optimized to fit our storage
   */
  modifier nonReentrant() {
    if (s_reentrancyGuard) revert ReentrantCall();
    s_reentrancyGuard = true;
    _;
    s_reentrancyGuard = false;
  }
}
