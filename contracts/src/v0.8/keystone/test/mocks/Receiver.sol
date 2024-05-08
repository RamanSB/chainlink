// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Receiver {
  event MessageReceived(bytes32 indexed workflowId, bytes32 indexed workflowOwner, bytes[] mercuryReports);

  constructor() {}

  function onReport(bytes32 workflowId, bytes32 workflowOwner, bytes calldata rawReport) external {
    // parse actual report
    bytes[] memory mercuryReports = abi.decode(rawReport, (bytes[]));
    emit MessageReceived(workflowId, workflowOwner, mercuryReports);
  }
}
