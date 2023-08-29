package main

import "fmt"

// Интерфейс для рисовальных инструментов
type DrawingTool interface {
	Draw() string
}

// Конкретная фабрика для создания кисти
type BrushFactory struct{}

func (bf BrushFactory) CreateTool() DrawingTool {
	return &Brush{}
}

// Конкретная фабрика для создания карандаша
type PencilFactory struct{}

func (pf PencilFactory) CreateTool() DrawingTool {
	return &Pencil{}
}

// Конкретная фабрика для создания ластика
type EraserFactory struct{}

func (ef EraserFactory) CreateTool() DrawingTool {
	return &Eraser{}
}

// Конкретные рисовальные инструменты: кисть, карандаш и ластик
type Brush struct{}

func (b Brush) Draw() string {
	return "Рисуем кистью"
}

type Pencil struct{}

func (p Pencil) Draw() string {
	return "Рисуем карандашом"
}

type Eraser struct{}

func (e Eraser) Draw() string {
	return "Используем ластик"
}

func main() {
	// Создаем фабрику для кисти
	brushFactory := BrushFactory{}
	// Используем фабрику для создания кисти
	brush := brushFactory.CreateTool()
	fmt.Println(brush.Draw()) // Вывод: "Рисуем кистью"

	// Аналогично для карандаша
	pencilFactory := PencilFactory{}
	pencil := pencilFactory.CreateTool()
	fmt.Println(pencil.Draw()) // Вывод: "Рисуем карандашом"

	// И для ластика
	eraserFactory := EraserFactory{}
	eraser := eraserFactory.CreateTool()
	fmt.Println(eraser.Draw()) // Вывод: "Используем ластик"
}
