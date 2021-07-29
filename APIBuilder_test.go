package biapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var okApi = NewAPIBuilder().APIKey("").APISecretkey("").ApiPassphrase("").Build(OKEX)

func TestAPIBuilder_Build(t *testing.T) {
	assert.Equal(t, okApi.GetExchangeName(), OKEX)
}