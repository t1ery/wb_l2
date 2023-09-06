package handlers

import (
	"net/http"
	"github.com/t1ery/wb_l2/calendar_api/data_store"
	"github.com/t1ery/wb_l2/calendar_api/utils"
)

func EventsForMonthHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получите месяц и год из параметров запроса
	month := r.FormValue("month")
	year := r.FormValue("year")
	if month == "" || year == "" {
		http.Error(w, "Отсутствует параметр месяца или года", http.StatusBadRequest)
		return
	}

	// Найдите события на указанную неделю
	events := dataStore.GetEventsForMonth(month, year)
	
	// Возвращаем JSON-ответ с найденными событиями
	utils.SendJSONResponse(w, http.StatusOK, events)
}

