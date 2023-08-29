package main

import "fmt"

// Handler - интерфейс обработчика запроса
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// ConcreteHandler - конкретный обработчик запроса
type ConcreteHandler struct {
	next        Handler
	handlerName string
}

func NewConcreteHandler(name string) *ConcreteHandler {
	return &ConcreteHandler{handlerName: name}
}

func (c *ConcreteHandler) SetNext(handler Handler) {
	c.next = handler
}

func (c *ConcreteHandler) HandleRequest(request string) {
	fmt.Printf("%s обрабатывает запрос: %s\n", c.handlerName, request)
	if c.next != nil {
		c.next.HandleRequest(request)
	}
}

func main() {
	handler1 := NewConcreteHandler("Обработчик 1")
	handler2 := NewConcreteHandler("Обработчик 2")
	handler3 := NewConcreteHandler("Обработчик 3")

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	handler1.HandleRequest("Запрос 1")
}
