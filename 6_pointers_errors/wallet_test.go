package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)
		assertBalance(t, &wallet, Bitcoin(15))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		assertError(t, wallet.Withdraw(Bitcoin(100)), ErrInsufficientFunds)
		assertBalance(t, &wallet, startingBalance)
	})
}

func assertBalance(t *testing.T, wallet *Wallet, expectedBalance Bitcoin) {
	t.Helper()
	actualBalance := wallet.Balance()
	if actualBalance != expectedBalance {
		t.Errorf("got %s, want %s", actualBalance, expectedBalance)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("got an error but didn't want one")
	}
}

func assertError(t testing.TB, err error, expectedError error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
	if err != ErrInsufficientFunds {
		t.Error("wanted an ErrInsufficientFunds but didn't get one")
	}
}
