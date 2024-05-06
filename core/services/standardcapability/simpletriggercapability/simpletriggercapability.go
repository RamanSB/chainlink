package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-plugin"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

const (
	loggerName = "PluginStandardCapability"
)

func main() {
	s := loop.MustNewStartedServer(loggerName)
	defer s.Stop()

	stopCh := make(chan struct{})
	defer close(stopCh)

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: loop.StandardCapabilityHandshakeConfig(),
		Plugins: map[string]plugin.Plugin{
			loop.PluginStandardCapabilityName: &loop.StandardCapabilityLoop{
				PluginServer: &CustomTriggerCapabilityService{},
				BrokerConfig: loop.BrokerConfig{Logger: s.Logger, StopCh: stopCh, GRPCOpts: s.GRPCOpts},
			},
		},
		GRPCServer: s.GRPCOpts.NewServer,
	})

}

type CustomTriggerCapabilityService struct {
	telemetryService core.TelemetryService
	store            core.KeyValueStore
}

func (c CustomTriggerCapabilityService) Start(ctx context.Context) error {
	return nil
}

func (c CustomTriggerCapabilityService) Close() error {
	return nil
}

func (c CustomTriggerCapabilityService) Ready() error {
	return nil
}

func (c CustomTriggerCapabilityService) HealthReport() map[string]error {
	return nil
}

func (c CustomTriggerCapabilityService) Name() string {
	return ""
}

func (c CustomTriggerCapabilityService) Info(ctx context.Context) (capabilities.CapabilityInfo, error) {
	return capabilities.CapabilityInfo{
		ID:             "SIMPLETRIGGERCAPABILITY",
		CapabilityType: capabilities.CapabilityTypeTrigger,
		Description:    "",
		Version:        "",
		DON:            nil,
	}, nil
}

func (c CustomTriggerCapabilityService) RegisterTrigger(ctx context.Context, request capabilities.CapabilityRequest) (<-chan capabilities.CapabilityResponse, error) {

	result := make(chan capabilities.CapabilityResponse, 100)

	err := c.store.Store(ctx, "key", []byte("value"))
	if err != nil {
		return nil, fmt.Errorf("failed to store key: %w", err)
	}

	go func() {
		defer close(result)
		for i := 0; i < 2; i++ {
			value, err := values.Wrap(fmt.Sprintf("Trigger World! %d", i))
			if err != nil {
				// log
				return
			}

			result <- capabilities.CapabilityResponse{
				Value: value,
				Err:   nil,
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return result, nil

}

func (c CustomTriggerCapabilityService) UnregisterTrigger(ctx context.Context, request capabilities.CapabilityRequest) error {
	return nil
}

func (c CustomTriggerCapabilityService) Initialise(ctx context.Context, config string, telemetryService core.TelemetryService,
	store core.KeyValueStore, capReg core.CapabilitiesRegistry) error {

	c.telemetryService = telemetryService
	c.store = store

	if err := capReg.Add(ctx, c); err != nil {
		return fmt.Errorf("error when adding capability to registry: %w", err)
	}

	return nil
}
