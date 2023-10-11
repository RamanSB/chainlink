package evm_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	evmdb "github.com/smartcontractkit/chainlink/v2/core/chains/evm/db"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/mocks"
	configtest "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/configtest/v2"
	evmtestdb "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/evmtest/db"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestLegacyChains(t *testing.T) {
	evmCfg := configtest.NewGeneralConfig(t, nil)

	c := mocks.NewChain(t)
	c.On("ID").Return(big.NewInt(7))
	m := map[string]evm.Chain{c.ID().String(): c}

	l := evm.NewLegacyChains(m, evmCfg.EVMConfigs())
	assert.NotNil(t, l.ChainNodeConfigs())
	got, err := l.Get(c.ID().String())
	assert.NoError(t, err)
	assert.Equal(t, c, got)

}

func TestChainOpts_Validate(t *testing.T) {
	type fields struct {
		AppConfig        evm.AppConfig
		EventBroadcaster pg.EventBroadcaster
		MailMon          *utils.MailboxMonitor
		DB               *evmdb.ScopedDB
	}
	cfg := configtest.NewTestGeneralConfig(t)
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				AppConfig:        cfg,
				EventBroadcaster: pg.NewNullEventBroadcaster(),
				MailMon:          &utils.MailboxMonitor{},
				DB:               evmtestdb.NewScopedDB(t, cfg.Database()),
			},
		},
		{
			name: "invalid",
			fields: fields{
				AppConfig:        nil,
				EventBroadcaster: nil,
				MailMon:          nil,
				DB:               nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := evm.ChainOpts{
				AppConfig:        tt.fields.AppConfig,
				EventBroadcaster: tt.fields.EventBroadcaster,
				MailMon:          tt.fields.MailMon,
				DB:               tt.fields.DB,
			}
			if err := o.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ChainOpts.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
