package store_test

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/internal/cltest"
	strpkg "github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/models"
	"github.com/smartcontractkit/chainlink/utils"
	"github.com/stretchr/testify/assert"
)

func TestTxManager_CreateTx_Success(t *testing.T) {
	t.Parallel()
	app, cleanup := cltest.NewApplicationWithKeyStore()
	defer cleanup()
	store := app.Store
	manager := store.TxManager

	to := cltest.NewAddress()
	data, err := hex.DecodeString("0000abcdef")
	assert.NoError(t, err)
	hash := cltest.NewHash()
	sentAt := uint64(23456)
	nonce := uint64(256)
	ethMock := app.MockEthClient()
	ethMock.Context("app.Start()", func(ethMock *cltest.EthMock) {
		ethMock.Register("eth_getTransactionCount", utils.Uint64ToHex(nonce))
	})
	assert.NoError(t, app.Start())

	ethMock.Context("manager.CreateTx#1", func(ethMock *cltest.EthMock) {
		ethMock.Register("eth_sendRawTransaction", hash)
		ethMock.Register("eth_blockNumber", utils.Uint64ToHex(sentAt))
	})

	a, err := manager.CreateTx(to, data)
	assert.NoError(t, err)
	tx := models.Tx{}
	assert.NoError(t, store.One("ID", a.TxID, &tx))
	assert.NoError(t, err)
	assert.Equal(t, nonce, tx.Nonce)
	assert.Equal(t, data, tx.Data)
	assert.Equal(t, to, tx.To)

	assert.NoError(t, store.One("From", tx.From, &tx))
	assert.Equal(t, nonce, tx.Nonce)
	attempts, err := store.AttemptsFor(tx.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(attempts))

	ethMock.EventuallyAllCalled(t)
}

func TestTxManager_CreateTx_AttemptErrorDeletesTxAndDoesNotIncrementNonce(t *testing.T) {
	t.Parallel()

	config, configCleanup := cltest.NewConfig()
	defer configCleanup()

	app, cleanup := cltest.NewApplicationWithConfigAndKeyStore(config)
	defer cleanup()

	store := app.Store
	manager := store.TxManager

	to := cltest.NewAddress()
	data, err := hex.DecodeString("0000abcdef")
	assert.NoError(t, err)
	sentAt := uint64(23456)
	nonce := uint64(256)
	ethMock := app.MockEthClient()
	ethMock.Context("app.Start()", func(ethMock *cltest.EthMock) {
		ethMock.Register("eth_getTransactionCount", utils.Uint64ToHex(nonce))
	})
	assert.NoError(t, app.Start())

	ethMock.Context("manager.CreateTx#1", func(ethMock *cltest.EthMock) {
		ethMock.Register("eth_sendRawTransaction", "invalid")
		ethMock.Register("eth_blockNumber", utils.Uint64ToHex(sentAt))
	})

	_, err = manager.CreateTx(to, data)
	assert.Error(t, err)

	var txs []models.Tx
	err = store.ORM.All(&txs)
	assert.Equal(t, 0, len(txs))

	var txAttempts []models.TxAttempt
	err = store.ORM.All(&txAttempts)
	assert.Equal(t, 0, len(txAttempts))

	hash := cltest.NewHash()
	ethMock.Context("manager.CreateTx#2", func(ethMock *cltest.EthMock) {
		ethMock.Register("eth_sendRawTransaction", hash)
		ethMock.Register("eth_blockNumber", utils.Uint64ToHex(sentAt))
	})

	a, err := manager.CreateTx(to, data)
	assert.NoError(t, err)
	tx := models.Tx{}
	assert.NoError(t, store.One("ID", a.TxID, &tx))
	assert.NoError(t, err)
	assert.Equal(t, nonce, tx.Nonce)
	assert.Equal(t, data, tx.Data)
	assert.Equal(t, to, tx.To)

	assert.NoError(t, store.One("From", tx.From, &tx))
	assert.Equal(t, nonce, tx.Nonce)
	attempts, err := store.AttemptsFor(tx.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(attempts))

	ethMock.EventuallyAllCalled(t)
}

func TestTxManager_MeetsMinConfirmations_beforeGasBumpThreshold(t *testing.T) {
	t.Parallel()

	app, cleanup := cltest.NewApplicationWithKeyStore()
	defer cleanup()
	store := app.Store
	config := store.Config
	txm := store.TxManager
	ethMock := app.MockEthClient()

	sentAt := uint64(23456)
	from := cltest.GetAccountAddress(store)

	tx := cltest.CreateTxAndAttempt(store, from, sentAt)
	attempts, err := store.AttemptsFor(tx.ID)
	assert.NoError(t, err)
	a := attempts[0]

	currentBlock := sentAt + config.EthGasBumpThreshold - 1
	ethMock.Context("txm.MeetsMinConfirmations", func(ethMock *cltest.EthMock) {
		ethMock.Register("eth_getTransactionReceipt", strpkg.TxReceipt{})
		ethMock.Register("eth_blockNumber", utils.Uint64ToHex(currentBlock))
	})

	confirmed, err := txm.MeetsMinConfirmations(a.Hash)
	assert.NoError(t, err)
	assert.False(t, confirmed)
	assert.NoError(t, store.One("ID", tx.ID, tx))
	attempts, err = store.AttemptsFor(tx.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(attempts))

	ethMock.EventuallyAllCalled(t)
}

