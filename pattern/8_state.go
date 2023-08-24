package main

import "fmt"

// Интерфейс состояния
type State interface {
	Handle()
}

type StateA struct{}

func (s *StateA) Handle() {
	fmt.Println("Handling State A")
}

type StateB struct{}

func (s *StateB) Handle() {
	fmt.Println("Handling State B")
}

// Контекст, использующий состояние
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) HandleState() {
	c.state.Handle()
}

func main() {
	// Создание контекста
	context := &Context{}

	// Установка состояния A
	stateA := &StateA{}
	context.SetState(stateA)
	context.HandleState()

	// Установка состояния B
	stateB := &StateB{}
	context.SetState(stateB)
	context.HandleState()
}
