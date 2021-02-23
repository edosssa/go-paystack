package paystack

import (
	"fmt"
	"net/url"
)

// ChargeService allows you to configure payment channel of your choice when initiating a payment.
// For more details see https://paystack.com/docs/api/#charge
type ChargeService service

// Card represents a Card object
type Card struct {
	Number            string `json:"card_number,omitempty"`
	CVV               string `json:"card_cvc,omitempty"`
	ExpirtyMonth      string `json:"expiry_month,omitempty"`
	ExpiryYear        string `json:"expiry_year,omitempty"`
	AddressLine1      string `json:"address_line1,omitempty"`
	AddressLine2      string `json:"address_line2,omitempty"`
	AddressLine3      string `json:"address_line3,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	AddressPostalCode string `json:"address_postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
}

// BankAccount is used as bank in a charge request
// See https://paystack.com/docs/payments/payment-channels#bank-accounts for more details
type BankAccount struct {
	Code          string `json:"code,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
}

// MobileMoney represents a mobile money payment channel
// See https://paystack.com/docs/payments/payment-channels#mobile-money for more details
type MobileMoney struct {
	Phone string `json:"phone,omitempty"`
	// valid providers are "mtn", "voda", and "tgo"
	Provider string `json:"provider,omitempty"`
}

// QRCode todo: add better documentation
// See https://paystack.com/docs/payments/payment-channels#qr-code for more details
type QRCode struct {
	Provider string `json:"provider,omitempty"`
}

// ChargeRequest represents a Paystack charge request
type ChargeRequest struct {
	Email    string  `json:"email,omitempty"`
	Amount   float32 `json:"amount,omitempty"`
	Birthday string  `json:"birthday,omitempty"`
	// Only use this payment channel if your organization is PCI DSS compliant
	// See https://paystack.com/docs/payments/payment-channels#cards for more details
	Card              *Card        `json:"card,omitempty"`
	Bank              *BankAccount `json:"bank,omitempty"`
	AuthorizationCode string       `json:"authorization_code,omitempty"`
	Pin               string       `json:"pin,omitempty"`
	MobileMoney       *MobileMoney `json:"mobile_money,omitempty"`
	Metadata          *Metadata    `json:"metadata,omitempty"`
	QR                *QRCode      `json:"qr,omitempty"`
}

// Create submits a charge request using card details or bank details or authorization code
// For more details see https://paystack.com/docs/api/#charge-create
func (s *ChargeService) Create(req *ChargeRequest) (Response, error) {
	resp := Response{}
	err := s.client.Call("POST", "/charge", req, &resp)
	return resp, err
}

// Tokenize tokenizes payment instrument before a charge
// For more details see https://developers.paystack.co/v1.0/reference#charge-tokenize
func (s *ChargeService) Tokenize(req *ChargeRequest) (Response, error) {
	resp := Response{}
	err := s.client.Call("POST", "/charge/tokenize", req, &resp)
	return resp, err
}

// SubmitPIN submits PIN to continue a charge
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitPIN(pin, reference string) (Response, error) {
	data := url.Values{}
	data.Add("pin", pin)
	data.Add("reference", reference)
	resp := Response{}
	err := s.client.Call("POST", "/charge/submit_pin", data, &resp)
	return resp, err
}

// SubmitOTP submits OTP to continue a charge
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitOTP(otp, reference string) (Response, error) {
	data := url.Values{}
	data.Add("pin", otp)
	data.Add("reference", reference)
	resp := Response{}
	err := s.client.Call("POST", "/charge/submit_otp", data, &resp)
	return resp, err
}

// SubmitPhone submits Phone when requested
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitPhone(phone, reference string) (Response, error) {
	data := url.Values{}
	data.Add("pin", phone)
	data.Add("reference", reference)
	resp := Response{}
	err := s.client.Call("POST", "/charge/submit_phone", data, &resp)
	return resp, err
}

// SubmitBirthday submits Birthday when requested
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *ChargeService) SubmitBirthday(birthday, reference string) (Response, error) {
	data := url.Values{}
	data.Add("pin", birthday)
	data.Add("reference", reference)
	resp := Response{}
	err := s.client.Call("POST", "/charge/submit_birthday", data, &resp)
	return resp, err
}

// CheckPending returns pending charges
// When you get "pending" as a charge status, wait 30 seconds or more,
// then make a check to see if its status has changed. Don't call too early as you may get a lot more pending than you should.
// For more details see https://developers.paystack.co/v1.0/reference#check-pending-charge
func (s *ChargeService) CheckPending(reference string) (Response, error) {
	u := fmt.Sprintf("/charge/%s", reference)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}
