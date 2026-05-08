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
		daysStr := strings.TrimSpace(repeat[2:])
		days, err := strconv.Atoi(daysStr)
		if err != nil {
			return "", fmt.Errorf("error converting repeat days to int: %w", err)
		}
		if days <= 0 || days > 400 {
			return "", fmt.Errorf("repeat days should be between 1 and 400")
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

	case strings.HasPrefix(repeat, "w "):
		return nextWeekdayDate(now, startDate, repeat[2:])

	case strings.HasPrefix(repeat, "m "):
		return nextMonthDate(now, startDate, repeat[2:])

	default:
		return "", fmt.Errorf("unsupported repeat rule: %s", repeat)
	}
}

func nextWeekdayDate(now, startDate time.Time, rule string) (string, error) {
	days := strings.Split(strings.TrimSpace(rule), ",")
	allowedDays := make(map[int]bool)

	for _, d := range days {
		d = strings.TrimSpace(d)
		day, err := strconv.Atoi(d)
		if err != nil {
			return "", fmt.Errorf("invalid weekday: %s", d)
		}
		if day < 1 || day > 7 {
			return "", fmt.Errorf("weekday should be between 1 and 7")
		}
		allowedDays[day] = true
	}

	future := startDate.AddDate(0, 0, 1)
	if future.Before(now) || future.Equal(now) {
		future = now.AddDate(0, 0, 1)
	}

	for {
		weekday := int(future.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		if allowedDays[weekday] {
			return future.Format("20060102"), nil
		}
		future = future.AddDate(0, 0, 1)
	}
}

func nextMonthDate(now, startDate time.Time, rule string) (string, error) {
	parts := strings.Fields(strings.TrimSpace(rule))
	if len(parts) == 0 || len(parts) > 2 {
		return "", fmt.Errorf("invalid month rule format")
	}

	daysStr := strings.Split(parts[0], ",")
	monthDays := make([]int, 0)
	for _, d := range daysStr {
		d = strings.TrimSpace(d)
		day, err := strconv.Atoi(d)
		if err != nil {
			return "", fmt.Errorf("invalid month day: %s", d)
		}
		if day != -1 && day != -2 && (day < 1 || day > 31) {
			return "", fmt.Errorf("month day should be -2, -1, or between 1 and 31")
		}
		monthDays = append(monthDays, day)
	}

	allowedMonths := make(map[int]bool)
	if len(parts) == 2 {
		monthsStr := strings.Split(parts[1], ",")
		for _, m := range monthsStr {
			m = strings.TrimSpace(m)
			month, err := strconv.Atoi(m)
			if err != nil {
				return "", fmt.Errorf("invalid month: %s", m)
			}
			if month < 1 || month > 12 {
				return "", fmt.Errorf("month should be between 1 and 12")
			}
			allowedMonths[month] = true
		}
	}

	checkDate := startDate.AddDate(0, 0, 1)
	if checkDate.Before(now) || checkDate.Equal(now) {
		checkDate = now.AddDate(0, 0, 1)
	}

	for {
		currentYear, currentMonth, _ := checkDate.Date()
		lastDay := time.Date(currentYear, currentMonth+1, 0, 0, 0, 0, 0, time.UTC).Day()

		if len(allowedMonths) == 0 || allowedMonths[int(currentMonth)] {
			currentDay := checkDate.Day()
			bestDay := -1

			for _, targetDay := range monthDays {
				var actualDay int
				if targetDay == -1 {
					actualDay = lastDay
				} else if targetDay == -2 {
					actualDay = lastDay - 1
				} else {
					actualDay = targetDay
				}

				if actualDay > lastDay {
					continue
				}

				if actualDay >= currentDay {
					if bestDay == -1 || actualDay < bestDay {
						bestDay = actualDay
					}
				}
			}

			if bestDay != -1 {
				result := time.Date(currentYear, currentMonth, bestDay, 0, 0, 0, 0, time.UTC)
				if !result.Before(now) && !result.Equal(now) {
					return result.Format("20060102"), nil
				}
			}
		}

		checkDate = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
	}
}
