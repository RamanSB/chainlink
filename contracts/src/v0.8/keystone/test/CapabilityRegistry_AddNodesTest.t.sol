// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {BaseTest} from "./BaseTest.t.sol";
import {CapabilityRegistry} from "../CapabilityRegistry.sol";

contract CapabilityRegistry_AddNodesTest is BaseTest {
  event NodeAdded(bytes32 p2pId, uint256 nodeOperatorId);

  uint32 private constant TEST_NODE_OPERATOR_ONE_ID = 0;
  uint256 private constant TEST_NODE_OPERATOR_TWO_ID = 1;

  function setUp() public override {
    BaseTest.setUp();
    changePrank(ADMIN);
    s_capabilityRegistry.addNodeOperators(_getNodeOperators());
    s_capabilityRegistry.addCapability(s_basicCapability);
    s_capabilityRegistry.addCapability(s_capabilityWithConfigurationContract);
  }

  function test_RevertWhen_CalledByNonNodeOperatorAdminAndNonOwner() public {
    changePrank(STRANGER);
    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);

    bytes32[] memory hashedCapabilityIds = new bytes32[](1);
    hashedCapabilityIds[0] = s_basicHashedCapabilityId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectRevert(CapabilityRegistry.AccessForbidden.selector);
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);
  }

  function test_RevertWhen_NumNodesNotEqualToNumCapabilityConfigs() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);

    bytes32[] memory hashedCapabilityIds = new bytes32[](1);
    hashedCapabilityIds[0] = s_basicHashedCapabilityId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](2);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectRevert(
      abi.encodeWithSelector(CapabilityRegistry.LengthMismatch.selector, nodes.length, nodeCapabilityIds.length)
    );
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);
  }

  function test_RevertWhen_SignerAddressEmpty() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);

    bytes32[] memory hashedCapabilityIds = new bytes32[](1);
    hashedCapabilityIds[0] = s_basicHashedCapabilityId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: address(0)
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectRevert(abi.encodeWithSelector(CapabilityRegistry.InvalidNodeSigner.selector));
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);
  }

  function test_RevertWhen_AddingDuplicateP2PId() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);

    bytes32[] memory hashedCapabilityIds = new bytes32[](1);
    hashedCapabilityIds[0] = s_basicHashedCapabilityId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);

    vm.expectRevert(abi.encodeWithSelector(CapabilityRegistry.InvalidNodeP2PId.selector, P2P_ID));
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);
  }

  function test_RevertWhen_P2PIDEmpty() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);

    bytes32[] memory hashedCapabilityIds = new bytes32[](1);
    hashedCapabilityIds[0] = s_basicHashedCapabilityId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: bytes32(""),
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectRevert(abi.encodeWithSelector(CapabilityRegistry.InvalidNodeP2PId.selector, bytes32("")));
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);
  }

  function test_RevertWhen_AddingNodeWithoutCapabilities() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);

    bytes32[] memory hashedCapabilityIds = new bytes32[](0);

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectRevert(abi.encodeWithSelector(CapabilityRegistry.InvalidNodeCapabilities.selector, hashedCapabilityIds));
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);
  }

  function test_RevertWhen_AddingNodeWithInvalidCapability() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);

    bytes32[] memory hashedCapabilityIds = new bytes32[](1);
    hashedCapabilityIds[0] = s_nonExistentHashedCapabilityId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectRevert(abi.encodeWithSelector(CapabilityRegistry.InvalidNodeCapabilities.selector, hashedCapabilityIds));
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);
  }

  function test_AddsNodeParams() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);

    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);
    bytes32[] memory hashedCapabilityIds = new bytes32[](2);
    hashedCapabilityIds[0] = s_basicHashedCapabilityId;
    hashedCapabilityIds[1] = s_capabilityWithConfigurationContractId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectEmit(address(s_capabilityRegistry));
    emit NodeAdded(P2P_ID, TEST_NODE_OPERATOR_ONE_ID);
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);

    (CapabilityRegistry.NodeParams memory node, bytes32[] memory supportedHashedCapabilityIds) = s_capabilityRegistry
      .getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(node.p2pId, P2P_ID);
    assertEq(supportedHashedCapabilityIds.length, 2);
    assertEq(supportedHashedCapabilityIds[0], s_basicHashedCapabilityId);
    assertEq(supportedHashedCapabilityIds[1], s_capabilityWithConfigurationContractId);
  }

  function test_OwnerCanAddNodes() public {
    changePrank(ADMIN);

    CapabilityRegistry.NodeParams[] memory nodes = new CapabilityRegistry.NodeParams[](1);
    bytes32[] memory hashedCapabilityIds = new bytes32[](2);
    hashedCapabilityIds[0] = s_basicHashedCapabilityId;
    hashedCapabilityIds[1] = s_capabilityWithConfigurationContractId;

    nodes[0] = CapabilityRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS
    });

    bytes32[][] memory nodeCapabilityIds = new bytes32[][](1);
    nodeCapabilityIds[0] = hashedCapabilityIds;

    vm.expectEmit(address(s_capabilityRegistry));
    emit NodeAdded(P2P_ID, TEST_NODE_OPERATOR_ONE_ID);
    s_capabilityRegistry.addNodes(nodes, nodeCapabilityIds);

    (CapabilityRegistry.NodeParams memory node, bytes32[] memory supportedHashedCapabilityIds) = s_capabilityRegistry
      .getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(node.p2pId, P2P_ID);
    assertEq(supportedHashedCapabilityIds.length, 2);
    assertEq(supportedHashedCapabilityIds[0], s_basicHashedCapabilityId);
    assertEq(supportedHashedCapabilityIds[1], s_capabilityWithConfigurationContractId);
  }
}
