package biapi


type MarketKlineResp struct {
	Open  string     `json:"open"`
	Close string     `json:"close"`
	Low   string     `json:"low"`
	High  string     `json:"high"`
	Ts    string     `json:"ts"`
	Vol   string     `json:"vol"`
}

type SymbolResp struct {
	Name  string     `json:"name"`
}