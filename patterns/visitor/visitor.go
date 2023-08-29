package main

import "fmt"

// Shape - интерфейс фигуры, который будет посещен
type Shape interface {
	Draw()
	Accept(visitor Visitor)
}

// Line - конкретная фигура (линия)
type Line struct {
}

func (l *Line) Draw() {
	fmt.Println("Рисуем линию")
}

func (l *Line) Accept(visitor Visitor) {
	visitor.VisitLine(l)
}

// Circle - конкретная фигура (круг)
type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Рисуем круг")
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

// Text - конкретная фигура (текст)
type Text struct {
}

func (t *Text) Draw() {
	fmt.Println("Выводим текст")
}

func (t *Text) Accept(visitor Visitor) {
	visitor.VisitText(t)
}

// Visitor - интерфейс посетителя
type Visitor interface {
	VisitLine(line *Line)
	VisitCircle(circle *Circle)
	VisitText(text *Text)
}

// ExportPNGVisitor - конкретный посетитель для экспорта в PNG
type ExportPNGVisitor struct {
}

func (v *ExportPNGVisitor) VisitLine(*Line) {
	fmt.Println("Экспортируем линию в PNG")
}

func (v *ExportPNGVisitor) VisitCircle(*Circle) {
	fmt.Println("Экспортируем круг в PNG")
}

func (v *ExportPNGVisitor) VisitText(*Text) {
	fmt.Println("Экспортируем текст в PNG")
}

// ExportSVGVisitor - конкретный посетитель для экспорта в SVG
type ExportSVGVisitor struct {
}

func (v *ExportSVGVisitor) VisitLine(*Line) {
	fmt.Println("Экспортируем линию в SVG")
}

func (v *ExportSVGVisitor) VisitCircle(*Circle) {
	fmt.Println("Экспортируем круг в SVG")
}

func (v *ExportSVGVisitor) VisitText(*Text) {
	fmt.Println("Экспортируем текст в SVG")
}

func main() {
	shapes := []Shape{&Line{}, &Circle{}, &Text{}}

	pngVisitor := &ExportPNGVisitor{}
	svgVisitor := &ExportSVGVisitor{}

	fmt.Println("Экспорт в PNG:")
	for _, shape := range shapes {
		shape.Draw()
		shape.Accept(pngVisitor)
	}

	fmt.Println("Экспорт в SVG:")
	for _, shape := range shapes {
		shape.Draw()
		shape.Accept(svgVisitor)
	}
}
