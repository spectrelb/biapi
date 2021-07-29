package okex

import (
	"github.com/okex/V3-Open-API-SDK/okex-go-sdk-api"
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

func NewOKEx(config okex.Config) *OKEx {
	ok := &OKEx{Client: &Client{
		Config: config,
	}}
	ok.OKExSpot = &OKExSpot{ok}
	return ok
}
