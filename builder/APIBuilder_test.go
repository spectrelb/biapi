package builder

import (
	"fmt"
	"github.com/spectrelb/biapi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var (
	proxy = func(_ *http.Request) (*url.URL, error) {
		//return url.Parse("http://127.0.0.1:1087")
		return url.Parse("http://127.0.0.1:19180")
	}
	transport = &http.Transport{Proxy: proxy}

	aa = &http.Client{
		Transport: transport,
		Timeout: time.Duration(5) * time.Second,
	}

	okApi = NewCustomAPIBuilder(aa).APIKey("xxxxx").APISecretkey("xxxx").ApiPassphrase("xxxx").Build(HUOBI)
)

func TestAPIBuilder_Build(t *testing.T) {
	assert.Equal(t, okApi.GetExchangeName(), OKEX)
}

func TestAPIGetMarketHistoryKlineBuilder_Build(t *testing.T) {
	resp, _ := okApi.GetMarketHistoryKline("ETH-USDT", biapi.KLINE_PERIOD_4H, 200)
	fmt.Println(resp)
}

func TestAPIGetAllSymbolBuilder_Build(t *testing.T) {
	resp, _ := okApi.GetAllSymbol()
	fmt.Println(resp)
}