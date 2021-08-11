package huobi

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	. "github.com/spectrelb/biapi"
	"net/http"
	"strconv"
	"time"
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

func (ok Huobi) GetMarketHistoryKline(symbol string, period KlinePeriod, limit int) ([]MarketKlineResp, error) {
	client := new(client.MarketClient).Init(config.Host)

	optionalRequest := market.GetCandlestickOptionalRequest{Period: HuobiKlinePeriodConverter[period], Size: 200}
	result, err := client.GetCandlestick(symbol, optionalRequest)

	if err !=nil {
		return nil, err
	}

	var data []MarketKlineResp
	for _, v := range result {
		date := time.Unix(v.Id, 0).Format("2006-01-02 15:04:05")

		open, _ := v.Open.Float64()
		high, _ := v.High.Float64()
		low, _ := v.Low.Float64()
		closeF, _ := v.Close.Float64()
		vol, _ := v.Vol.Float64()

		kline := MarketKlineResp{
			Ts:date,
			Open: strconv.FormatFloat(open, 'g', -1, 64),
			High: strconv.FormatFloat(high, 'g', -1, 64),
			Low: strconv.FormatFloat(low, 'g', -1, 64),
			Close: strconv.FormatFloat(closeF, 'g', -1, 64),
			Vol: strconv.FormatFloat(vol, 'g', -1, 64),
		}
		data = append(data, kline)
	}

	return data, nil
}

func (ok Huobi) GetAllSymbol() ([]string, error) {
	client := new(client.CommonClient).Init(config.Host)
	result, err := client.GetSymbols()
	if err != nil {
		return nil, err
	}

	var resp []string
	for _, v := range result {
		if v.State == "live" && (v.QuoteCurrency == "usdt" || v.QuoteCurrency == "btc" || v.QuoteCurrency == "eth" ) {
			resp = append(resp, fmt.Sprintf("%s-%s", v.BaseCurrency, v.QuoteCurrency))
		}
	}

	return resp, nil
}

func NewHuobi() *Huobi {
	return &Huobi{}
}