package interfaces

import "fmt"

type PaymentMethod interface {
	Process(amount float64) bool
	Provider() string
}

type CardPayment struct {
	CardNumber string
	Limit      float64
}

func (c *CardPayment) Process(amount float64) bool {
	check := c.Limit >= amount
	if check {
		c.Limit -= amount
	}
	return check
}

func (c CardPayment) Provider() string {
	return "CARD"
}

type UPIPayment struct {
	VPA string
}

func (u UPIPayment) Process(amount float64) bool {
	return true
}

func (u UPIPayment) Provider() string {
	return "UPI"
}

type CryptoPayment struct {
	Wallet  string
	Balance float64
}

func (c *CryptoPayment) Process(amount float64) bool {
	check := c.Balance >= amount
	if check {
		c.Balance -= amount
	}
	return check
}

func (c CryptoPayment) Provider() string {
	return "CRYPTO"
}

func Checkout(p PaymentMethod, amount float64) string {
	// Read README.md for the instructions
	succeed := p.Process(amount)
	if succeed {
		return fmt.Sprintf("Payment successful via %s", p.Provider())
	}

	return fmt.Sprintf("Payment failed via %s", p.Provider())
}

func DetectCrypto(p PaymentMethod) bool {
	_, ok := p.(*CryptoPayment)
	return ok
}
