package paystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var c *Client

func init() {
	apiKey := mustGetTestKey()
	c = NewClient(apiKey, nil)
}

func TestResolveCardBIN(t *testing.T) {
	assert := assert.New(t)
	r, err := c.ResolveCardBIN(59983)
	assert.NoError(err)
	assert.NotEmpty(r["bin"])
}

func TestCheckBalance(t *testing.T) {
	assert := assert.New(t)
	balances, err := c.CheckBalance()
	assert.NoError(err)
	// Ideally, there are two balances, NGN and USD
	assert.GreaterOrEqual(len(balances), 2)
}

func TestSessionTimeout(t *testing.T) {
	assert := assert.New(t)
	r, err := c.GetSessionTimeout()
	assert.NoError(err)
	assert.NotNil(r["payment_session_timeout"])
}
