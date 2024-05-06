package standardcapability

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/generic"
	"github.com/smartcontractkit/chainlink/v2/core/services/telemetry"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type Delegate struct {
	logger                logger.Logger
	ds                    sqlutil.DataSource
	registry              core.CapabilitiesRegistry
	cfg                   plugins.RegistrarConfig
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
}

func NewDelegate(logger logger.Logger, ds sqlutil.DataSource, registry core.CapabilitiesRegistry,
	cfg plugins.RegistrarConfig, monitoringEndpointGen telemetry.MonitoringEndpointGenerator) *Delegate {
	return &Delegate{logger: logger, ds: ds, registry: registry, cfg: cfg, monitoringEndpointGen: monitoringEndpointGen}
}

func (d Delegate) JobType() job.Type {
	return job.StandardCapability
}

func (d Delegate) BeforeJobCreated(job job.Job) {}

func (d Delegate) ServicesForSpec(ctx context.Context, jb job.Job) ([]job.ServiceCtx, error) {

	log := d.logger.Named("StandardCapability").Named(jb.StandardCapabilitySpec.GetID())

	cmdName := jb.StandardCapabilitySpec.BinaryUrl

	// TEMP override
	cmdName = "/Users/matthewpendrey/Projects/chainlink/core/services/standardcapability/simpletriggercapability/simpletriggercapability" // get a better version of this from the test code

	cmdFn, opts, err := d.cfg.RegisterLOOP(plugins.CmdConfig{
		ID:  log.Name(),
		Cmd: cmdName,
		Env: nil,
	})

	if err != nil {
		return nil, fmt.Errorf("error registering loop: %v", err)
	}

	capabilityLoop := loop.NewStandardCapability(log, opts, cmdFn)

	// TODO Move the below into service context

	err = capabilityLoop.Start(ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting standard capability service: %v", err)
	}

	err = capabilityLoop.WaitCtx(ctx)
	if err != nil {
		return nil, fmt.Errorf("error waiting for standard capability service to start: %v", err)
	}

	info, err := capabilityLoop.Service.Info(ctx)

	d.logger.Info("Standard capability service", "info", info)

	if err != nil {
		return nil, fmt.Errorf("error getting standard capability service info: %v", err)
	}

	kvStore := job.NewKVStore(jb.ID, d.ds, log)
	telemetryService := generic.NewTelemetryAdapter(d.monitoringEndpointGen)

	err = capabilityLoop.Service.Initialise(ctx, jb.StandardCapabilitySpec.CapabilityConfig, telemetryService, kvStore, d.registry)
	if err != nil {
		return nil, fmt.Errorf("error initialising standard capability service: %v", err)
	}

	//  temp test for now to test registration and communication with the capability

	capability, err := d.registry.GetTrigger(ctx, "SIMPLETRIGGERCAPABILITY")
	if err != nil {
		return nil, fmt.Errorf("error getting action capability: %v", err)
	}

	resultCh, err := capability.RegisterTrigger(ctx, capabilities.CapabilityRequest{})
	if err != nil {
		return nil, fmt.Errorf("error creating standard capability: %v", err)
	}

	for resp := range resultCh {
		fmt.Printf("Got response from standard capability: %v\n", resp.Value)
	}

	// closing logic is cleaning up resources as expected
	err = capabilityLoop.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing standard capability service: %v", err)
	}

	// end of temp test

	// TODO Move initialisation into the service context
	return []job.ServiceCtx{capabilityLoop}, nil
}

func (d Delegate) AfterJobCreated(job job.Job) {}

func (d Delegate) BeforeJobDeleted(job job.Job) {}

func (d Delegate) OnDeleteJob(ctx context.Context, jb job.Job) error { return nil }

func ValidatedStandardCapabilitySpec(tomlString string) (job.Job, error) {
	var jb = job.Job{ExternalJobID: uuid.New()}

	tree, err := toml.Load(tomlString)
	if err != nil {
		return jb, errors.Wrap(err, "toml error on load standard capability")
	}

	err = tree.Unmarshal(&jb)
	if err != nil {
		return jb, errors.Wrap(err, "toml unmarshal error on standard capability spec")
	}

	var spec job.StandardCapabilitySpec
	err = tree.Unmarshal(&spec)
	if err != nil {
		return jb, errors.Wrap(err, "toml unmarshal error on standard capability job")
	}

	jb.StandardCapabilitySpec = &spec
	if jb.Type != job.StandardCapability {
		return jb, errors.Errorf("standard capability unsupported job type %s", jb.Type)
	}

	// TODO other validation

	return jb, nil
}
