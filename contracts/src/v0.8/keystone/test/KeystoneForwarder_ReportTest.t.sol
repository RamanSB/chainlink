// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Vm} from "forge-std/Test.sol";
import {BaseTest} from "./KeystoneForwarderBaseTest.t.sol";
import {KeystoneForwarder} from "../KeystoneForwarder.sol";

contract KeystoneForwarder_ReportTest is BaseTest {
  bytes32 internal workflowId = hex"6d795f6964000000000000000000000000000000000000000000000000000000";
  address internal workflowOwner = address(51);
  bytes32 internal executionId = hex"6d795f657865637574696f6e5f69640000000000000000000000000000000000";
  bytes[] internal mercuryReports = new bytes[](2);
  bytes internal rawReports;
  bytes internal report;
  uint256 internal requiredSignaturesNum = F + 1;
  bytes[] internal signatures = new bytes[](2);

  function setUp() public override {
    BaseTest.setUp();

    s_forwarder.setConfig(DON_ID, F, _getSignerAddresses());

    mercuryReports[0] = hex"010203";
    mercuryReports[1] = hex"aabbccdd";

    rawReports = abi.encode(mercuryReports);
    report = abi.encodePacked(workflowId, DON_ID, executionId, workflowOwner, rawReports);

    bytes32 hash = keccak256(report);

    for (uint256 i = 0; i < requiredSignaturesNum; i++) {
      (uint8 v, bytes32 r, bytes32 s) = vm.sign(s_signers[i].mockPrivateKey, hash);
      signatures[i] = bytes.concat(r, s, bytes1(v));
    }

    vm.startPrank(TRANSMITTER);
  }

  function test_RevertWhen_ReportHasDuplicateSignatures() public {
    signatures[1] = signatures[0]; // repeat a signature

    vm.expectRevert(KeystoneForwarder.DuplicateSigner.selector);
    s_forwarder.report(address(s_receiver), report, signatures);
  }

  function test_RevertWhen_ReportAlreadyProcessed() public {
    s_forwarder.report(address(s_receiver), report, signatures);

    vm.expectRevert(KeystoneForwarder.ReportAlreadyProcessed.selector);
    s_forwarder.report(address(s_receiver), report, signatures);
  }

  function test_report() public {
    // taken from https://github.com/smartcontractkit/chainlink/blob/2390ec7f3c56de783ef4e15477e99729f188c524/core/services/relay/evm/cap_encoder_test.go#L42-L55
    // bytes memory report = hex"6d795f6964000000000000000000000000000000000000000000000000000000010203046d795f657865637574696f6e5f696400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000301020300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004aabbccdd00000000000000000000000000000000000000000000000000000000";
    vm.expectCall(address(s_receiver), abi.encodeCall(s_receiver.onReport, (workflowId, workflowOwner, rawReports)));
    vm.recordLogs();

    s_forwarder.report(address(s_receiver), report, signatures);

    // {
    //   // validate s_receiver was called
    //   Vm.Log[] memory entries = vm.getRecordedLogs();
    //   assertEq(entries.length, 2);
    //   assertEq(entries[0].emitter, address(s_receiver));
    //   // validate params sent to s_receiver
    //   assertEq(entries[0].topics[1], workflowId, "workflow id mismatch");
    //   assertEq(entries[0].topics[2], bytes20(workflowOwner), "owner id mismatch");
    //   bytes[] memory decodedMercuryReports = abi.decode(entries[0].data, (bytes[]));
    //   assertEq(mercuryReports, decodedMercuryReports, "mercury reports mismatch");

    //   assertEq(entries[1].emitter, address(s_forwarder));
    // }

    // {
    //   // validate transmitter was recorded
    //   address transmitter = s_forwarder.getTransmitter(address(s_receiver), workflowOwner, executionId);
    //   assertEq(transmitter, TRANSMITTER, "transmitter mismatch");
    // }
  }
}
