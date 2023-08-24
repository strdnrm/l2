package main

import "fmt"

// Интерфейс команды объявляет метод выполнения команды
type Command interface {
	Execute()
}

// Получатель команды выполняет действие
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Action executed")
}

// Конкретная команда реализует метод выполнения команды
type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Инициатор команды связывает команду с получателем и инициирует выполнение команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := &ConcreteCommand{receiver: *receiver}
	invoker := &Invoker{}

	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
