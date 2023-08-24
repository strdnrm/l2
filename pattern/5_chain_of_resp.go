package main

import "fmt"

// Интерфейс обработчика объявляет метод обработки запроса и установку следующего обработчика
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// Базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *BaseHandler) HandleRequest(request string) {
	if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

// Конкретная реализация обработчика
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("Request handled by ConcreteHandlerA")
	} else {
		h.BaseHandler.HandleRequest(request)
	}
}

// Конкретная реализация обработчика
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("Request handled by ConcreteHandlerB")
	} else {
		h.BaseHandler.HandleRequest(request)
	}
}

// Конкретная реализация обработчика
type ConcreteHandlerC struct {
	BaseHandler
}

func (h *ConcreteHandlerC) HandleRequest(request string) {
	if request == "C" {
		fmt.Println("Request handled by ConcreteHandlerC")
	} else {
		h.BaseHandler.HandleRequest(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	handlerA.HandleRequest("A")
	handlerA.HandleRequest("B")
	handlerA.HandleRequest("C")
	handlerA.HandleRequest("D")
}
