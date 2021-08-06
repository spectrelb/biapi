package okex

import (
	"github.com/okex/V3-Open-API-SDK/okex-go-sdk-api"
	. "github.com/spectrelb/biapi"
)

type OKExSpot struct {
	*OKEx
}

func (ok *OKExSpot) GetAccount(name string) (string, error)  {
	_, _ = ok.Client.GetV5AccountBalance(&name)
	return "123", nil
}

func (ok *OKExSpot) GetMarketHistoryKline(symbol string, period KlinePeriod) ([]MarketKlineResp, error)  {
	var data []MarketKlineResp
	var after string
	for i := 0; i < 2; i++ {
		optionals := okex.NewParams()
		optionals["instId"] = symbol
		optionals["bar"] = OkexKlinePeriodConverter[period]
		optionals["limit"] = "100"
		optionals["after"] = after
		result, err := ok.Client.GetV5MarketCandles(&optionals)
		if err !=nil {
			return nil, err
		}

		a := len(result.Data) - 1

		if len(result.Data) >0 {
			after = result.Data[a:][0][0]
		}

		for _, v := range result.Data {
			kline := MarketKlineResp{
				Ts:v[0],
				Open:v[1],
				High:v[2],
				Low:v[3],
				Close:v[4],
				Vol:v[5],
			}
			data = append(data, kline)
		}
	}

	return data, nil
}

func (ok *OKExSpot) GetAllSymbol() ([]string, error)  {
	instType := SPOT
	result, err := ok.Client.GetV5PublicInstruments(&instType)
	if err !=nil {
		return nil, err
	}

	var resp []string
	for _, v := range result.Data {
		if v.State == "live" && (v.QuoteCcy == "USDT" || v.QuoteCcy == "BTC" || v.QuoteCcy == "ETH" ) {
			resp = append(resp, v.InstID)
		}
	}

	return resp, nil
}