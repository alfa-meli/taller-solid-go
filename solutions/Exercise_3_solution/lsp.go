package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

type BankPayment struct {
}

func (BankPayment) Pay(account int) {
	fmt.Printf("Payment using Bankaccount %d\n", account)
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bank := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 12345,
	}
	ProcessPayment(bank)
}
