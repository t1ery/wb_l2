package main

import "fmt"

// Интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Конкретная стратегия: оплата кредитной картой
type CreditCardPayment struct {
	cardNumber string
}

func NewCreditCardPayment(cardNumber string) *CreditCardPayment {
	return &CreditCardPayment{cardNumber}
}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено $%.2f с помощью кредитной карты %s", amount, c.cardNumber)
}

// Конкретная стратегия: оплата через PayPal
type PayPalPayment struct {
	email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{email}
}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено $%.2f через PayPal с адреса %s", amount, p.email)
}

// Контекст, использующий стратегию оплаты
type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (s *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	s.paymentStrategy = strategy
}

func (s *ShoppingCart) Checkout(amount float64) string {
	return s.paymentStrategy.Pay(amount)
}

func main() {
	cart := NewShoppingCart()

	// Выбираем стратегию оплаты: кредитная карта
	creditCardPayment := NewCreditCardPayment("1234-5678-9876-5432")
	cart.SetPaymentStrategy(creditCardPayment)

	// Выполняем оплату
	result := cart.Checkout(100.0)
	fmt.Println(result)

	// Меняем стратегию оплаты: PayPal
	payPalPayment := NewPayPalPayment("example@example.com")
	cart.SetPaymentStrategy(payPalPayment)

	// Выполняем оплату
	result = cart.Checkout(50.0)
	fmt.Println(result)
}
