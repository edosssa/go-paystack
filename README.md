# Paystack client for Golang

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/snapay-labs/go-paystack)

> This library would not be possible without the awesome work on [paystack-go](https://github.com/snapay/rn-paystack) by Yao Adzaku

paystack-go is a Go client library for accessing the Paystack API.

## Installation

go-paystack uses go modules so make sure you have a mod file in your project or generate one using 

```bash
go mod init github.com/my/repo
```

And then install go-paystack

```bash
go get github.com/snapay-labs/go-paystack
```

## Quickstart

Getting up and running using go-paystack is simple, see for yourself.

``` go
import "github.com/snapay-labs/go-paystack"

apiKey := "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"

// second param is an optional http client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
client := paystack.NewClient(apiKey)

recipient := &TransferRecipient{
    Type:          "Nuban",
    Name:          "Customer 1",
    Description:   "Demo customer",
    AccountNumber: "0100000010",
    BankCode:      "044",
    Currency:      "NGN",
    Metadata:      map[string]interface{}{"job": "Plumber"},
}

recipient1, err := client.Transfer.CreateRecipient(recipient)

req := &TransferRequest{
    Source:    "balance",
    Reason:    "Delivery pickup",
    Amount:    30,
    Recipient: recipient1.RecipientCode,
}

transfer, err := client.Transfer.Initiate(req)
if err != nil {
    // do something with error
}

// retrieve list of plans
plans, err := client.Plan.List()

for i, plan := range plans.Values {
  fmt.Printf("%+v", plan)
}

cust := &Customer{
    FirstName: "User123",
    LastName:  "AdminUser",
    Email:     "user123@gmail.com",
    Phone:     "+23400000000000000",
}
// create the customer
customer, err := client.Customer.Create(cust)
if err != nil {
    // do something with error
}

// Get customer by ID
customer, err := client.Customers.Get(customer.ID)
```

See the test files for more examples.

<br>

Made wth ❤️ by [Snapay](http://www.snapay.ng)