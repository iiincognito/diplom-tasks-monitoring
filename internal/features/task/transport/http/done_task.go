package task_http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *TaskHTTPHandler) DoneTask(rw http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]string{"error": "Не указан идентификатор"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]string{"error": "Неверный идентификатор"})
		return
	}

	fmt.Printf("DEBUG DoneTask id: %d\n", id)
	resp, _ := h.taskService.DoneTask(id)
	fmt.Printf("DEBUG DoneTask response: %+v\n", resp)

	rw.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(rw).Encode(resp)
}
