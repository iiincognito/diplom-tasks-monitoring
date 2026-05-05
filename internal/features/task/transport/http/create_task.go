package task_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iiincognito/diplom-tasks-monitoring/internal/core/domain"
)

func (h *TaskHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	var req domain.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("DEBUG HANDLER JSON ERROR: %v\n", err)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(domain.CreateTaskResponse{Error: err.Error()})
		return
	}
	fmt.Printf("DEBUG HANDLER REQUEST: %+v\n", req)

	resp, _ := h.taskService.CreateTask(&req)
	fmt.Printf("DEBUG HANDLER RESPONSE: %+v\n", resp)

	rw.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(rw).Encode(resp)
}
