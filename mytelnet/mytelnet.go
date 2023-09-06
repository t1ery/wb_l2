package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Парсим аргументы командной строки
	host := flag.String("host", "", "Хост для подключения")
	port := flag.String("port", "23", "Порт для подключения")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут для подключения")

	flag.Parse()

	// Собираем адрес для подключения
	address := fmt.Sprintf("%s:%s", *host, *port)

	// Устанавливаем таймаут для подключения
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Не удалось подключиться к серверу: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Подключено к %s\n", address)

	// Запускаем горутину для чтения данных из сокета и вывода на STDOUT
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Соединение разорвано.")
				os.Exit(0)
			}
			fmt.Print(string(buffer[:n]))
		}
	}()

	fmt.Println("Начинаем передавать ввод в сокет и получать вывод...")

	// Читаем ввод пользователя с консоли и отправляем в сокет
	buffer := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			os.Exit(1)
		}
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Ошибка отправки данных:", err)
			os.Exit(1)
		}
	}
}
