package task_http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *TaskHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("DEBUG DeleteTask id: %d\n", id)
	resp, _ := h.taskService.DeleteTask(id)
	fmt.Printf("DEBUG DeleteTask response: %+v\n", resp)

	rw.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(rw).Encode(resp)
}
