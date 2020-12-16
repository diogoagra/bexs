package bexs

import "github.com/shopspring/decimal"

// DefaultReturn -
type DefaultReturn struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// ConfigStruct -
type ConfigStruct struct {
	PercentDiscountOnTrade    int             `json:"PercentDiscountOnTrade"`
	PercentDiscountOnWithdraw int             `json:"PercentDiscountOnWithdraw"`
	EffectiveMakerFee         decimal.Decimal `json:"EffectiveMakerFee"`
	EffectiveTakerFee         decimal.Decimal `json:"EffectiveTakerFee"`
}

// GetBalanceStruct -
type GetBalanceStruct struct {
	Asset          string          `json:"Asset"`
	AssetName      string          `json:"AssetName"`
	Balance        decimal.Decimal `json:"Balance"`
	Available      decimal.Decimal `json:"Available"`
	Pending        decimal.Decimal `json:"Pending"`
	IsActive       bool            `json:"IsActive"`
	AllowDeposit   bool            `json:"AllowDeposit"`
	AllowWithdraw  bool            `json:"AllowWithdraw"`
	DepositAddress string          `json:"DepositAddress"`
	InfoMessage    string          `json:"InfoMessage"`
	BtcEquivalent  decimal.Decimal `json:"BtcEquivalent"`
	Currency       string          `json:"Currency"`
	CurrencyName   string          `json:"CurrencyName"`
	CryptoAddress  string          `json:"CryptoAddress"`
}

// GetAssetsStruct -
type GetAssetsStruct struct {
	Asset                string          `json:"Asset"`
	AssetLong            string          `json:"AssetLong"`
	MinConfirmation      int             `json:"MinConfirmation"`
	WithdrawTxFee        decimal.Decimal `json:"WithdrawTxFee"`
	WithdrawTxFeePercent decimal.Decimal `json:"WithdrawTxFeePercent"`
	SystemProtocol       string          `json:"SystemProtocol"`
	IsActive             bool            `json:"IsActive"`
	InfoMessage          string          `json:"InfoMessage"`
	MaintenanceMode      bool            `json:"MaintenanceMode"`
	MaintenanceMessage   string          `json:"MaintenanceMessage"`
	FormatPrefix         string          `json:"FormatPrefix"`
	FormatSufix          string          `json:"FormatSufix"`
	DecimalSeparator     string          `json:"DecimalSeparator"`
	ThousandSeparator    string          `json:"ThousandSeparator"`
	DecimalPlaces        int             `json:"DecimalPlaces"`
	Currency             string          `json:"Currency"`
	CurrencyLong         string          `json:"CurrencyLong"`
	CoinType             string          `json:"CoinType"`
}

// GetOrdersStruct -
type GetOrdersStruct struct {
	OrderID            int             `json:"OrderID"`
	Exchange           string          `json:"Exchange"`
	Type               string          `json:"Type"`
	Quantity           decimal.Decimal `json:"Quantity"`
	QuantityRemaining  decimal.Decimal `json:"QuantityRemaining"`
	QuantityBaseTraded decimal.Decimal `json:"QuantityBaseTraded"`
	Price              decimal.Decimal `json:"Price"`
	Status             string          `json:"Status"`
	Created            string          `json:"Created"`
	Comments           interface{}     `json:"Comments"`
}

// GetDepositAddressStruct -
type GetDepositAddressStruct struct {
	Asset          string `json:"Asset"`
	AssetName      string `json:"AssetName"`
	DepositAddress string `json:"DepositAddress"`
	Currency       string `json:"Currency"`
	CurrencyName   string `json:"CurrencyName"`
}

// GetHistoryStruct -
type GetHistoryStruct struct {
	ID            int             `json:"ID"`
	Timestamp     string          `json:"Timestamp"`
	Asset         string          `json:"Asset"`
	Amount        decimal.Decimal `json:"Amount"`
	TransactionID string          `json:"TransactionID"`
	Status        string          `json:"Status"`
	Label         string          `json:"Label"`
	Symbol        string          `json:"Symbol"`
}

