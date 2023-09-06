package handlers

import (
	"net/http"
	"github.com/t1ery/wb_l2/calendar_api/data_store"
	"github.com/t1ery/wb_l2/calendar_api/utils"
)

func EventsForWeekHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получите начальную и конечную дату недели из параметров запроса
	startDate := r.FormValue("start_date")
	endDate := r.FormValue("end_date")
	if startDate == "" || endDate == "" {
		http.Error(w, "Отсутствует параметр начальной или конечной даты", http.StatusBadRequest)
		return
	}

	// Найдите события на указанную неделю
	events := dataStore.GetEventsForWeek(startDate, endDate)

	// Возвращаем JSON-ответ с найденными событиями
	utils.SendJSONResponse(w, http.StatusOK, events)
}
