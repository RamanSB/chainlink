// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {KeystoneForwarder} from "../KeystoneForwarder.sol";

contract KeystoneForwarder_SetConfigTest is Test {
  address internal constant TRANSMITTER = address(50);
  uint256 internal constant MAX_ORACLES = 31;

  struct Signer {
    uint256 mockPrivateKey;
    address signerAddress;
  }
  Signer[MAX_ORACLES] internal s_signers;

  function setUp() public virtual {
    uint256 seed = 0;

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

  function test_RevertWhen_FaultToleranceIsZero() public {
    KeystoneForwarder forwarder = new KeystoneForwarder();
    setUp();

    uint8 f = 0;
    bytes4 donId = 0x01020304;

    address[] memory signers = _getSignerAddresses();

    vm.expectRevert(KeystoneForwarder.FaultToleranceMustBePositive.selector);
    forwarder.setConfig(donId, f, signers);
  }

  function test_RevertWhen_InsufficientSigners() public {
    KeystoneForwarder forwarder = new KeystoneForwarder();
    setUp();

    uint8 f = 1;
    bytes4 donId = 0x01020304;
    address[] memory signers = new address[](1);

    vm.expectRevert(abi.encodeWithSelector(KeystoneForwarder.InsufficientSigners.selector, 1, 4));
    forwarder.setConfig(donId, f, signers);
  }

  function test_RevertWhen_ExcessSigners() public {
    KeystoneForwarder forwarder = new KeystoneForwarder();

    setUp();

    uint8 f = 1;
    bytes4 donId = 0x01020304;
    address[] memory signers = new address[](64);

    vm.expectRevert(abi.encodeWithSelector(KeystoneForwarder.ExcessSigners.selector, 64, 31));
    forwarder.setConfig(donId, f, signers);
  }

  function test_RevertWhen_ProvidingDuplicateSigners() public {
    KeystoneForwarder forwarder = new KeystoneForwarder();
    setUp();

    uint8 f = 1;
    bytes4 donId = 0x01020304;
    address[] memory signers = _getSignerAddresses();
    signers[1] = signers[0];

    vm.expectRevert(KeystoneForwarder.DuplicateSigner.selector);
    forwarder.setConfig(donId, f, signers);
  }

  function test_SetConfig() public {
    KeystoneForwarder forwarder = new KeystoneForwarder();
    setUp();

    // configure contract with signers
    uint8 f = 1;
    bytes4 donId = 0x01020304;
    address[] memory signers = _getSignerAddresses();

    forwarder.setConfig(donId, f, signers);
  }
}
