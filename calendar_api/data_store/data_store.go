package data_store

import (
	"errors"
	"sync"
	"github.com/t1ery/wb_l2/calendar_api/domain"
)

// DataStore представляет хранилище событий.
type DataStore struct {
	events []domain.Event
	mu     sync.Mutex
}

// NewDataStore создает новое хранилище данных.
func NewDataStore() *DataStore {
	return &DataStore{
		events: make([]domain.Event, 0),
	}
}

// CreateEvent создает новое событие и добавляет его в хранилище.
func (ds *DataStore) CreateEvent(event domain.Event) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	// Присваиваем уникальный ID
	event.ID = len(ds.events) + 1

	// Добавляем событие в хранилище
	ds.events = append(ds.events, event)
}

// GetEventByID возвращает событие по его ID.
func (ds *DataStore) GetEventByID(id int) (domain.Event, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for _, event := range ds.events {
		if event.ID == id {
			return event, nil
		}
	}

	return domain.Event{}, errors.New("событие не найдено")
}

// UpdateEvent обновляет событие в хранилище по его ID.
func (ds *DataStore) UpdateEvent(id int, updatedEvent domain.Event) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for i, event := range ds.events {
		if event.ID == id {
			// Обновляем событие
			ds.events[i] = updatedEvent
			return nil
		}
	}

	return errors.New("событие не найдено")
}

// DeleteEvent удаляет событие из хранилища по его ID.
func (ds *DataStore) DeleteEvent(id int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for i, event := range ds.events {
		if event.ID == id {
			// Удаляем событие из среза
			ds.events = append(ds.events[:i], ds.events[i+1:]...)
			return nil
		}
	}

	return errors.New("событие не найдено")
}

// GetEventsForDay возвращает события на указанный день.
func (ds *DataStore) GetEventsForDay(date string) []domain.Event {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	var eventsInDay []domain.Event

	for _, event := range ds.events {
		if event.Date == date {
			eventsInDay = append(eventsInDay, event)
		}
	}

	return eventsInDay
}


// GetEventsForWeek возвращает события в указанном диапазоне недель.
func (ds *DataStore) GetEventsForWeek(startDate, endDate string) []domain.Event {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	var eventsInWeek []domain.Event

	for _, event := range ds.events {
		if event.Date >= startDate && event.Date <= endDate {
			eventsInWeek = append(eventsInWeek, event)
		}
	}

	return eventsInWeek
}

// GetEventsForMonth возвращает события в указанном месяце и году.
func (ds *DataStore) GetEventsForMonth(month, year string) []domain.Event {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	var eventsInMonth []domain.Event

	for _, event := range ds.events {
		if event.Date[:7] == year+"-"+month {
			eventsInMonth = append(eventsInMonth, event)
		}
	}

	return eventsInMonth
}
