// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	cosmos "github.com/smartcontractkit/chainlink/v2/core/chains/cosmos"
	config "github.com/smartcontractkit/chainlink/v2/core/config"

	mock "github.com/stretchr/testify/mock"

	solana "github.com/smartcontractkit/chainlink-solana/pkg/solana"

	starknet "github.com/smartcontractkit/chainlink/v2/core/chains/starknet"

	time "time"

	toml "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml"

	uuid "github.com/google/uuid"

	zapcore "go.uber.org/zap/zapcore"
)

// GeneralConfig is an autogenerated mock type for the GeneralConfig type
type GeneralConfig struct {
	mock.Mock
}

// AppID provides a mock function with given fields:
func (_m *GeneralConfig) AppID() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// AuditLogger provides a mock function with given fields:
func (_m *GeneralConfig) AuditLogger() config.AuditLogger {
	ret := _m.Called()

	var r0 config.AuditLogger
	if rf, ok := ret.Get(0).(func() config.AuditLogger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.AuditLogger)
		}
	}

	return r0
}

// AutoPprof provides a mock function with given fields:
func (_m *GeneralConfig) AutoPprof() config.AutoPprof {
	ret := _m.Called()

	var r0 config.AutoPprof
	if rf, ok := ret.Get(0).(func() config.AutoPprof); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.AutoPprof)
		}
	}

	return r0
}

// ConfigTOML provides a mock function with given fields:
func (_m *GeneralConfig) ConfigTOML() (string, string) {
	ret := _m.Called()

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func() (string, string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// CosmosConfigs provides a mock function with given fields:
func (_m *GeneralConfig) CosmosConfigs() cosmos.CosmosConfigs {
	ret := _m.Called()

	var r0 cosmos.CosmosConfigs
	if rf, ok := ret.Get(0).(func() cosmos.CosmosConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos.CosmosConfigs)
		}
	}

	return r0
}

// CosmosEnabled provides a mock function with given fields:
func (_m *GeneralConfig) CosmosEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Database provides a mock function with given fields:
func (_m *GeneralConfig) Database() config.Database {
	ret := _m.Called()

	var r0 config.Database
	if rf, ok := ret.Get(0).(func() config.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Database)
		}
	}

	return r0
}

// EVMConfigs provides a mock function with given fields:
func (_m *GeneralConfig) EVMConfigs() toml.EVMConfigs {
	ret := _m.Called()

	var r0 toml.EVMConfigs
	if rf, ok := ret.Get(0).(func() toml.EVMConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(toml.EVMConfigs)
		}
	}

	return r0
}

// EVMEnabled provides a mock function with given fields:
func (_m *GeneralConfig) EVMEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EVMRPCEnabled provides a mock function with given fields:
func (_m *GeneralConfig) EVMRPCEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Feature provides a mock function with given fields:
func (_m *GeneralConfig) Feature() config.Feature {
	ret := _m.Called()

	var r0 config.Feature
	if rf, ok := ret.Get(0).(func() config.Feature); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Feature)
		}
	}

	return r0
}

// FluxMonitor provides a mock function with given fields:
func (_m *GeneralConfig) FluxMonitor() config.FluxMonitor {
	ret := _m.Called()

	var r0 config.FluxMonitor
	if rf, ok := ret.Get(0).(func() config.FluxMonitor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.FluxMonitor)
		}
	}

	return r0
}

// Insecure provides a mock function with given fields:
func (_m *GeneralConfig) Insecure() config.Insecure {
	ret := _m.Called()

	var r0 config.Insecure
	if rf, ok := ret.Get(0).(func() config.Insecure); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Insecure)
		}
	}

	return r0
}

// InsecureFastScrypt provides a mock function with given fields:
func (_m *GeneralConfig) InsecureFastScrypt() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JobPipeline provides a mock function with given fields:
func (_m *GeneralConfig) JobPipeline() config.JobPipeline {
	ret := _m.Called()

	var r0 config.JobPipeline
	if rf, ok := ret.Get(0).(func() config.JobPipeline); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.JobPipeline)
		}
	}

	return r0
}

// Keeper provides a mock function with given fields:
func (_m *GeneralConfig) Keeper() config.Keeper {
	ret := _m.Called()

	var r0 config.Keeper
	if rf, ok := ret.Get(0).(func() config.Keeper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Keeper)
		}
	}

	return r0
}

// Log provides a mock function with given fields:
func (_m *GeneralConfig) Log() config.Log {
	ret := _m.Called()

	var r0 config.Log
	if rf, ok := ret.Get(0).(func() config.Log); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Log)
		}
	}

	return r0
}

// LogConfiguration provides a mock function with given fields: log
func (_m *GeneralConfig) LogConfiguration(log config.LogfFn) {
	_m.Called(log)
}

// Mercury provides a mock function with given fields:
func (_m *GeneralConfig) Mercury() config.Mercury {
	ret := _m.Called()

	var r0 config.Mercury
	if rf, ok := ret.Get(0).(func() config.Mercury); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Mercury)
		}
	}

	return r0
}

