package bexs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var publicTest = New("staging", "", "", true)

func TestBexs_GetAssets(t *testing.T) {
	_, err := publicTest.GetAssets()
	assert.NoError(t, err)
}

func TestBexs_GetMarkets(t *testing.T) {
	_, err := publicTest.GetMarkets()
	assert.NoError(t, err)
}

func TestBexs_GetTicker(t *testing.T) {
	_, err := publicTest.GetTicker("BTC_USDT")
	assert.NoError(t, err)
}

func TestBexs_GetMarketSummary(t *testing.T) {
	_, err := publicTest.GetMarketSummary("BTC_USDT")
	assert.NoError(t, err)
}

func TestBexs_GetOrderBook(t *testing.T) {
	_, err := publicTest.GetOrderBook("BTC_USDT", "ALL", 25)
	assert.NoError(t, err)
}

func TestBexs_GetMarketHistory(t *testing.T) {
	_, err := publicTest.GetMarketHistory("BTC_USDT", 25)
	assert.NoError(t, err)
}

func TestBexs_GetCandles(t *testing.T) {
	_, err := publicTest.GetCandles("BTC_USDT", "1d")
	assert.NoError(t, err)
}
