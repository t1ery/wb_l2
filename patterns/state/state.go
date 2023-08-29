package main

import (
	"fmt"
)

// Интерфейс состояния заказа
type OrderState interface {
	SubmitOrder() string
	PayOrder() string
	ShipOrder() string
	CompleteOrder() string
}

// Конкретное состояние: Новый заказ
type NewOrder struct{}

func (o *NewOrder) SubmitOrder() string {
	return "Заказ отправлен на обработку."
}

func (o *NewOrder) PayOrder() string {
	return "Оплата заказа принята."
}

func (o *NewOrder) ShipOrder() string {
	return "Заказ не может быть отправлен до оплаты."
}

func (o *NewOrder) CompleteOrder() string {
	return "Заказ не может быть завершен до отправки и оплаты."
}

// Конкретное состояние: Заказ оплачен
type PaidOrder struct{}

func (o *PaidOrder) SubmitOrder() string {
	return "Заказ отправлен на обработку."
}

func (o *PaidOrder) PayOrder() string {
	return "Заказ уже оплачен."
}

func (o *PaidOrder) ShipOrder() string {
	return "Заказ отправлен клиенту."
}

func (o *PaidOrder) CompleteOrder() string {
	return "Заказ не может быть завершен без подтверждения доставки."
}

// Контекст заказа
type Order struct {
	state OrderState
}

func NewOrderInstance() *Order {
	return &Order{state: &NewOrder{}}
}

func (o *Order) SetState(state OrderState) {
	o.state = state
}

func (o *Order) SubmitOrder() string {
	return o.state.SubmitOrder()
}

func (o *Order) PayOrder() string {
	return o.state.PayOrder()
}

func (o *Order) ShipOrder() string {
	return o.state.ShipOrder()
}

func (o *Order) CompleteOrder() string {
	return o.state.CompleteOrder()
}

func main() {
	order := NewOrderInstance()

	fmt.Println(order.SubmitOrder())
	fmt.Println(order.PayOrder())
	fmt.Println(order.ShipOrder())

	// Меняем состояние заказа на "Заказ оплачен"
	order.SetState(&PaidOrder{})

	fmt.Println(order.ShipOrder())
	fmt.Println(order.CompleteOrder())
}
