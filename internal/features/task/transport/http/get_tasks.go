package task_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *TaskHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	fmt.Printf("DEBUG GetTasks search: %s\n", search)

	resp, _ := h.taskService.GetTasks(search)
	fmt.Printf("DEBUG GetTasks response: %+v\n", resp)

	rw.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(rw).Encode(resp)
}
