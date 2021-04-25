package paystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankList(t *testing.T) {
	assert := assert.New(t)
	banks, err := c.Bank.List()
	if err != nil {
		t.Error(err)
	}
	assert.Greater(len(banks.Values), 0)
}

func TestResolveBVN(t *testing.T) {
	assert := assert.New(t)

	// Test invlaid BVN.
	// Err not nill. Resp status code is 400
	resp, err := c.Bank.ResolveBVN(21212917)
	assert.Error(err, "Expected error for invalid BVN")

	// Test free calls limit
	// Error is nil
	// &{Meta:{CallsThisMonth:0 FreeCallsLeft:0} BVN:cZ+MKrsLAqJCUi+hxIdQqw==}’
	resp, err = c.Bank.ResolveBVN(21212917741)
	if resp.Meta.FreeCallsLeft != 0 {
		t.Errorf("Expected free calls limit exceeded, got %+v'", resp)
	}
	// TODO(yao): Reproduce error: Your balance is not enough to fulfill this request
}

func TestResolveAccountNumber(t *testing.T) {
	assert := assert.New(t)
	account, err := c.Bank.ResolveAccountNumber("2208713487", "057")
	if err != nil {
		t.Error(err)
	}
	assert.Equal("EDOSA OSARIEMEN KELVIN", account.AccountName)
	assert.Equal("2208713487", account.AccountNumber)
	assert.Equal(21, account.BankID)
}
