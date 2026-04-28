package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", nil
	}

	startDate, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", fmt.Errorf("error parsing start date: %w", err)
	}

	repeat = strings.TrimSpace(repeat)

	switch {
	case strings.HasPrefix(repeat, "d "):
		daysStr := strings.TrimSpace(repeat[1:])
		days, err := strconv.Atoi(daysStr)
		if err != nil {
			return "", fmt.Errorf("error converting repeat days to int: %w", err)
		}
		if days > 400 {
			return "", fmt.Errorf("repeat days should be less than or equal to 400")
		}

		future := startDate.AddDate(0, 0, days)
		for !future.After(now) {
			future = future.AddDate(0, 0, days)
		}
		return future.Format("20060102"), nil

	case repeat == "y":
		future := startDate.AddDate(1, 0, 0)
		for !future.After(now) {
			future = future.AddDate(1, 0, 0)
		}
		return future.Format("20060102"), nil

	default:
		return "", fmt.Errorf("unsupported repeat rule: %s", repeat)
	}

}
