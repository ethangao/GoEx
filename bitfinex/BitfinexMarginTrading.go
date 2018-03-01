package bitfinex

import . "github.com/ethangao/GoEx"

type MarginLimits struct {
	Pair              string  `json:"on_pair"`
	InitialMargin     float64 `json:"initial_margin,string"`
	MarginRequirement float64 `json:"margin_requirement,string"`
	TradableBalance   float64 `json:"tradable_balance,string"`
}

type MarginInfo struct {
	MarginBalance     float64        `json:"margin_balance,string"`
	TradableBalance   float64        `json:"tradable_balance,string"`
	UnrealizedPl      float64        `json:"unrealized_pl,string"`
	UnrealizedSwap    float64        `json:"unrealized_swap,string"`
	NetValue          float64        `json:"net_value,string"`
	RequiredMargin    float64        `json:"required_margin,string"`
	Leverage          float64        `json:"leverage,string"`
	MarginRequirement float64        `json:"margin_requirement,string"`
	MarginLimits      []MarginLimits `json:"margin_limits"`
}

type MarginPositionInfo struct {
	Id        int64   `json:"id"`
	Symbol    string  `json:"symbol"`
	Status    string  `json:"status"`
	Base      float64 `json:"base,string"`
	Amount    float64 `json:"amount,string"`
	Timestamp string  `json:"timestamp"`
	Swap      float64 `json:"swap,string"`
	Pl        float64 `json:"pl,string"`
}

func (bfx *Bitfinex) GetMarginTradingWalletBalance() (*Account, error) {
	balancemap, err := bfx.GetWalletBalances()
	if err != nil {
		return nil, err
	}
	return balancemap["trading"], nil
}

func (bfx *Bitfinex) MarginLimitBuy(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("limit", "buy", amount, price, currencyPair)
}

func (bfx *Bitfinex) MarginLimitSell(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("limit", "sell", amount, price, currencyPair)
}

func (bfx *Bitfinex) MarginMarketBuy(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("Market", "buy", amount, price, currencyPair)
}

func (bfx *Bitfinex) MarginMarketSell(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("Market", "sell", amount, price, currencyPair)
}

func (bfx *Bitfinex) GetMarginInfos() ([]MarginInfo, error) {
	var marginInfo []MarginInfo
	err := bfx.doAuthenticatedRequest("POST", "margin_infos", map[string]interface{}{}, &marginInfo)
	if err != nil {
		return nil, err
	}
	return marginInfo, nil
}

func (bfx *Bitfinex) GetMarginPositions() ([]MarginPositionInfo, error) {
	var ret []MarginPositionInfo
	err := bfx.doAuthenticatedRequest("POST", "positions", map[string]interface{}{}, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (bfx *Bitfinex) ClaimMarginPosition(id int64, amount float64) (*MarginPositionInfo, error) {
	var ret MarginPositionInfo
	err := bfx.doAuthenticatedRequest("POST", "positions/claim",
		map[string]interface{}{"position_id": id, "amount": amount}, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
