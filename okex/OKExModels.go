package okex


type V5AccountBalanceResult struct {
	Code string `json:"code"`
	Data []Data `json:"data"`
	Msg string `json:"msg"`
}
type Details struct {
	AvailBal string `json:"availBal"`
	AvailEq string `json:"availEq"`
	CashBal string `json:"cashBal"`
	Ccy string `json:"ccy"`
	CrossLiab string `json:"crossLiab"`
	DisEq string `json:"disEq"`
	Eq string `json:"eq"`
	EqUsd string `json:"eqUsd"`
	FrozenBal string `json:"frozenBal"`
	Interest string `json:"interest"`
	IsoEq string `json:"isoEq"`
	IsoLiab string `json:"isoLiab"`
	Liab string `json:"liab"`
	MaxLoan string `json:"maxLoan"`
	MgnRatio string `json:"mgnRatio"`
	NotionalLever string `json:"notionalLever"`
	OrdFrozen string `json:"ordFrozen"`
	Twap string `json:"twap"`
	UTime string `json:"uTime"`
	Upl string `json:"upl"`
	UplLiab string `json:"uplLiab"`
}
type Data struct {
	AdjEq string `json:"adjEq"`
	Details []Details `json:"details"`
	Imr string `json:"imr"`
	IsoEq string `json:"isoEq"`
	MgnRatio string `json:"mgnRatio"`
	Mmr string `json:"mmr"`
	NotionalUsd string `json:"notionalUsd"`
	OrdFroz string `json:"ordFroz"`
	TotalEq string `json:"totalEq"`
	UTime string `json:"uTime"`
}


type V5PublicInstrumentsResult struct {
	Code string `json:"code"`
	Data []V5PublicInstrumentsData `json:"data"`
	Msg string `json:"msg"`
}
type V5PublicInstrumentsData struct {
	Alias string `json:"alias"`
	BaseCcy string `json:"baseCcy"`
	Category string `json:"category"`
	CtMult string `json:"ctMult"`
	CtType string `json:"ctType"`
	CtVal string `json:"ctVal"`
	CtValCcy string `json:"ctValCcy"`
	ExpTime string `json:"expTime"`
	InstID string `json:"instId"`
	InstType string `json:"instType"`
	Lever string `json:"lever"`
	ListTime string `json:"listTime"`
	LotSz string `json:"lotSz"`
	MinSz string `json:"minSz"`
	OptType string `json:"optType"`
	QuoteCcy string `json:"quoteCcy"`
	SettleCcy string `json:"settleCcy"`
	State string `json:"state"`
	Stk string `json:"stk"`
	TickSz string `json:"tickSz"`
	Uly string `json:"uly"`
}


type V5PMarketTickerResult struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data []V5PMarketTickerData `json:"data"`
}
type V5PMarketTickerData struct {
	InstType string `json:"instType"`
	InstID string `json:"instId"`
	Last string `json:"last"`
	LastSz string `json:"lastSz"`
	AskPx string `json:"askPx"`
	AskSz string `json:"askSz"`
	BidPx string `json:"bidPx"`
	BidSz string `json:"bidSz"`
	Open24H string `json:"open24h"`
	High24H string `json:"high24h"`
	Low24H string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H string `json:"vol24h"`
	Ts string `json:"ts"`
	SodUtc0 string `json:"sodUtc0"`
	SodUtc8 string `json:"sodUtc8"`
}



type V5TradeOrderResult struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data []V5TradeOrderData `json:"data"`
}

type V5TradeOrderData struct {
	ClOrdID string `json:"clOrdId"`
	OrdID string `json:"ordId"`
	Tag string `json:"tag"`
	SCode string `json:"sCode"`
	SMsg string `json:"sMsg"`
}

type V5TradeOrderAlgoResult struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data []V5TradeOrderAlgoData `json:"data"`
}

type V5TradeOrderAlgoData struct {
	AlgoId string `json:"algoId"`
	SCode string `json:"sCode"`
	SMsg string `json:"sMsg"`
}

type V5OrderHistory struct {
	Code string `json:"code"`
	Data []V5OrderHistoryData `json:"data"`
	Msg string `json:"msg"`
}
type V5OrderHistoryData struct {
	AccFillSz string `json:"accFillSz"`
	AvgPx string `json:"avgPx"`
	CTime string `json:"cTime"`
	Category string `json:"category"`
	Ccy string `json:"ccy"`
	ClOrdID string `json:"clOrdId"`
	Fee string `json:"fee"`
	FeeCcy string `json:"feeCcy"`
	FillPx string `json:"fillPx"`
	FillSz string `json:"fillSz"`
	FillTime string `json:"fillTime"`
	InstID string `json:"instId"`
	InstType string `json:"instType"`
	Lever string `json:"lever"`
	OrdID string `json:"ordId"`
	OrdType string `json:"ordType"`
	Pnl string `json:"pnl"`
	PosSide string `json:"posSide"`
	Px string `json:"px"`
	Rebate string `json:"rebate"`
	RebateCcy string `json:"rebateCcy"`
	Side string `json:"side"`
	SlOrdPx string `json:"slOrdPx"`
	SlTriggerPx string `json:"slTriggerPx"`
	State string `json:"state"`
	Sz string `json:"sz"`
	Tag string `json:"tag"`
	TdMode string `json:"tdMode"`
	TpOrdPx string `json:"tpOrdPx"`
	TpTriggerPx string `json:"tpTriggerPx"`
	TradeID string `json:"tradeId"`
	UTime string `json:"uTime"`
}

type GetV5MarketCandlesResult struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data [][]string `json:"data"`
}