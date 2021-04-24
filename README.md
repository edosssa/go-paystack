# Paystack client for Golang

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ZappieLabs/go-paystack)

go-paystack is a Go client library for accessing the Paystack API.

## Installation

go-paystack is available using the standard go get command.

```bash
go get github.com/ZappieLabs/go-paystack
```

Before running tests, you need to set the ```PAYSTACK_KEY``` environment variable.
You can easily do this by creating a ```.env``` file in your project root with the following contents and we'll load it automatically when tests are run. Replace ```paystack-secret-key``` with your **test** secret key.

```
PAYSTACK_KEY=<paystack-secret-key>
```

If you don't have a test secret key, you can create one from your [paystack dashboard](https://dashboard.paystack.com/#/settings/profile).


> ⚠️ We don't allow the use of production secret keys while running tests. Test secret keys always have an **sk_test_** prefix.

Run tests

```bash
go test ./...
```



## Quickstart

Getting up and running using go-paystack is simple, see for yourself.

```go
import "github.com/ZappieLabs/go-paystack"

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

Made wth ❤️ by [Zappie](http://zappie.co)
