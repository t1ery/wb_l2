package handlers

import (
	"net/http"
	"github.com/t1ery/wb_l2/calendar_api/utils"
	"github.com/t1ery/wb_l2/calendar_api/data_store"
)


func CreateEventHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Распарсите и валидируйте параметры запроса
	newEvent, err := utils.ValidateCreateEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Добавляем событие в хранилище
	dataStore.CreateEvent(newEvent)


	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Событие создано успешно"}
	utils.SendJSONResponse(w, http.StatusOK, response)
}
