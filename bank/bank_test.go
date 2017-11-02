package bank

import "testing"

func TestDeposit(t *testing.T) {
	Deposit(10)
	balance = Balance()

	if balance != 10 {
		t.Errorf("Incorrect balance, got: %d, want: %d", balance, 10)
	}
}
