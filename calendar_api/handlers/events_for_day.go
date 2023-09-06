package handlers

import (
	"net/http"
	"github.com/t1ery/wb_l2/calendar_api/data_store"
	"github.com/t1ery/wb_l2/calendar_api/utils"
)

func EventsForDayHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получите дату из параметров запроса
	date := r.FormValue("date")
	if date == "" {
		http.Error(w, "Отсутствует параметр даты", http.StatusBadRequest)
		return
	}

	// Найдите события на указанную дату
	events := dataStore.GetEventsForDay(date)

	// Возвращаем JSON-ответ с найденными событиями
	utils.SendJSONResponse(w, http.StatusOK, events)
}
