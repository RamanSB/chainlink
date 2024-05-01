
import {Utils} from "../../libraries/Utils.sol";

contract Receiver {
  event MessageReceived(bytes32 indexed workflowId, bytes32 indexed workflowExecutionId, bytes4 indexed donId, bytes[] mercuryReports);

  constructor() {}

  function foo(bytes calldata rawReport) external {
    // decode metadata
    (bytes32 workflowId, bytes4 donId, bytes32 workflowExecutionId) = Utils._splitReport(rawReport);
    // parse actual report
    bytes[] memory mercuryReports = abi.decode(rawReport[Utils.REPORT_HEADER_LENGTH:], (bytes[]));
    emit MessageReceived(workflowId, workflowExecutionId, donId, mercuryReports);
  }
}

