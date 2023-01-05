package banking

import (
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	Balance Bitcoin
}

func (w *Wallet) Deposit(n Bitcoin) {
	w.Balance += n
}

func (w *Wallet) Withdraw(n Bitcoin) error {
	if n > w.Balance {
		return InsufficientBalance
	}

	w.Balance -= n
	return nil
}

type Prettier interface {
	ToPretty() string
}

func (b Bitcoin) ToPretty() string {
	return fmt.Sprintf("%.1f BTC", b)
}
