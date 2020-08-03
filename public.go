package bexs

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetAssets -
func (c *Bexs) GetAssets() (result []GetAssetsStruct, err error) {
	response, err := c.getURL("/api/v3/public/getassets?", false)
	if err != nil {
		return
	}
	response, err = c.parseResult(response)
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
	response, err := c.getURL("/api/v3/public/getmarkets?", false)
	if err != nil {
		return
	}
	response, err = c.parseResult(response)
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

	params := make(url.Values)
	params.Add("market", market)

	response, err := c.getURL(fmt.Sprintf("/api/v3/public/getticker?%s", params.Encode()), false)
	if err != nil {
		return
	}
	response, err = c.parseResult(response)
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

	params := make(url.Values)
	params.Add("market", market)

	response, err := c.getURL(fmt.Sprintf("/api/v3/public/getmarketsummary?%s", params.Encode()), false)
	if err != nil {
		return
	}
	response, err = c.parseResult(response)
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

	params := make(url.Values)
	params.Add("market", market)
	params.Add("type", kind)
	params.Add("depth", fmt.Sprintf("%d", depth))

	response, err := c.getURL(fmt.Sprintf("/api/v3/public/getorderbook?%s", params.Encode()), false)
	if err != nil {
		return
	}
	response, err = c.parseResult(response)
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

	params := make(url.Values)
	params.Add("market", market)
	params.Add("count", fmt.Sprintf("%d", count))

	response, err := c.getURL(fmt.Sprintf("/api/v3/public/getmarkethistory?%s", params.Encode()), false)
	if err != nil {
		return
	}
	response, err = c.parseResult(response)
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

	params := make(url.Values)
	params.Add("market", market)
	params.Add("period", period)

	response, err := c.getURL(fmt.Sprintf("/api/v3/public/getcandles?%s", params.Encode()), false)
	if err != nil {
		return
	}
	response, err = c.parseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}
