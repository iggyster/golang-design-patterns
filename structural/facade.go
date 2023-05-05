package structural

import "fmt"

type Engine struct {
}

func (e *Engine) StartEngine() {
	fmt.Println("Engine started")
}

type Tank struct {
	Amount int
}

func (t *Tank) CheckGasoline() bool {
	return t.Amount > 0
}

type Car struct {
	Engine
	Tank
}

func (c Car) Drive() error {
	if !c.CheckGasoline() {
		return fmt.Errorf("no gasoline in a tank\n")
	}

	c.StartEngine()

	return nil
}
