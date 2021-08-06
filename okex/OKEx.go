package okex

import (
	"github.com/okex/V3-Open-API-SDK/okex-go-sdk-api"
	. "github.com/spectrelb/biapi"
	"net/http"
)

type OKEx struct {
	Client          *Client
	OKExSpot		*OKExSpot
}

func (ok OKEx) GetAccount(name string) (string, error) {
	return ok.OKExSpot.GetAccount(name)
}

func (ok OKEx) GetExchangeName() (string) {
	return "okex.com"
}

func (ok OKEx) GetMarketHistoryKline(symbol string, period KlinePeriod) ([]MarketKlineResp, error) {
	return ok.OKExSpot.GetMarketHistoryKline(symbol,period)
}

func (ok OKEx) GetAllSymbol() ([]string, error) {
	return ok.OKExSpot.GetAllSymbol()
}

func NewOKEx(config okex.Config, client *http.Client) *OKEx {
	ok := &OKEx{Client: &Client{
		Config: config,
		HttpClient:client,
	}}
	ok.OKExSpot = &OKExSpot{ok}
	return ok
}
