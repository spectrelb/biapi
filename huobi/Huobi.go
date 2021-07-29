package huobi

import (
	"net/http"
)

type Huobi struct {
	Client          *http.Client
}

func (ok Huobi) GetAccount(name string) (string, error) {
	return "123", nil
}

func (ok Huobi) GetExchangeName() (string) {
	return "huobi.com"
}

func NewHuobi() *Huobi {
	return &Huobi{}
}