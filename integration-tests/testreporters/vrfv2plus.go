package testreporters

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	tc "github.com/smartcontractkit/chainlink/integration-tests/testconfig"

	"github.com/slack-go/slack"

	"github.com/smartcontractkit/chainlink-testing-framework/testreporters"
)

type VRFV2PlusTestReporter struct {
	TestType                     string
	RequestCount                 *big.Int
	FulfilmentCount              *big.Int
	AverageFulfillmentInMillions *big.Int
	SlowestFulfillment           *big.Int
	FastestFulfillment           *big.Int
	testConfig                   *tc.TestConfig
}

func (o *VRFV2PlusTestReporter) SetReportData(
	testType string,
	RequestCount *big.Int,
	FulfilmentCount *big.Int,
	AverageFulfillmentInMillions *big.Int,
	SlowestFulfillment *big.Int,
	FastestFulfillment *big.Int,
	testConfig *tc.TestConfig,
) {
	o.TestType = testType
	o.RequestCount = RequestCount
	o.FulfilmentCount = FulfilmentCount
	o.AverageFulfillmentInMillions = AverageFulfillmentInMillions
	o.SlowestFulfillment = SlowestFulfillment
	o.FastestFulfillment = FastestFulfillment
	o.testConfig = testConfig
}

// SendSlackNotification sends a slack message to a slack webhook
func (o *VRFV2PlusTestReporter) SendSlackNotification(t *testing.T, slackClient *slack.Client, testConfig *tc.TestConfig) error {
	if slackClient == nil {
		slackClient = slack.New(testreporters.SlackAPIKey)
	}

	testFailed := t.Failed()
	headerText := fmt.Sprintf(":white_check_mark: VRF V2 Plus %s Test PASSED :white_check_mark:", o.TestType)
	if testFailed {
		headerText = fmt.Sprintf(":x: VRF V2 Plus %s Test FAILED :x:", o.TestType)
	}

	vrfv2lusConfig := o.testConfig.VRFv2Plus.Common
	messageBlocks := testreporters.SlackNotifyBlocks(headerText, strings.Join(testConfig.Network.SelectedNetworks, ","), []string{
		fmt.Sprintf(
			"Summary\n"+
				"Perf Test Type: %s\n"+
				"Test Duration set in parameters: %s\n"+
				"Use Existing Env: %t\n"+
				"Request Count: %s\n"+
				"Fulfilment Count: %s\n"+
				"AverageFulfillmentInMillions: %s\n"+
				"Slowest Fulfillment: %s\n"+
				"Fastest Fulfillment: %s \n"+
				"RPS: %d\n"+
				"RateLimitUnitDuration: %s\n"+
				"RandomnessRequestCountPerRequest: %d\n"+
				"RandomnessRequestCountPerRequestDeviation: %d\n",
			o.TestType,
			vrfv2lusConfig.TestDuration.Truncate(time.Second).String(),
			vrfv2lusConfig.UseExistingEnv,
			o.RequestCount.String(),
			o.FulfilmentCount.String(),
			o.AverageFulfillmentInMillions.String(),
			o.SlowestFulfillment.String(),
			o.FastestFulfillment.String(),
			vrfv2lusConfig.RPS,
			vrfv2lusConfig.RateLimitUnitDuration.String(),
			vrfv2lusConfig.RandomnessRequestCountPerRequest,
			vrfv2lusConfig.RandomnessRequestCountPerRequestDeviation,
		),
	})

	_, err := testreporters.SendSlackMessage(slackClient, slack.MsgOptionBlocks(messageBlocks...))
	if err != nil {
		return err
	}
	return nil
}
