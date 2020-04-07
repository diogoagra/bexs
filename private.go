package bexs

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Config -
func (c *Bexs) Config() (result ConfigStruct, err error) {
	response, err := c.GetURL("/api/v3/private/config?", true)
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

// GetBalances -
func (c *Bexs) GetBalances() (result []GetBalanceStruct, err error) {
	response, err := c.GetURL("/api/v3/private/getbalances?", true)
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

// GetBalance -
func (c *Bexs) GetBalance(asset string) (result GetBalanceStruct, err error) {
	if asset == "" {
		err = fmt.Errorf("Invalid input asset %s", asset)
		return
	}
	var tmp []GetBalanceStruct
	response, err := c.GetURL("/api/v3/private/getbalance?"+fmt.Sprintf("asset=%s", asset), true)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &tmp)
	if err != nil {
		return
	}
	result = tmp[0]
	return
}

// BuyMarket - Send a buy market order
func (c *Bexs) BuyMarket(market string, quantity float64) (result string, err error) {
	if market == "" || quantity <= 0 {
		err = fmt.Errorf("Invalid input market %s quantity %f", market, quantity)
		return
	}
	response, err := c.GetURL("/api/v3/private/buymarket?"+fmt.Sprintf("market=%s&quantity=%.8f", market, quantity), true)

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

// SellMarket - Send a sell market order
func (c *Bexs) SellMarket(market string, quantity float64) (result string, err error) {
	if market == "" || quantity <= 0 {
		err = fmt.Errorf("Invalid input market %s quantity %f", market, quantity)
		return
	}
	response, err := c.GetURL("/api/v3/private/sellmarket?"+fmt.Sprintf("market=%s&quantity=%.8f", market, quantity), true)

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

// BuyLimit - Send a buy limit order
func (c *Bexs) BuyLimit(market string, rate, quantity float64) (result string, err error) {
	if market == "" || rate <= 0 || quantity <= 0 {
		err = fmt.Errorf("Invalid input market %s rate %f quantity %f", market, rate, quantity)
		return
	}
	response, err := c.GetURL("/api/v3/private/buylimit?"+fmt.Sprintf("market=%s&rate=%.8f&quantity=%.8f", market, rate, quantity), true)

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

// SellLimit - Send a sell limit order
func (c *Bexs) SellLimit(market string, rate, quantity float64) (result string, err error) {
	if market == "" || rate <= 0 || quantity <= 0 {
		err = fmt.Errorf("Invalid input market %s rate %f quantity %f", market, rate, quantity)
		return
	}
	response, err := c.GetURL("/api/v3/private/selllimit?"+fmt.Sprintf("market=%s&rate=%.8f&quantity=%.8f", market, rate, quantity), true)
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

// BuyAmi - Send ami buy order, must have amirate > rate
func (c *Bexs) BuyAmi(market string, rate, quantity, amirate float64) (result string, err error) {
	if market == "" || rate <= 0 || quantity <= 0 || amirate <= 0 {
		err = fmt.Errorf("Invalid input market %s rate %f quantity %f amirate %f", market, rate, quantity, amirate)
		return
	}
	if rate > amirate {
		err = fmt.Errorf("Invalid AMI rate %f amirate %f", rate, amirate)
		return
	}
	response, err := c.GetURL("/api/v3/private/buylimitami?"+fmt.Sprintf("market=%s&rate=%.8f&amirate=%.8f&quantity=%.8f", market, rate, amirate, quantity), true)
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

// SellAmi - Send ami sell order, must have rate > amirate
func (c *Bexs) SellAmi(market string, rate, quantity, amirate float64) (result string, err error) {
	if market == "" || rate <= 0 || quantity <= 0 || amirate <= 0 {
		err = fmt.Errorf("Invalid input market %s rate %f quantity %f ami rate %f", market, rate, quantity, amirate)
		return
	}
	if amirate > rate {
		err = fmt.Errorf("Invalid AMI rate %f amirate %f", rate, amirate)
		return
	}
	response, err := c.GetURL("/api/v3/private/selllimitami?"+fmt.Sprintf("market=%s&rate=%.8f&quantity=%.8f", market, rate, quantity), true)
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

// CancelOrder -
func (c *Bexs) CancelOrder(orderID string) (result string, err error) {
	if orderID == "" {
		err = errors.New("Invalid input empty orderid")
		return
	}
	response, err := c.GetURL("/api/v3/private/ordercancel?"+fmt.Sprintf("orderid=%s", orderID), true)
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
	result = orderID
	return
}

// CancelOrders -
func (c *Bexs) CancelOrders(ordersID []int) (result string, err error) {

	if len(ordersID) == 0 {
		err = errors.New("Invalid input empty orders")
		return
	}

	var orders string

	for _, a := range ordersID {
		orders += fmt.Sprintf("orderids=%v&", a)
	}

	response, err := c.GetURL("/api/v3/private/cancelorders?"+orders[0:len(orders)-1], true)
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
	result = "true" // TODO
	return
}

// GetOpenOrders -
func (c *Bexs) GetOpenOrders(market string) (result []GetOrdersStruct, err error) {
	if market == "" {
		err = fmt.Errorf("Invalid input empty market")
		return
	}
	response, err := c.GetURL("/api/v3/private/getopenorders?"+fmt.Sprintf("market=%s", market), true)
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

// OrderStatus -
func (c *Bexs) OrderStatus(orderID string) (result GetOrdersStruct, err error) {
	if orderID == "" {
		err = fmt.Errorf("Invalid input empty order")
		return
	}
	var tmp []GetOrdersStruct
	response, err := c.GetURL("/api/v3/private/getorder?"+fmt.Sprintf("orderid=%s", orderID), true)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &tmp)
	if err != nil {
		return
	}

	if len(tmp) > 0 {
		result = tmp[0]
	}

	return
}

// GetDepositAddress -
func (c *Bexs) GetDepositAddress(asset string) (result GetDepositAddressStruct, err error) {
	if asset == "" {
		err = fmt.Errorf("Invalid input empty asset")
		return
	}
	response, err := c.GetURL("/api/v3/private/getdepositaddress?"+fmt.Sprintf("asset=%s", asset), true)
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

// GetDepositHistory -
func (c *Bexs) GetDepositHistory() (result []GetHistoryStruct, err error) {
	response, err := c.GetURL("/api/v3/private/getdeposithistory?", true)
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

// GetWithdrawHistory -
func (c *Bexs) GetWithdrawHistory() (result []GetHistoryStruct, err error) {
	response, err := c.GetURL("/api/v3/private/getwithdrawhistory?", true)
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

// Withdraw -
func (c *Bexs) Withdraw(asset, address string, quantity float64) (result bool, err error) {
	if asset == "" || address == "" || quantity <= 0 {
		err = fmt.Errorf("Invalid input")
		return
	}
	response, err := c.GetURL("/api/v3/private/withdraw?"+fmt.Sprintf("asset=%s&address=%s&quantity=%.8f", asset, address, quantity), true)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	result = true
	return
}

// DirectTransfer -
func (c *Bexs) DirectTransfer(asset, to string, quantity float64, exchange int) (result bool, err error) {
	if asset == "" || to == "" || quantity <= 0 || exchange <= 0 {
		err = fmt.Errorf("Invalid input asset %s to %s quantity %f exchange %v", asset, to, quantity, exchange)
		return
	}
	response, err := c.GetURL("/api/v3/private/directtransfer?"+fmt.Sprintf("asset=%s&accountto=%s&quantity=%.8f&exchangeto=%v", asset, to, quantity, exchange), true)
	if err != nil {
		return
	}
	response, err = c.ParseResult(response)
	if err != nil {
		return
	}
	result = true
	return
}