// OCR provides a mock function with given fields:
func (_m *GeneralConfig) OCR() config.OCR {
	ret := _m.Called()

	var r0 config.OCR
	if rf, ok := ret.Get(0).(func() config.OCR); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.OCR)
		}
	}

	return r0
}

// OCR2 provides a mock function with given fields:
func (_m *GeneralConfig) OCR2() config.OCR2 {
	ret := _m.Called()

	var r0 config.OCR2
	if rf, ok := ret.Get(0).(func() config.OCR2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.OCR2)
		}
	}

	return r0
}

// P2P provides a mock function with given fields:
func (_m *GeneralConfig) P2P() config.P2P {
	ret := _m.Called()

	var r0 config.P2P
	if rf, ok := ret.Get(0).(func() config.P2P); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.P2P)
		}
	}

	return r0
}

// Password provides a mock function with given fields:
func (_m *GeneralConfig) Password() config.Password {
	ret := _m.Called()

	var r0 config.Password
	if rf, ok := ret.Get(0).(func() config.Password); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Password)
		}
	}

	return r0
}

// Prometheus provides a mock function with given fields:
func (_m *GeneralConfig) Prometheus() config.Prometheus {
	ret := _m.Called()

	var r0 config.Prometheus
	if rf, ok := ret.Get(0).(func() config.Prometheus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Prometheus)
		}
	}

	return r0
}

// Pyroscope provides a mock function with given fields:
func (_m *GeneralConfig) Pyroscope() config.Pyroscope {
	ret := _m.Called()

	var r0 config.Pyroscope
	if rf, ok := ret.Get(0).(func() config.Pyroscope); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Pyroscope)
		}
	}

	return r0
}

// RootDir provides a mock function with given fields:
func (_m *GeneralConfig) RootDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Sentry provides a mock function with given fields:
func (_m *GeneralConfig) Sentry() config.Sentry {
	ret := _m.Called()

	var r0 config.Sentry
	if rf, ok := ret.Get(0).(func() config.Sentry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Sentry)
		}
	}

	return r0
}

// SetLogLevel provides a mock function with given fields: lvl
func (_m *GeneralConfig) SetLogLevel(lvl zapcore.Level) error {
	ret := _m.Called(lvl)

	var r0 error
	if rf, ok := ret.Get(0).(func(zapcore.Level) error); ok {
		r0 = rf(lvl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogSQL provides a mock function with given fields: logSQL
func (_m *GeneralConfig) SetLogSQL(logSQL bool) {
	_m.Called(logSQL)
}

// SetPasswords provides a mock function with given fields: keystore, vrf
func (_m *GeneralConfig) SetPasswords(keystore *string, vrf *string) {
	_m.Called(keystore, vrf)
}

// ShutdownGracePeriod provides a mock function with given fields:
func (_m *GeneralConfig) ShutdownGracePeriod() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// SolanaConfigs provides a mock function with given fields:
func (_m *GeneralConfig) SolanaConfigs() solana.TOMLConfigs {
	ret := _m.Called()

	var r0 solana.TOMLConfigs
	if rf, ok := ret.Get(0).(func() solana.TOMLConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(solana.TOMLConfigs)
		}
	}

	return r0
}

// SolanaEnabled provides a mock function with given fields:
func (_m *GeneralConfig) SolanaEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// StarkNetEnabled provides a mock function with given fields:
func (_m *GeneralConfig) StarkNetEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// StarknetConfigs provides a mock function with given fields:
func (_m *GeneralConfig) StarknetConfigs() starknet.StarknetConfigs {
	ret := _m.Called()

	var r0 starknet.StarknetConfigs
	if rf, ok := ret.Get(0).(func() starknet.StarknetConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(starknet.StarknetConfigs)
		}
	}

	return r0
}

// TelemetryIngress provides a mock function with given fields:
func (_m *GeneralConfig) TelemetryIngress() config.TelemetryIngress {
	ret := _m.Called()

	var r0 config.TelemetryIngress
	if rf, ok := ret.Get(0).(func() config.TelemetryIngress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.TelemetryIngress)
		}
	}

	return r0
}

// Threshold provides a mock function with given fields:
func (_m *GeneralConfig) Threshold() config.Threshold {
	ret := _m.Called()

	var r0 config.Threshold
	if rf, ok := ret.Get(0).(func() config.Threshold); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Threshold)
		}
	}

	return r0
}

// Tracing provides a mock function with given fields:
func (_m *GeneralConfig) Tracing() config.Tracing {
	ret := _m.Called()

	var r0 config.Tracing
	if rf, ok := ret.Get(0).(func() config.Tracing); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Tracing)
		}
	}

	return r0
}

// Validate provides a mock function with given fields:
func (_m *GeneralConfig) Validate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateDB provides a mock function with given fields:
func (_m *GeneralConfig) ValidateDB() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebServer provides a mock function with given fields:
func (_m *GeneralConfig) WebServer() config.WebServer {
	ret := _m.Called()

	var r0 config.WebServer
	if rf, ok := ret.Get(0).(func() config.WebServer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.WebServer)
		}
	}

	return r0
}

// NewGeneralConfig creates a new instance of GeneralConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGeneralConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *GeneralConfig {
	mock := &GeneralConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
