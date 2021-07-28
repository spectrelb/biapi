package biapi


type TradeSide int

const (
	BUY TradeSide = 1 + iota
	SELL
	BUY_MARKET
	SELL_MARKET
)

func (ts TradeSide) String() string {
	switch ts {
	case 1:
		return "BUY"
	case 2:
		return "SELL"
	case 3:
		return "BUY_MARKET"
	case 4:
		return "SELL_MARKET"
	default:
		return "UNKNOWN"
	}
}

const (
	OPEN_BUY   = 1 + iota //开多
	OPEN_SELL             //开空
	CLOSE_BUY             //平多
	CLOSE_SELL            //平空
)

type KlinePeriod int

const (
	KLINE_PERIOD_1MIN = 1 + iota
	KLINE_PERIOD_3MIN
	KLINE_PERIOD_5MIN
	KLINE_PERIOD_15MIN
	KLINE_PERIOD_30MIN
	KLINE_PERIOD_60MIN
	KLINE_PERIOD_1H
	KLINE_PERIOD_2H
	KLINE_PERIOD_3H
	KLINE_PERIOD_4H
	KLINE_PERIOD_6H
	KLINE_PERIOD_8H
	KLINE_PERIOD_12H
	KLINE_PERIOD_1DAY
	KLINE_PERIOD_3DAY
	KLINE_PERIOD_1WEEK
	KLINE_PERIOD_1MONTH
	KLINE_PERIOD_1YEAR
)


var (
	THIS_WEEK_CONTRACT  = "this_week"  //周合约
	NEXT_WEEK_CONTRACT  = "next_week"  //次周合约
	QUARTER_CONTRACT    = "quarter"    //季度合约
	BI_QUARTER_CONTRACT = "bi_quarter" // NEXT QUARTER
	SWAP_CONTRACT       = "swap"       //永续合约
	SWAP_USDT_CONTRACT  = "swap-usdt"
)

//exchanges const
const (
	OKEX            = "okex.com"
	HUOBI           = "huobi.com"
	HUOBI_PRO       = "huobi.pro"
	BINANCE         = "binance.com"
)