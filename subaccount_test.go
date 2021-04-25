package paystack

import "testing"

func TestSubAccountCRUD(t *testing.T) {
	subAccount1 := &SubAccount{
		BusinessName:     "Dummy Business",
		SettlementBank:   "Zenith Bank",
		AccountNumber:    "0000000000",
		PercentageCharge: 18.2,
	}

	// create the subAccount
	subAccount, err := c.SubAccount.Create(subAccount1)
	if err != nil {
		t.Errorf("CREATE SubAccount returned error: %v", err)
	}

	if subAccount.SubAccountCode == "" {
		t.Errorf("Expected SubAccount code to be set")
	}

	// Retrieve the subAccount
	subAccount, err = c.SubAccount.Get(subAccount.ID)
	if err != nil {
		t.Errorf("GET SubAccount returned error: %v", err)
	}

	if subAccount.BusinessName != subAccount1.BusinessName {
		t.Errorf("Expected SubAccount BusinessName %v, got %v", subAccount.BusinessName, subAccount1.BusinessName)
	}

	// Retrieve the subAccount list
	// Todo: fix list subaccounts, skip for now
	t.SkipNow()
	subAccounts, err := c.SubAccount.List()
	if err != nil || !(len(subAccounts.Values) > 0) || !(subAccounts.Meta.Total > 0) {
		t.Errorf("Expected SubAccount list, got %d, returned error %v", len(subAccounts.Values), err)
	}
}
