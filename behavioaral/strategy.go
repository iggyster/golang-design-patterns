package behavioaral

import "fmt"

type PaymentStrategy interface {
	Pay(money int)
}

type VisaPayment struct {
	Address string
}

func (v *VisaPayment) Pay(money int) {
	fmt.Printf("Pay %v with %v\n", money, v.Address)
}

type MasterCardPayment struct {
	Address string
}

func (m *MasterCardPayment) Pay(money int) {
	fmt.Printf("Pay %v with %v\n", money, m.Address)
}

type PayPalPayment struct {
	Address string
}

func (p *PayPalPayment) Pay(money int) {
	fmt.Printf("Pay %v with %v\n", money, p.Address)
}

type Cash struct {
	Address string
}

func (c *Cash) Pay(money int) {
	fmt.Printf("Pay %v with %v\n", money, c.Address)
}
