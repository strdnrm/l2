package main

import "fmt"

// Интерфейс стратегии
type Strategy interface {
	Execute()
}

type StrategyA struct{}

func (s *StrategyA) Execute() {
	fmt.Println("Executing Strategy A")
}

type StrategyB struct{}

func (s *StrategyB) Execute() {
	fmt.Println("Executing Strategy B")
}

// Контекст, использующий стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	// Создание контекста
	context := &Context{}

	// Установка стратегии A
	strategyA := &StrategyA{}
	context.SetStrategy(strategyA)
	context.ExecuteStrategy() // Вывод: Executing Strategy A

	// Установка стратегии B
	strategyB := &StrategyB{}
	context.SetStrategy(strategyB)
	context.ExecuteStrategy() // Вывод: Executing Strategy B
}
