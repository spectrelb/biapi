package okex

import (
	. "github.com/spectrelb/biapi"
)

const (
	ASSET_BALANCE                  = "/api/v5/asset/balances"
	ACCOUNT_BALANCE                = "/api/v5/account/balance"
	PUBLICINSTRUMENTS              = "/api/v5/public/instruments"

	TRADEORDER                     = "/api/v5/trade/order"
	TRADEBATCHORDER                = "/api/v5/trade/batch-orders"
	TRADEORDERALGO                 = "/api/v5/trade/order-algo"
	ORDERHISTORY                   = "/api/v5/trade/orders-history"
	ORDERHISTORYPENDING            = "/api/v5/trade/orders-pending"

	MARKETTICKER                   = "/api/v5/market/ticker"
	MARKETCANDLES                  = "/api/v5/market/candles"
	MARKETHISTORYCANDLES           = "/api/v5/market/history-candles"

)

var OkexKlinePeriodConverter = map[KlinePeriod]string{
	KLINE_PERIOD_5MIN:   "5m",
	KLINE_PERIOD_15MIN:  "15m",
	KLINE_PERIOD_30MIN:  "30m",
	KLINE_PERIOD_1H:     "1H",
	KLINE_PERIOD_4H:     "4H",
	KLINE_PERIOD_1DAY:   "1D",
	KLINE_PERIOD_1WEEK:  "1W",
	KLINE_PERIOD_1MONTH: "1M",
}

const (
	SPOT        = "SPOT"
	MARGIN      = "MARGIN"
	SWAP        = "SWAP"
	FUTURES     = "FUTURES"
)