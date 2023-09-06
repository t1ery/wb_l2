package utils

import (
    "fmt"
    "net/http"
    "github.com/t1ery/wb_l2/calendar_api/domain"
    "time"
)

// Валидация параметров создания события
func ValidateCreateEventParams(r *http.Request) (domain.Event, error) {
    var newEvent domain.Event
    title := r.FormValue("title")
    date := r.FormValue("date")
    details := r.FormValue("details")

    // Валидация параметров
    if title == "" || date == "" || details == "" {
        return newEvent, fmt.Errorf("Отсутствуют обязательные параметры")
    }

    // Проверяем формат даты
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return newEvent, fmt.Errorf("Неверный формат даты")
	}

    newEvent = domain.Event{
        Title:   title,
        Date:    date,
        Details: details,
    }

    return newEvent, nil
}
