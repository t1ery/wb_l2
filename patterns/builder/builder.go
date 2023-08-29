package main

import "fmt"

// Document - класс документа, который мы хотим построить
type Document struct {
	Header    string
	Text      string
	ImagePath string
}

// DocumentBuilder - интерфейс строителя для документа
type DocumentBuilder interface {
	SetHeader(header string)
	SetText(text string)
	SetImage(imagePath string)
	GetDocument() *Document
}

// SimpleDocumentBuilder - конкретный строитель для простого документа
type SimpleDocumentBuilder struct {
	document *Document
}

func NewSimpleDocumentBuilder() *SimpleDocumentBuilder {
	return &SimpleDocumentBuilder{
		document: &Document{},
	}
}

func (b *SimpleDocumentBuilder) SetHeader(header string) {
	b.document.Header = header
}

func (b *SimpleDocumentBuilder) SetText(text string) {
	b.document.Text = text
}

func (b *SimpleDocumentBuilder) SetImage(imagePath string) {
	b.document.ImagePath = imagePath
}

func (b *SimpleDocumentBuilder) GetDocument() *Document {
	return b.document
}

// Director - директор, который управляет процессом конструирования
type Director struct {
	builder DocumentBuilder
}

func NewDirector(builder DocumentBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.SetHeader("Простой документ")
	d.builder.SetText("Это текст простого документа.")
	d.builder.SetImage("image.jpg")
}

func main() {
	builder := NewSimpleDocumentBuilder()
	director := NewDirector(builder)

	director.Construct()
	document := builder.GetDocument()

	fmt.Println("Заголовок:", document.Header)
	fmt.Println("Текст:", document.Text)
	fmt.Println("Изображение:", document.ImagePath)
}
