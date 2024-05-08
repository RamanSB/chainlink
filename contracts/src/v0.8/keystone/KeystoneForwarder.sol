// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/console.sol";
import {IForwarder} from "./interfaces/IForwarder.sol";
import {IReceiver} from "./interfaces/IReceiver.sol";
import {ConfirmedOwner} from "../shared/access/ConfirmedOwner.sol";
import {TypeAndVersionInterface} from "../interfaces/TypeAndVersionInterface.sol";
import {Utils} from "./libraries/Utils.sol";

// solhint-disable gas-custom-errors, no-unused-vars
contract KeystoneForwarder is IForwarder, ConfirmedOwner, TypeAndVersionInterface {
  error ReentrantCall();

  /// @notice This error is returned when the report data is invalid.
  /// This can happen if the data is shorter than REPORT_HEADER_LENGTH.
  error InvalidReport();

  /// @notice This error is thrown whenever trying to set a config
  /// with a fault tolerance of 0
  error FaultToleranceMustBePositive();

  /// @notice This error is thrown whenever configuration provides
  /// more signers than the maximum allowed number
  /// @param numSigners The number of signers who have signed the report
  /// @param maxSigners The maximum number of signers that can sign a report
  error ExcessSigners(uint256 numSigners, uint256 maxSigners);

  /// @notice This error is thrown whenever a configuration is provided
  /// with less than the minimum number of signers
  /// @param numSigners The number of signers provided
  /// @param minSigners The minimum number of signers expected
  error InsufficientSigners(uint256 numSigners, uint256 minSigners);

  error DuplicateSigner();

  error WrongNumberOfSignatures(uint256 expected, uint256 received);

  error InvalidDonId(uint32 donId);
  error InvalidSigner(address signer);
  error ReportAlreadyProcessed();

  bool internal s_reentrancyGuard; // guard against reentrancy

  /// @notice Contains the signing address of each oracle
  struct OracleSet {
    uint8 f; // Number of faulty nodes allowed
    address[] signers;
    mapping(address => uint256) _positions; // 1-indexed to detect unset values
  }
  mapping(uint32 donId => OracleSet) internal s_configs;

  struct DeliveryStatus {
    address transmitter;
    bool success;
  }

  // reportId = keccak256(bytes20(receiver) | workflowOwner | workflowExecutionId)
  mapping(bytes32 reportId => DeliveryStatus status) internal s_reports;

  event ReportProcessed(
    address indexed receiver,
    address indexed workflowOwner,
    bytes32 indexed workflowExecutionId,
    bool result
  );

  constructor() ConfirmedOwner(msg.sender) {}

  uint256 internal constant MAX_ORACLES = 31;

  function setConfig(uint32 donId, uint8 f, address[] calldata signers) external nonReentrant {
    if (f == 0) revert FaultToleranceMustBePositive();
    if (signers.length > MAX_ORACLES) revert ExcessSigners(signers.length, MAX_ORACLES);
    if (signers.length <= 3 * f) revert InsufficientSigners(signers.length, 3 * f + 1);

    // TODO: how does setConfig handle expiration? e.g. if the signer set changes

    // remove any old signer addresses
    for (uint256 i = 0; i < s_configs[donId].signers.length; ++i) {
      address signer = s_configs[donId].signers[i];
      delete s_configs[donId]._positions[signer];
    }

    // add new signer addresses
    s_configs[donId].signers = signers;
    for (uint256 i = 0; i < signers.length; ++i) {
      // assign indices, detect duplicates
      address signer = signers[i];
      if (s_configs[donId]._positions[signer] != 0) revert DuplicateSigner();
      s_configs[donId]._positions[signer] = uint8(i) + 1;
      s_configs[donId].signers.push(signer);
    }
    s_configs[donId].f = f;
  }

  // send a report to receiver
  function report(
    address receiverAddress,
    bytes calldata rawReport,
    bytes[] calldata signatures
  ) external nonReentrant {
    if (rawReport.length < Utils.REPORT_HEADER_LENGTH) {
      revert InvalidReport();
    }

    (bytes32 workflowId, bytes4 donIdBytes, bytes32 workflowExecutionId, bytes32 workflowOwner) = Utils._splitReport(
      rawReport
    );

    uint32 donId = uint32(donIdBytes);

    if (signatures.length != s_configs[donId].f + 1) {
      revert WrongNumberOfSignatures(s_configs[donId].f + 1, signatures.length);
    }

    // f can never be 0, so this means the config doesn't actually exist
    if (s_configs[donId].f == 0) revert InvalidDonId(donId);

    bytes32 reportId = _reportId(receiverAddress, address(bytes20(workflowOwner)), workflowExecutionId);

    if (s_reports[reportId].transmitter != address(0)) {
      revert ReportAlreadyProcessed();
    }

    // validate signatures
    {
      bytes32 hash = keccak256(rawReport);

      address[MAX_ORACLES] memory signed;
      uint8 index = 0;
      for (uint256 i = 0; i < signatures.length; i++) {
        // TODO: is libocr-style multiple bytes32 arrays more optimal, gas-wise?
        (bytes32 r, bytes32 s, uint8 v) = Utils._splitSignature(signatures[i]);
        address signer = ecrecover(hash, v, r, s);

        // validate signer is trusted and signature is unique
        index = uint8(s_configs[donId]._positions[signer]);
        if (index == 0) revert InvalidSigner(signer); // index is 1-indexed so we can detect unset signers
        index -= 1;
        if (signed[index] != address(0)) revert DuplicateSigner();
        signed[index] = signer;
      }
    }

    console.log("before call");

    IReceiver receiver = IReceiver(receiverAddress);
    bool success;
    try receiver.onReport(workflowId, address(bytes20(workflowOwner)), rawReport[Utils.REPORT_HEADER_LENGTH:]) {
      console.log("called");
      success = true;
    } catch (bytes memory reason) {
      console.log(string(reason));
      console.log("error");
      success = false;
    }

    s_reports[reportId] = DeliveryStatus(msg.sender, success);

    emit ReportProcessed(receiverAddress, address(bytes20(workflowOwner)), workflowExecutionId, success);
  }

  function _reportId(
    address receiver,
    address workflowOwner,
    bytes32 workflowExecutionId
  ) internal pure returns (bytes32) {
    // TODO: gas savings: could we just use a bytes key and avoid another keccak256 call
    return keccak256(bytes.concat(bytes20(uint160(receiver)), bytes20(workflowOwner), workflowExecutionId));
  }

  // get transmitter of a given report or 0x0 if it wasn't transmitted yet
  function getTransmitter(
    address receiver,
    address workflowOwner,
    bytes32 workflowExecutionId
  ) external view returns (address) {
    bytes32 reportId = _reportId(receiver, workflowOwner, workflowExecutionId);
    return s_reports[reportId].transmitter;
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
