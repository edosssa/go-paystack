package paystack

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChargeCardAndAuthorization(t *testing.T) {
	assert := assert.New(t)
	// Card does not require validation
	card := Card{
		Number:      "4084084084084081",
		CVV:         "408",
		ExpiryMonth: fmt.Sprintf("%d", time.Now().Month()),
		ExpiryYear:  fmt.Sprintf("%d", time.Now().Year()),
	}
	charge := &ChargeRequest{
		Email:  "test@user.com",
		Amount: 10000,
		Card:   &card,
	}

	r, err := c.Charge.Create(charge)
	assert.NoError(err)
	assert.Equal(r["status"], "success")

	auth := r["authorization"].(map[string]interface{})
	authCode := auth["authorization_code"].(string)
	reusable := auth["reusable"].(bool)

	assert.True(reusable, "Card instruments should always be reusable")

	// Use the auth code to perform another charge
	charge = &ChargeRequest{
		Email:             "test@user.com",
		Amount:            10000,
		AuthorizationCode: authCode,
	}

	r, err = c.Charge.Create(charge)
	assert.NoError(err)
	assert.Equal(r["status"], "success")
}

func TestChargeBank(t *testing.T) {
	assert := assert.New(t)

	bankAccount := BankAccount{
		Code:          "057",
		AccountNumber: "0000000000",
	}

	charge := ChargeRequest{
		Email:    "test@user.com",
		Amount:   10000,
		Bank:     &bankAccount,
		Birthday: "1993-02-24",
	}

	r, err := c.Charge.Create(&charge)
	assert.NoError(err)
	assert.NotEmpty(r["reference"])
	assert.NotEmpty(r["status"])

	for {
		switch r["status"] {
		case "success":
			return
		case "send_otp":
			r, err = c.Charge.SubmitOTP("123456", r["reference"].(string))
			assert.NoError(err)
			assert.NotEmpty(r["reference"])
			assert.NotEmpty(r["status"])
		default:
			t.Errorf(`Expected status of success or send_otp, got %s`, r["status"])
		}
	}
}

func TestChargeServiceCheckPending(t *testing.T) {
	assert := assert.New(t)

	bankAccount := BankAccount{
		Code:          "057",
		AccountNumber: "0000000000",
	}

	charge := ChargeRequest{
		Email:    "test@user.com",
		Amount:   10000,
		Bank:     &bankAccount,
		Birthday: "1999-12-31",
	}

	r, err := c.Charge.Create(&charge)
	assert.NoError(err)
	assert.NotEmpty(r["reference"])
	assert.NotEmpty(r["status"])

	r, err = c.Charge.CheckPending(r["reference"].(string))
	assert.NoError(err)
	assert.NotEmpty(r["status"])
	assert.NotEmpty(r["reference"])
}
