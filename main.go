package main

import (
	"github.com/iggyster/golang-design-patterns/behavioaral"
	"github.com/iggyster/golang-design-patterns/creational"
	"github.com/iggyster/golang-design-patterns/structural"
)

func main() {
	// Singleton
	creational.GetInstance()
	creational.GetInstance()

	// Decorator
	logger := structural.Logger{}
	upperLogger := structural.UpperCaseLogger{Logger: logger}

	upperLogger.LogMessage("lower case")

	// Strategy
	var strategy behavioaral.PaymentStrategy
	method := "cash"

	switch method {
	case "paypal":
		strategy = &behavioaral.PayPalPayment{Address: method}
	case "mastercard":
		strategy = &behavioaral.MasterCardPayment{Address: method}
	case "visa":
		strategy = &behavioaral.VisaPayment{Address: method}
	default:
		strategy = &behavioaral.Cash{Address: method}
	}

	strategy.Pay(100)

	// Facade
	engine := structural.Engine{}
	tank := structural.Tank{Amount: 40}
	car := structural.Car{Engine: engine, Tank: tank}

	car.StartEngine()
}
