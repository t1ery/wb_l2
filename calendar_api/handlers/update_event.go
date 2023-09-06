package handlers

import (
	"fmt"
	"net/http"
	"github.com/t1ery/wb_l2/calendar_api/utils"
	"github.com/t1ery/wb_l2/calendar_api/data_store"
	"strconv"
)

func UpdateEventHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получите ID обновляемого события
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

	// Получаем данные для обновления из запроса
	updatedEvent, err := utils.ValidateCreateEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Выводим данные для отладки
	fmt.Printf("Обновляем событие с ID: %d\n", id)
	fmt.Printf("Новые данные: %+v\n", updatedEvent)

	// Обновляем событие в хранилище
	err = dataStore.UpdateEvent(id, updatedEvent)
	if err != nil {
		fmt.Printf("Ошибка при обновлении события: %v\n", err)
		http.Error(w, "Не удалось обновить событие", http.StatusInternalServerError)
		return
	}

	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Событие обновлено успешно"}
	utils.SendJSONResponse(w, http.StatusOK, response)
}