// GetMarketsStruct -
type GetMarketsStruct struct {
	MarketName         string          `json:"MarketName"`
	MarketAsset        string          `json:"MarketAsset"`
	BaseAsset          string          `json:"BaseAsset"`
	MarketAssetLong    string          `json:"MarketAssetLong"`
	BaseAssetLong      string          `json:"BaseAssetLong"`
	IsActive           bool            `json:"IsActive"`
	MinTradeSize       decimal.Decimal `json:"MinTradeSize"`
	InfoMessage        string          `json:"InfoMessage"`
	MarketCurrency     string          `json:"MarketCurrency"`
	BaseCurrency       string          `json:"BaseCurrency"`
	MarketCurrencyLong string          `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string          `json:"BaseCurrencyLong"`
}

// GetTickerStruct -
type GetTickerStruct struct {
	Market string          `json:"Market"`
	Bid    decimal.Decimal `json:"Bid"`
	Ask    decimal.Decimal `json:"Ask"`
	Last   decimal.Decimal `json:"Last"`
}

// GetMarketSummarieStruct -
type GetMarketSummarieStruct struct {
	TimeStamp       string          `json:"TimeStamp"`
	MarketName      string          `json:"MarketName"`
	MarketAsset     string          `json:"MarketAsset"`
	BaseAsset       string          `json:"BaseAsset"`
	MarketAssetName string          `json:"MarketAssetName"`
	BaseAssetName   string          `json:"BaseAssetName"`
	PrevDay         decimal.Decimal `json:"PrevDay"`
	High            decimal.Decimal `json:"High"`
	Low             decimal.Decimal `json:"Low"`
	Last            decimal.Decimal `json:"Last"`
	Average         decimal.Decimal `json:"Average"`
	Volume          decimal.Decimal `json:"Volume"`
	BaseVolume      decimal.Decimal `json:"BaseVolume"`
	Bid             decimal.Decimal `json:"Bid"`
	Ask             decimal.Decimal `json:"Ask"`
	IsActive        bool            `json:"IsActive"`
	InfoMessage     string          `json:"InfoMessage"`
	MarketCurrency  string          `json:"MarketCurrency"`
	BaseCurrency    string          `json:"BaseCurrency"`
}

// Book -
type Book struct {
	Quantity decimal.Decimal `json:"Quantity"`
	Rate     decimal.Decimal `json:"Rate"`
}

// OrderBookStruct -
type OrderBookStruct struct {
	Buy  []Book `json:"buy"`
	Sell []Book `json:"sell"`
}

// GetMarketHistoryStruct -
type GetMarketHistoryStruct struct {
	TradeID    string          `json:"TradeID"`
	TimeStamp  string          `json:"TimeStamp"`
	OrderType  string          `json:"OrderType"`
	Price      decimal.Decimal `json:"Price"`
	Quantity   decimal.Decimal `json:"Quantity"`
	BaseVolume decimal.Decimal `json:"BaseVolume"`
	Total      decimal.Decimal `json:"Total"`
}

// GetCandlesStruct -
type GetCandlesStruct struct {
	TimeStamp  string          `json:"TimeStamp"`
	Open       decimal.Decimal `json:"Open"`
	High       decimal.Decimal `json:"High"`
	Low        decimal.Decimal `json:"Low"`
	Close      decimal.Decimal `json:"Close"`
	Volume     decimal.Decimal `json:"Volume"`
	BaseVolume decimal.Decimal `json:"BaseVolume"`
}

// TransactionStructs -
type TransactionStructs struct {
	ID          int     `json:"ID"`
	TimeStamp   string  `json:"TimeStamp"`
	Asset       string  `json:"Asset"`
	AssetName   string  `json:"AssetName"`
	Amount      float64 `json:"Amount"`
	Type        string  `json:"Type"`
	Description string  `json:"Description"`
	Comments    string  `json:"Comments"`
	CoinSymbol  string  `json:"CoinSymbol"`
	CoinName    string  `json:"CoinName"`
}
