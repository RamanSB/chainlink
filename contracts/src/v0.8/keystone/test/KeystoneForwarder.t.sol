// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Test.sol";

import "../KeystoneForwarder.sol";
import "./mocks/Receiver.sol";

contract KeystoneForwarderTest is Test {
  function test_abi_partial_decoding_works() public {
    bytes memory report = hex"0102";
    uint256 amount = 1;
    bytes memory payload = abi.encode(report, amount);
    bytes memory decodedReport = abi.decode(payload, (bytes));
    assertEq(decodedReport, report, "not equal");
  }

  address internal constant TRANSMITTER = address(50);
  uint256 internal constant MAX_ORACLES = 31;

  struct Signer {
    uint256 mockPrivateKey;
    address signerAddress;
  }
  Signer[MAX_ORACLES] internal s_signers;

  function setUp() public virtual {
    uint256 seed = 0;
    generateSigners(seed);
  }

  function generateSigners(uint256 seed) internal {
    // generate signers
    for (uint256 i; i < MAX_ORACLES; i++) {
      uint256 mockPK = seed + i + 1;
      s_signers[i].mockPrivateKey = mockPK;
      s_signers[i].signerAddress = vm.addr(mockPK);
    }
  }

  function _getSignerAddresses() internal view returns (address[] memory) {
    address[] memory signerAddrs = new address[](s_signers.length);
    for (uint256 i = 0; i < signerAddrs.length; i++) {
      signerAddrs[i] = s_signers[i].signerAddress;
    }
    return signerAddrs;
  }

  function _generateSignatures(bytes memory report, uint256 numSignatures) internal view returns (bytes[] memory) {
    bytes32 hash = keccak256(report);
    bytes[] memory signatures = new bytes[](numSignatures);

    for (uint256 i = 0; i < numSignatures; i++) {
      (uint8 v, bytes32 r, bytes32 s) = vm.sign(s_signers[i].mockPrivateKey, hash);
      signatures[i] = bytes.concat(r, s, bytes1(v));
    }
    return signatures;
  }

  function test_it_works() public {
    KeystoneForwarder forwarder = new KeystoneForwarder();
    Receiver receiver = new Receiver();

    // generate signers
    setUp();

    // configure contract with signers
    uint8 f = 1;
    bytes4 donId = 0x01020304;
    {
      address[] memory signers = _getSignerAddresses();
      forwarder.setConfig(donId, f, signers);
    }

    // set config doesn't allow duplicate signers
    {
      address[] memory signers = _getSignerAddresses();
      signers[1] = signers[0];

      vm.expectRevert(KeystoneForwarder.RepeatedSigner.selector);
      forwarder.setConfig(donId, f, signers);
    }

    // taken from https://github.com/smartcontractkit/chainlink/blob/2390ec7f3c56de783ef4e15477e99729f188c524/core/services/relay/evm/cap_encoder_test.go#L42-L55
    bytes
      memory report = hex"6d795f6964000000000000000000000000000000000000000000000000000000010203046d795f657865637574696f6e5f696400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000301020300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004aabbccdd00000000000000000000000000000000000000000000000000000000";
    bytes memory data = abi.encodeWithSignature("foo(bytes)", report);

    // generate signatures
    uint256 numSignatures = f + 1;
    bytes[] memory signatures = _generateSignatures(report, numSignatures);

    vm.expectCall(address(receiver), data);
    vm.recordLogs();

    vm.startPrank(TRANSMITTER);
    bool delivered1 = forwarder.report(address(receiver), data, signatures);
    assertTrue(delivered1, "report not delivered");

    bytes32 workflowId = hex"6d795f6964000000000000000000000000000000000000000000000000000000";
    bytes32 executionId = hex"6d795f657865637574696f6e5f69640000000000000000000000000000000000";
    {
      // validate receiver was called
      Vm.Log[] memory entries = vm.getRecordedLogs();
      assertEq(entries[0].emitter, address(receiver));
      // validate workflow id and workflow execution id
      assertEq(entries[0].topics[1], workflowId);
      assertEq(entries[0].topics[2], executionId);
      assertEq(entries[0].topics[3], donId);
      bytes[] memory mercuryReports = abi.decode(entries[0].data, (bytes[]));
      assertEq(mercuryReports.length, 2);
      assertEq(mercuryReports[0], hex"010203");
      assertEq(mercuryReports[1], hex"aabbccdd");
    }

    {
      // validate transmitter was recorded
      address transmitter = forwarder.getTransmitter(address(receiver), receiver.foo.selector, executionId);
      assertEq(transmitter, TRANSMITTER);
    }

    {
      // doesn't deliver the same report more than once
      bool delivered2 = forwarder.report(address(receiver), data, signatures);
      assertFalse(delivered2, "report redelivered");
    }

    {
      // doesn't allow duplicate signers
      signatures[1] = signatures[0];
      vm.expectRevert(KeystoneForwarder.RepeatedSigner.selector);
      bool delivered3 = forwarder.report(address(receiver), data, signatures);
    }
  }
}