func TestTxManager_MeetsMinConfirmations_AtThreshold(t *testing.T) {
	t.Parallel()
	app, cleanup := cltest.NewApplicationWithKeyStore()
	defer cleanup()
	store := app.Store
	config := store.Config
	txm := store.TxManager
	ethMock := app.MockEthClient()

	sentAt := uint64(23456)
	from := cltest.GetAccountAddress(store)

	tx := cltest.CreateTxAndAttempt(store, from, sentAt)
	attempts, err := store.AttemptsFor(tx.ID)
	assert.NoError(t, err)
	a := attempts[0]
	ethMock.Context("txm.MeetsMinConfirmations", func(ethMock *cltest.EthMock) {
		ethMock.Register("eth_getTransactionReceipt", strpkg.TxReceipt{})
		ethMock.Register("eth_blockNumber", utils.Uint64ToHex(sentAt+config.EthGasBumpThreshold))
		ethMock.Register("eth_sendRawTransaction", cltest.NewHash())
	})

	confirmed, err := txm.MeetsMinConfirmations(a.Hash)
	assert.NoError(t, err)
	assert.False(t, confirmed)
	assert.NoError(t, store.One("ID", tx.ID, tx))
	attempts, err = store.AttemptsFor(tx.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(attempts))

	ethMock.EventuallyAllCalled(t)
}

func TestTxManager_MeetsMinConfirmations_confirmed(t *testing.T) {
	t.Parallel()

	config, configCleanup := cltest.NewConfig()
	defer configCleanup()

	sentAt := uint64(1)
	receiptAt := uint64(2)
	config.MinOutgoingConfirmations = 2

	app, cleanup := cltest.NewApplicationWithConfigAndKeyStore(config)
	defer cleanup()
	store := app.Store
	txm := store.TxManager

	from := cltest.GetAccountAddress(store)

	tests := []struct {
		name          string
		currentHeight uint64
		want          bool
	}{
		{"less than min confs", 2, false},
		{"equal min confs", 3, true},
		{"1 greater than min confs", 4, true},
		{"2 greater than min confs", 5, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ethMock := app.MockEthClient()
			confirmationReceipt := strpkg.TxReceipt{
				Hash:        cltest.NewHash(),
				BlockNumber: cltest.BigHexInt(receiptAt),
			}
			tx := cltest.CreateTxAndAttempt(store, from, sentAt)
			a := tx.TxAttempt
			ethMock.Context("txm.MeetsMinConfirmations", func(ethMock *cltest.EthMock) {
				ethMock.Register("eth_getTransactionReceipt", confirmationReceipt)
				ethMock.Register("eth_blockNumber", utils.Uint64ToHex(test.currentHeight))
			})

			actual, err := txm.MeetsMinConfirmations(a.Hash)
			assert.NoError(t, err)
			assert.Equal(t, test.want, actual)

			attempts, err := store.AttemptsFor(tx.ID)
			assert.NoError(t, err)
			assert.Equal(t, 1, len(attempts))

			ethMock.EventuallyAllCalled(t)
		})
	}
}

func TestTxManager_ActivateAccount(t *testing.T) {
	t.Parallel()

	ethMock := &cltest.EthMock{}
	txm := &strpkg.TxManager{
		EthClient: &strpkg.EthClient{CallerSubscriber: ethMock},
	}
	account := accounts.Account{Address: common.HexToAddress("0xbf4ed7b27f1d666546e30d74d50d173d20bca754")}

	ethMock.Register("eth_getTransactionCount", `0x2D0`)
	assert.NoError(t, txm.ActivateAccount(account))
	ethMock.EventuallyAllCalled(t)

	aa := txm.GetActiveAccount()
	assert.Equal(t, account.Address, aa.Address)
	assert.Equal(t, uint64(0x2d0), aa.GetNonce())
}

func TestActiveAccount_GetAndIncrementNonce_YieldsCurrentNonceAndIncrements(t *testing.T) {
	account := accounts.Account{Address: common.HexToAddress("0xbf4ed7b27f1d666546e30d74d50d173d20bca754")}
	activeAccount := strpkg.ActiveAccount{
		Account: account,
	}

	activeAccount.GetAndIncrementNonce(func(y uint64) error {
		assert.Equal(t, uint64(0), y)
		return nil
	})
	assert.Equal(t, uint64(1), activeAccount.GetNonce())

	activeAccount.GetAndIncrementNonce(func(y uint64) error {
		assert.Equal(t, uint64(1), y)
		return nil
	})
	assert.Equal(t, uint64(2), activeAccount.GetNonce())
}

func TestActiveAccount_GetAndIncrementNonce_DoesNotIncrementWhenCallbackThrowsException(t *testing.T) {
	account := accounts.Account{Address: common.HexToAddress("0xbf4ed7b27f1d666546e30d74d50d173d20bca754")}
	activeAccount := strpkg.ActiveAccount{
		Account: account,
	}

	err := activeAccount.GetAndIncrementNonce(func(y uint64) error {
		assert.Equal(t, uint64(0), y)
		return errors.New("Should not increment")
	})
	assert.Error(t, err)
	err = activeAccount.GetAndIncrementNonce(func(y uint64) error {
		assert.Equal(t, uint64(0), y)
		return errors.New("Should not increment again")
	})
	assert.Error(t, err)
	assert.Equal(t, uint64(0), activeAccount.GetNonce())
}
