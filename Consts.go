package biapi

type KlinePeriod int

//k线周期
const (
	KLINE_PERIOD_5MIN = 1 + iota
	KLINE_PERIOD_15MIN
	KLINE_PERIOD_30MIN
	KLINE_PERIOD_60MIN
	KLINE_PERIOD_1H
	KLINE_PERIOD_4H
	KLINE_PERIOD_1DAY
	KLINE_PERIOD_1WEEK
	KLINE_PERIOD_1MONTH
)

var HuobiKlinePeriodConverter = map[KlinePeriod]string{
	KLINE_PERIOD_5MIN:   "5min",
	KLINE_PERIOD_15MIN:  "15min",
	KLINE_PERIOD_30MIN:  "30min",
	KLINE_PERIOD_60MIN:  "1hour",
	KLINE_PERIOD_1H:     "1hour",
	KLINE_PERIOD_4H:     "4hour",
	KLINE_PERIOD_1DAY:   "1day",
	KLINE_PERIOD_1WEEK:  "7day",
	KLINE_PERIOD_1MONTH: "30day",
}

var Stablecoins = []string{"USDT", "USDC", "BUSD", "DAI", "UST", "TUSD", "PAX", "LUSD", "HUSD", "USDN", "GUSD", "FRAX", "ALUSD", "SUSD", "USDP", "USDX"}