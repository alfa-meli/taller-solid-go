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

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bank := &BankPayment{}
	ProcessPayment(bank)
}
