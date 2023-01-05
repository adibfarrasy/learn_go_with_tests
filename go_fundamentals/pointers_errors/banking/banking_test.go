package banking_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pointers_errors/banking"
)

var _ = Describe("Banking", func() {
	When("user wants to store money in the wallet", func() {
		Context("given the user deposits 10 Bitcoins", func() {
			It("should store 10 Bitcoins in the wallet", func() {
				wallet := banking.Wallet{}
				wallet.Deposit(10.0)

				Expect(wallet.Balance.ToPretty()).To(Equal("10.0 BTC"))
			})
		})
	})

	When("user wants to withdraw money from the wallet", func() {
		Context("given the user withdraw 10 Bitcoins and whe wallet amount is 20 Bitcoins", func() {
			It("should deduct 10 Bitcoins from the wallet", func() {
				wallet := banking.Wallet{Balance: banking.Bitcoin(20.0)}
				err := wallet.Withdraw(banking.Bitcoin(10.0))

				Expect(wallet.Balance.ToPretty()).To(Equal("10.0 BTC"))
				Expect(err).To(Succeed())
			})
		})

		Context("given the user withdraw 100 Bitcoins and whe wallet amount is 20 Bitcoins", func() {
			It("should not allow the deduction", func() {
				wallet := banking.Wallet{Balance: banking.Bitcoin(20.0)}
				err := wallet.Withdraw(banking.Bitcoin(100.0))

				Expect(wallet.Balance.ToPretty()).To(Equal("20.0 BTC"))
				Expect(err).To(MatchError(banking.InsufficientBalance))
			})
		})
	})
})
