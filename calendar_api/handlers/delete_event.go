package handlers

import (
	"net/http"
	"github.com/t1ery/wb_l2/calendar_api/data_store"
	"github.com/t1ery/wb_l2/calendar_api/utils"
	"strconv"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получите ID удаляемого события
	eventID := r.FormValue("id")
	if eventID == "" {
		http.Error(w, "Отсутствует идентификатор события", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "Неверный формат идентификатора события", http.StatusBadRequest)
		return
	}

	// Удалите событие из хранилища
	err = dataStore.DeleteEvent(id)
	if err != nil {
		http.Error(w, "Не удалось удалить событие", http.StatusInternalServerError)
		return
	}

	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Событие удалено успешно"}
	utils.SendJSONResponse(w, http.StatusOK, response)
}
