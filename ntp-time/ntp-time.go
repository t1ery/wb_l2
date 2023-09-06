package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем точное время с NTP-сервера
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	// Выводим текущее время
	fmt.Printf("Точное время: %s\n", ntpTime.Format(time.RFC3339))
}
