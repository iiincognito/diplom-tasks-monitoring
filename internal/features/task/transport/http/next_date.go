package task_http

import (
	"net/http"
	"time"

	"github.com/iiincognito/diplom-tasks-monitoring/internal/features/task/service"
)

func (h *TaskHTTPHandler) NextDate(rw http.ResponseWriter, r *http.Request) {
	nowStr := r.URL.Query().Get("now")
	dateStr := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	if dateStr == "" {
		http.Error(rw, "date parameter is required", http.StatusBadRequest)
		return
	}

	var now time.Time
	var err error

	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse("20060102", nowStr)
		if err != nil {
			http.Error(rw, "invalid now format, expected YYYYMMDD", http.StatusBadRequest)
			return
		}
	}

	nextDate, err := service.NextDate(now, dateStr, repeat)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(nextDate))
}
