package biapi

import (
	okex2 "github.com/okex/V3-Open-API-SDK/okex-go-sdk-api"
	"github.com/spectrelb/biapi/huobi"
	"github.com/spectrelb/biapi/okex"
	"time"
)

const (
	OKEX            = "okex.com"
	HUOBI           = "huobi.com"
	BINANCE         = "binance.com"
)

type APIBuilder struct {
	httpTimeout      time.Duration
	apiKey           string
	secretkey        string
	apiPassphrase    string
}

func NewAPIBuilder() (builder *APIBuilder) {
	return  &APIBuilder{}
}

func (builder *APIBuilder) APIKey(key string) (_builder *APIBuilder) {
	builder.apiKey = key
	return builder
}

func (builder *APIBuilder) APISecretkey(key string) (_builder *APIBuilder) {
	builder.secretkey = key
	return builder
}

func (builder *APIBuilder) ApiPassphrase(apiPassphrase string) (_builder *APIBuilder) {
	builder.apiPassphrase = apiPassphrase
	return builder
}

//NewAPI return Api instance by type
func (builder *APIBuilder) Build(exName string) API {
	var _api API
	switch exName {
	case HUOBI:
		_api = huobi.NewHuobi()
	case OKEX:
		_api = okex.NewOKEx(okex2.Config{
			ApiKey:     builder.apiKey,
			SecretKey:  builder.secretkey,
			Passphrase: builder.apiPassphrase,
		})
	case BINANCE:
	default:
		println("exchange name error [" + exName + "].")
	}

	return _api
}