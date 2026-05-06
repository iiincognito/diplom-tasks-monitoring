package task_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iiincognito/diplom-tasks-monitoring/internal/features/task/service"
)

func (h *TaskHTTPHandler) UpdateTask(rw http.ResponseWriter, r *http.Request) {
	var req service.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("DEBUG HANDLER JSON ERROR: %v\n", err)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
		return
	}
	fmt.Printf("DEBUG UpdateTask REQUEST: %+v\n", req)

	resp, _ := h.taskService.UpdateTask(&req)
	fmt.Printf("DEBUG UpdateTask RESPONSE: %+v\n", resp)

	rw.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(rw).Encode(resp)
}
