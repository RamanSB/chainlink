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
  /// This can happen if the data is shorter than SELECTOR_LENGTH + REPORT_LENGTH.
  /// @param data the data that was received
  error InvalidData(bytes data);

  error RepeatedSigner();

  error WrongNumberOfSignatures();

  error InvalidSigner();

  uint256 private constant SELECTOR_LENGTH = 4;
  uint256 private constant REPORT_LENGTH = 64;

  struct HotVars {
    bool reentrancyGuard; // guard against reentrancy
    uint8 n;
    uint8 f;
  }

  HotVars internal s_hotVars; // Mixture of config and state, commonly accessed

  /// @notice Contains the signing address of each oracle
  address[] internal s_signers;
  mapping(address => uint8 index) internal s_signersMap; // TODO: is this more gas efficient than a for loop { eq }?

  mapping(bytes32 => address) internal s_reports;

  constructor() ConfirmedOwner(msg.sender) {}

  uint256 internal constant MAX_ORACLES = 31;

  // NOTE: we don't inherit OCR2Base since unlike aggregator we only care about signers, not transmitters
  // and the signers don't fetch their config from the forwarder
  function setConfig(uint8 f, address[] calldata signers) external nonReentrant {
    // TODO: how does setConfig handle expiration? e.g. if the signer set changes

    // remove any old signer/transmitter addresses
    while (s_signers.length != 0) {
      uint256 lastIdx = s_signers.length - 1;
      address signer = s_signers[lastIdx];
      delete s_signersMap[signer];
      s_signers.pop();
    }

    // add new signer/transmitter addresses
    for (uint256 i = 0; i < signers.length; ++i) {
      if(s_signersMap[signers[i]] != 0) revert RepeatedSigner();
      s_signersMap[signers[i]] = uint8(i) + 1;
      s_signers.push(signers[i]);
    }
    s_hotVars.f = f;
  }

  // send a report to targetAddress
  function report(
    address targetAddress,
    bytes calldata data,
    bytes[] calldata signatures
  ) external nonReentrant returns (bool) {
    if (data.length < SELECTOR_LENGTH + REPORT_LENGTH) {
      revert InvalidData(data);
    }

    if (signatures.length != s_hotVars.f + 1) {
      revert WrongNumberOfSignatures();
    }

    // data is an encoded call with the selector prefixed: (bytes4 selector, bytes report, ...)
    // we are able to partially decode just the first param, since we don't know the rest
    bytes memory rawReport = abi.decode(data[4:], (bytes));

    bytes32 hash = keccak256(rawReport);

    address[MAX_ORACLES] memory signed;

    // validate signatures
    uint8 index = 0;
    for (uint256 i = 0; i < signatures.length; i++) {
      // TODO: is libocr-style multiple bytes32 arrays more optimal, gas-wise?
      (bytes32 r, bytes32 s, uint8 v) = Utils._splitSignature(signatures[i]);
      address signer = ecrecover(hash, v, r, s);

      // validate signer is trusted and signature is unique
      index = s_signersMap[signer];
      if (index == 0) revert InvalidSigner(); // index is 1-indexed so we can detect unset signers
      index -= 1;
      if(signed[index] != address(0)) revert RepeatedSigner();
      signed[index] = signer;
    }

    (bytes32 workflowId, bytes32 workflowExecutionId) = Utils._splitReport(rawReport);

    // report was already processed
    if (s_reports[workflowExecutionId] != address(0)) {
      return false;
    }

    // solhint-disable-next-line avoid-low-level-calls
    (bool success, bytes memory result) = targetAddress.call(data);

    s_reports[workflowExecutionId] = msg.sender;
    return true;
  }

  // get transmitter of a given report or 0x0 if it wasn't transmitted yet
  function getTransmitter(bytes32 workflowExecutionId) external view returns (address) {
    return s_reports[workflowExecutionId];
  }

  /// @inheritdoc TypeAndVersionInterface
  function typeAndVersion() external pure override returns (string memory) {
    return "KeystoneForwarder 1.0.0";
  }

  /**
   * @dev replicates Open Zeppelin's ReentrancyGuard but optimized to fit our storage
   */
  modifier nonReentrant() {
    if (s_hotVars.reentrancyGuard) revert ReentrantCall();
    s_hotVars.reentrancyGuard = true;
    _;
    s_hotVars.reentrancyGuard = false;
  }
}
