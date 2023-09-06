package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/t1ery/wb_l2/calendar_api/data_store"
    "github.com/t1ery/wb_l2/calendar_api/domain"
    "strings"
)

func TestCreateEventHandler(t *testing.T) {

    // Создаём фейковый объект DataStore для теста
    dataStore := data_store.NewDataStore()

    // Создаём фейковый запрос HTTP для теста
    requestBody := "title=Моя встреча&date=2023-09-10&details=Детали встречи"
    req := httptest.NewRequest("POST", "/create_event", strings.NewReader(requestBody))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // Создаём фейковый ResponseWriter
    w := httptest.NewRecorder()

    // Вызовим функцию CreateEventHandler с фейковыми параметрами
    CreateEventHandler(w, req, dataStore)

    // Проверим, что ответ соответствует ожиданиям
    if w.Code != http.StatusOK {
        t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
    }
}

func TestUpdateEventHandler(t *testing.T) {

    // Создаём фейковый объект DataStore для теста
    dataStore := data_store.NewDataStore()

    // Внесём фейковое событие в хранилище для теста
    event := domain.Event{ID: 1, Title: "Тестовое событие", Date: "2023-09-10", Details: "Детали тестового события"}
    dataStore.CreateEvent(event)

    // Создаём фейковый запрос HTTP для теста
    requestBody := "id=1&title=Измененное событие&date=2023-09-15&details=Измененные детали"
    req := httptest.NewRequest("POST", "/update_event", strings.NewReader(requestBody))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // Создаём фейковый ResponseWriter
    w := httptest.NewRecorder()

    // Вызовим функцию UpdateEventHandler с фейковыми параметрами
    UpdateEventHandler(w, req, dataStore)

    // Проверим, что ответ соответствует ожиданиям
    if w.Code != http.StatusOK {
        t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
    }
}

func TestDeleteEventHandler(t *testing.T) {

    // Создаём фейковый объект DataStore для теста
    dataStore := data_store.NewDataStore()

    // Внесём фейковое событие в хранилище для теста
    event := domain.Event{ID: 1, Title: "Тестовое событие", Date: "2023-09-10", Details: "Детали тестового события"}
    dataStore.CreateEvent(event)

    // Создаём фейковый запрос HTTP для теста
    req := httptest.NewRequest("POST", "/delete_event?id=1", nil)

    // Создаём фейковый ResponseWriter
    w := httptest.NewRecorder()

    // Вызовим функцию DeleteEventHandler с фейковыми параметрами
    DeleteEventHandler(w, req, dataStore)

    // Проверим, что ответ соответствует ожиданиям
    if w.Code != http.StatusOK {
        t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
    }
}

func TestEventsForDayHandler(t *testing.T) {

    // Создаём фейковый объект DataStore для теста
    dataStore := data_store.NewDataStore()

    // Внесём фейковое событие в хранилище для теста
    event := domain.Event{ID: 1, Title: "Тестовое событие", Date: "2023-09-10", Details: "Детали тестового события"}
    dataStore.CreateEvent(event)

    // Создаём фейковый запрос HTTP для теста
    req := httptest.NewRequest("GET", "/events_for_day?date=2023-09-10", nil)

    // Создаём фейковый ResponseWriter
    w := httptest.NewRecorder()

    // Вызовим функцию EventsForDayHandler с фейковыми параметрами
    EventsForDayHandler(w, req, dataStore)

    // Проверим, что ответ соответствует ожиданиям
    if w.Code != http.StatusOK {
        t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
    }

}

func TestEventsForWeekHandler(t *testing.T) {

    // Создаём фейковый объект DataStore для теста
    dataStore := data_store.NewDataStore()

    // Внесём фейковое событие в хранилище для теста
    event := domain.Event{ID: 1, Title: "Тестовое событие", Date: "2023-09-10", Details: "Детали тестового события"}
    dataStore.CreateEvent(event)

    // Создаём фейковый запрос HTTP для теста
    req := httptest.NewRequest("GET", "/events_for_week?start_date=2023-09-10&end_date=2023-09-16", nil)

    // Создаём фейковый ResponseWriter
    w := httptest.NewRecorder()

    // Вызовим функцию EventsForWeekHandler с фейковыми параметрами
    EventsForWeekHandler(w, req, dataStore)

    // Проверим, что ответ соответствует ожиданиям
    if w.Code != http.StatusOK {
        t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
    }
}

func TestEventsForMonthHandler(t *testing.T) {

    // Создаём фейковый объект DataStore для теста
    dataStore := data_store.NewDataStore()

    // Внесём фейковое событие в хранилище для теста
    event := domain.Event{ID: 1, Title: "Тестовое событие", Date: "2023-09-10", Details: "Детали тестового события"}
    dataStore.CreateEvent(event)

    // Создаём фейковый запрос HTTP для теста
    req := httptest.NewRequest("GET", "/events_for_month?month=09&year=2023", nil)

    // Создаём фейковый ResponseWriter
    w := httptest.NewRecorder()

    // Вызовим функцию EventsForMonthHandler с фейковыми параметрами
    EventsForMonthHandler(w, req, dataStore)

    // Проверим, что ответ соответствует ожиданиям
    if w.Code != http.StatusOK {
        t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
    }
}
