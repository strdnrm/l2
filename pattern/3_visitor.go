package main

import "fmt"

// Интерфейс посетителя объявляет методы посещения для каждого класса элемента
type Visitor interface {
	VisitElementA(element ElementA)
	VisitElementB(element ElementB)
}

// Конкретный посетитель реализует методы посещения для каждого класса элемента
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitElementA(element ElementA) {
	fmt.Println("Visit Element A")
}

func (v *ConcreteVisitor) VisitElementB(element ElementB) {
	fmt.Println("Visit Element B")
}

// Интерфейс элемента объявляет метод принятия посетителя
type Element interface {
	Accept(visitor Visitor)
}

// Конкретные элементы реализуют метод принятия посетителя
type ElementA struct{}

func (e ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

type ElementB struct{}

func (e ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

// Клиентский код работает с объектами через интерфейсы посетителя и элемента,
// не завися от конкретных классов элементов и посетителей
func ClientCode(elements []Element, visitor Visitor) {
	for _, element := range elements {
		element.Accept(visitor)
	}
}

func main() {
	elements := []Element{ElementA{}, ElementB{}}
	visitor := &ConcreteVisitor{}

	ClientCode(elements, visitor)
}
