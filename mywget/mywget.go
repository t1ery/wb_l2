package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Проверяем, был ли передан URL как аргумент командной строки.
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run wget.go <URL>")
		return
	}

	// Получаем URL из аргументов командной строки.
	url := os.Args[1]

	// Вызываем функцию для загрузки сайта.
	err := DownloadSite(url)
	if err != nil {
		fmt.Printf("Ошибка при загрузке сайта: %v\n", err)
	}
}

// DownloadSite загружает сайт целиком и сохраняет его в файл index.html.
func DownloadSite(url string) error {
	// Открываем HTTP-запрос к указанному URL.
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Создаем файл index.html для сохранения сайта.
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем содержимое ответа в файл.
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Сайт успешно загружен и сохранен в файл index.html\n")
	return nil
}
