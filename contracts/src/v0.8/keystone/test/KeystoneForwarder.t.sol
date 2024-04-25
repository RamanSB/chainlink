// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Test.sol";

import "../KeystoneForwarder.sol";
import {Utils} from "../libraries/Utils.sol";

contract Receiver {
  event MessageReceived(bytes32 indexed workflowId, bytes32 indexed workflowExecutionId, bytes[] mercuryReports);

  constructor() {}

  function foo(bytes calldata rawReport) external {
    // decode metadata
    (bytes32 workflowId, bytes32 workflowExecutionId) = Utils._splitReport(rawReport);
    // parse actual report
    bytes[] memory mercuryReports = abi.decode(rawReport[64:], (bytes[]));
    emit MessageReceived(workflowId, workflowExecutionId, mercuryReports);
  }
}

contract KeystoneForwarderTest is Test {
  function setUp() public virtual {}

  function test_abi_partial_decoding_works() public {
    bytes memory report = hex"0102";
    uint256 amount = 1;
    bytes memory payload = abi.encode(report, amount);
    bytes memory decodedReport = abi.decode(payload, (bytes));
    assertEq(decodedReport, report, "not equal");
  }

  uint256 internal constant MAX_ORACLES = 31;

  struct Signer {
    uint256 mockPrivateKey;
    address signerAddress;
  }
  Signer[MAX_ORACLES] internal s_signers;

  function test_it_works() public {
    KeystoneForwarder forwarder = new KeystoneForwarder();
    Receiver receiver = new Receiver();

    // generate signers
    for (uint256 i; i < MAX_ORACLES; i++) {
      uint256 mockPK = i + 1;
      s_signers[i].mockPrivateKey = mockPK;
      s_signers[i].signerAddress = vm.addr(mockPK);
    }

    // configure contract with signers
    uint8 f = 0;
    address[] memory signers = new address[](s_signers.length);
    for (uint256 i = 0; i < s_signers.length; i++) {
      signers[i] = s_signers[i].signerAddress;
    }
    forwarder.setConfig(f, signers);

    // taken from https://github.com/smartcontractkit/chainlink/blob/2390ec7f3c56de783ef4e15477e99729f188c524/core/services/relay/evm/cap_encoder_test.go#L42-L55
    bytes
      memory report = hex"6d795f69640000000000000000000000000000000000000000000000000000006d795f657865637574696f6e5f696400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000301020300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004aabbccdd00000000000000000000000000000000000000000000000000000000";
    bytes memory data = abi.encodeWithSignature("foo(bytes)", report);

    // generate signatures
    bytes32 hash = keccak256(report);

    uint256 numSignatures = f + 1;
    bytes[] memory signatures = new bytes[](numSignatures);

    for (uint256 i = 0; i < numSignatures; i++) {
      (uint8 v, bytes32 r, bytes32 s) = vm.sign(s_signers[i].mockPrivateKey, hash);
      // rs[i] = r;
      // ss[i] = s;
      // vs[i] = bytes1(v - 27);
      signatures[i] = bytes.concat(r, s, bytes1(v));
    }

    vm.expectCall(address(receiver), data);
    vm.recordLogs();

    bool delivered1 = forwarder.report(address(receiver), data, signatures);
    assertTrue(delivered1, "report not delivered");

    Vm.Log[] memory entries = vm.getRecordedLogs();
    assertEq(entries[0].emitter, address(receiver));
    // validate workflow id and workflow execution id
    bytes32 workflowId = hex"6d795f6964000000000000000000000000000000000000000000000000000000";
    bytes32 executionId = hex"6d795f657865637574696f6e5f69640000000000000000000000000000000000";
    assertEq(entries[0].topics[1], workflowId);
    assertEq(entries[0].topics[2], executionId);
    bytes[] memory mercuryReports = abi.decode(entries[0].data, (bytes[]));
    assertEq(mercuryReports.length, 2);
    assertEq(mercuryReports[0], hex"010203");
    assertEq(mercuryReports[1], hex"aabbccdd");

    // doesn't deliver the same report more than once
    bool delivered2 = forwarder.report(address(receiver), data, signatures);
    assertFalse(delivered2, "report redelivered");
  }
}
