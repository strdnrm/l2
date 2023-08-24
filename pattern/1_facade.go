package main

import "fmt"

// SubsystemA и SubsystemB представляют сложные подсистемы системы
type SubsystemA struct{}

func (s *SubsystemA) OperationA() {
	fmt.Println("SubsystemA: OperationA")
}

type SubsystemB struct{}

func (s *SubsystemB) OperationB() {
	fmt.Println("SubsystemB: OperationB")
}

// Facade предоставляет простой интерфейс для работы с подсистемами
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

func (f *Facade) Operation() {
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
