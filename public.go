package bexs

import (
	"encoding/json"
	"fmt"
)

// GetAssets -
func (c *Bexs) GetAssets() (result []GetAssetsStruct, err error) {
	response, err := c.GetURL("/api/v3/public/getassets?", false)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

// GetMarkets -
func (c *Bexs) GetMarkets() (result []GetMarketsStruct, err error) {
	response, err := c.GetURL("/api/v3/public/getmarkets?", false)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

// GetTicker -
func (c *Bexs) GetTicker(market string) (result []GetTickerStruct, err error) {
	if market == "" {
		err = fmt.Errorf("Invalid input empty market")
		return
	}
	response, err := c.GetURL("/api/v3/public/getticker?"+fmt.Sprintf("market=%s", market), false)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

// GetMarketSummary -
func (c *Bexs) GetMarketSummary(market string) (result GetMarketSummarieStruct, err error) {
	if market == "" {
		err = fmt.Errorf("Invalid input empty market")
		return
	}
	response, err := c.GetURL("/api/v3/public/getmarketsummary?"+fmt.Sprintf("market=%s", market), false)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

// GetOrderBook -
func (c *Bexs) GetOrderBook(market, kind string, depth int) (result OrderBookStruct, err error) {
	if market == "" {
		err = fmt.Errorf("Invalid input empty market")
		return
	}
	response, err := c.GetURL("/api/v3/public/getorderbook?"+fmt.Sprintf("market=%s&type=%s&depth=%v", market, kind, depth), false)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

// GetMarketHistory -
func (c *Bexs) GetMarketHistory(market string, count int) (result []GetMarketHistoryStruct, err error) {
	if market == "" {
		err = fmt.Errorf("Invalid input empty market")
		return
	}
	response, err := c.GetURL("/api/v3/public/getmarkethistory?"+fmt.Sprintf("market=%s&count=%v", market, count), false)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

// GetCandles -
func (c *Bexs) GetCandles(market, period string) (result []GetCandlesStruct, err error) {
	if market == "" {
		err = fmt.Errorf("Invalid input empty market")
		return
	}
	response, err := c.GetURL("/api/v3/public/getcandles?"+fmt.Sprintf("market=%s&period=%s", market, period), false)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}
