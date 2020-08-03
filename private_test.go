package bexs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var privateTest = New("staging", "", "", true)

func TestBexs_Config(t *testing.T) {
	_, err := privateTest.Config()
	assert.NoError(t, err)

}

func TestBexs_GetBalances(t *testing.T) {
	_, err := privateTest.GetBalances()
	assert.NoError(t, err)

}

func TestBexs_GetBalance(t *testing.T) {
	_, err := privateTest.GetBalance("BTC")
	assert.NoError(t, err)
}

func TestBexs_BuyMarket(t *testing.T) {
	_, err := privateTest.BuyMarket("BTC_USDT", 1)
	assert.NoError(t, err)
}

func TestBexs_SellMarket(t *testing.T) {
	_, err := privateTest.SellMarket("BTC_USDT", 1)
	assert.NoError(t, err)

}

func TestBexs_BuyLimit(t *testing.T) {
	_, err := privateTest.BuyLimit("BTC_USDT", 10000, 1)
	assert.NoError(t, err)
}

func TestBexs_SellLimit(t *testing.T) {
	_, err := privateTest.SellLimit("BTC_USDT", 10000, 1)
	assert.NoError(t, err)
}

func TestBexs_BuyAmi(t *testing.T) {
	_, err := privateTest.BuyAmi("BTC_USDT", 10000, 1, 11000)
	assert.NoError(t, err)

}

func TestBexs_SellAmi(t *testing.T) {
	_, err := privateTest.SellAmi("BTC_USDT", 13000, 1, 12000)
	assert.NoError(t, err)
}

func TestBexs_CancelOrder(t *testing.T) {
	order, err := privateTest.SellLimit("BTC_USDT", 50000, 1)
	assert.NoError(t, err)

	_, err = privateTest.CancelOrder(order)
	assert.NoError(t, err)

}

func TestBexs_GetOpenOrders(t *testing.T) {
	_, err := privateTest.GetOpenOrders("BTC_USDT")
	assert.NoError(t, err)
}

func TestBexs_OrderStatus(t *testing.T) {
	_, err := privateTest.OrderStatus("317550")
	assert.NoError(t, err)
}

func TestBexs_GetDepositAddress(t *testing.T) {
	_, err := privateTest.GetDepositAddress("BTC")
	assert.Error(t, err, "Error: No addresses available at this time.")
}

func TestBexs_GetDepositHistory(t *testing.T) {
	_, err := privateTest.GetDepositHistory()
	assert.NoError(t, err)
}

func TestBexs_GetWithdrawHistory(t *testing.T) {
	_, err := privateTest.GetWithdrawHistory()
	assert.NoError(t, err)
}

func TestBexs_Withdraw(t *testing.T) {
	_, err := privateTest.Withdraw("BTC", "abcde12346", 1)
	assert.Error(t, err, "Your Daily/Monthly limit has been exceeded. You need to submit the necessary documentation to increase yours limits.")
}

func TestBexs_DirectTransfer(t *testing.T) {
	_, err := privateTest.DirectTransfer("BTC", "teste@localhost", 1, 5)
	assert.Error(t, err, "Your Daily/Monthly limit has been exceeded. You need to submit the necessary documentation to increase yours limits.")
}
