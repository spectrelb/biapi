package okex

import (
	"github.com/okex/V3-Open-API-SDK/okex-go-sdk-api"
)

/**
获取账户中所有资产余额
GET /api/v5/asset/balance

获取账户中BTC、ETH两种资产余额
GET /api/v5/asset/balance?ccy=BTC,ETH
*/
func (client *Client) GetV5AssetBalance(ccy *string) (*[]map[string]interface{}, error) {
	uri := ASSET_BALANCE
	if ccy != nil && len(*ccy) > 0 {
		params := okex.NewParams()
		params["ccy"] = *ccy
		uri =  okex.BuildParams(uri, params)
	}

	r := []map[string]interface{}{}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}


/**
获取账户中所有资产余额
GET /api/v5/account/balance

获取账户中BTC、ETH两种资产余额
GET /api/v5/account/balance?ccy=BTC,ETH
*/
func (client *Client) GetV5AccountBalance(ccy *string) (*V5AccountBalanceResult, error) {
	uri := ACCOUNT_BALANCE
	if ccy != nil && len(*ccy) > 0 {
		params :=  okex.NewParams()
		params["ccy"] = *ccy
		uri =  okex.BuildParams(uri, params)
	}

	r := V5AccountBalanceResult{}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/**
获取交易产品基础信息
获取所有可交易产品的信息列表。
GET /api/v5/public/instruments?instType=SPOT

instType  产品类型
SPOT：币币
MARGIN：币币杠杆
SWAP：永续合约
FUTURES：交割合约
OPTION：期权
*/
func (client *Client) GetV5PublicInstruments(instType *string) (*V5PublicInstrumentsResult, error) {
	uri := PUBLICINSTRUMENTS
	if instType != nil && len(*instType) > 0 {
		params :=  okex.NewParams()
		params["instType"] = *instType
		uri =  okex.BuildParams(uri, params)
	}

	r := V5PublicInstrumentsResult{}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/**
获取单个产品行情信息
GET /api/v5/market/ticker?instId=BTC-USD-SWAP
*/
func (client *Client) GetV5PMarketTicker(instId *string) (*V5PMarketTickerResult, error) {
	uri := MARKETTICKER
	if instId != nil && len(*instId) > 0 {
		params :=  okex.NewParams()
		params["instId"] = *instId
		uri =  okex.BuildParams(uri, params)
	}

	r := V5PMarketTickerResult{}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/**
下单
POST /api/v5/trade/order
body
{
    "instId":"BTC-USDT",
    "tdMode":"cash",
    "clOrdId":"b15",
    "side":"buy",
    "ordType":"limit",
    "px":"2.15",
    "sz":"2"
}
*/
func (client *Client) PostV5TradeOrder(params *map[string]string) (*V5TradeOrderResult, error) {
	or := V5TradeOrderResult{}

	if _, err := client.Request(okex.POST, TRADEORDER, params, &or); err != nil {
		return nil, err
	}
	return &or, nil
}

/**
批量下单
POST /api/v5/trade/batch-order
body
[{
    "instId":"BTC-USDT",
    "tdMode":"cash",
    "clOrdId":"b15",
    "side":"buy",
    "ordType":"limit",
    "px":"2.15",
    "sz":"2"
},{
    "instId":"BTC-USDT",
    "tdMode":"cash",
    "clOrdId":"b15",
    "side":"buy",
    "ordType":"limit",
    "px":"2.15",
    "sz":"2"
}
]
*/
func (client *Client) PostV5TradeBatchOrder(params *[]map[string]string) (*V5TradeOrderResult, error) {
	or := V5TradeOrderResult{}

	if _, err := client.Request(okex.POST, TRADEBATCHORDER, params, &or); err != nil {
		return nil, err
	}
	return &or, nil
}

func (client *Client) GetV5TradeOrder(options *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := TRADEORDER
	if options != nil && len(*options) > 0 {
		params :=  okex.NewParams()
		params["instId"] = (*options)["instId"]
		params["ordId"] =  (*options)["ordId"]
		uri =  okex.BuildParams(uri, params)
	}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (client *Client) PostV5TradeOrderAlgo(optionalOrderInfo *map[string]string) (*V5TradeOrderAlgoResult, error) {
	or := V5TradeOrderAlgoResult{}

	if _, err := client.Request(okex.POST, TRADEORDERALGO, optionalOrderInfo, &or); err != nil {
		return nil, err
	}
	return &or, nil
}

func (client *Client) GetV5OrderHistory(options *map[string]string) (*V5OrderHistory, error) {
	r := V5OrderHistory{}

	uri := ORDERHISTORY
	if options != nil && len(*options) > 0 {
		params :=  okex.NewParams()
		params["instType"] = (*options)["instType"]
		params["limit"] = (*options)["limit"]
		uri =  okex.BuildParams(uri, params)
	}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}


func (client *Client) GetV5OrderPending(options *map[string]string) (*V5OrderHistory, error) {
	r := V5OrderHistory{}

	uri := ORDERHISTORYPENDING
	if options != nil && len(*options) > 0 {
		params :=  okex.NewParams()
		params["instType"] = (*options)["instType"]
		params["limit"] = (*options)["limit"]
		uri =  okex.BuildParams(uri, params)
	}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (client *Client) GetV5MarketCandles(options *map[string]string) (*GetV5MarketCandlesResult, error) {
	r := GetV5MarketCandlesResult{}

	uri := MARKETCANDLES
	if options != nil && len(*options) > 0 {
		params :=  okex.NewParams()
		params["instId"] = (*options)["instId"]
		params["limit"] = (*options)["limit"]
		params["bar"] = (*options)["bar"]
		params["after"] = (*options)["after"]
		uri =  okex.BuildParams(uri, params)
	}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (client *Client) GetV5MarketHistoryCandles(options *map[string]string) (*GetV5MarketCandlesResult, error) {
	r := GetV5MarketCandlesResult{}

	uri := MARKETHISTORYCANDLES
	if options != nil && len(*options) > 0 {
		params :=  okex.NewParams()
		params["instId"] = (*options)["instId"]
		params["limit"] = (*options)["limit"]
		params["bar"] = (*options)["bar"]
		params["after"] = (*options)["after"]
		uri =  okex.BuildParams(uri, params)
	}

	if _, err := client.Request(okex.GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}