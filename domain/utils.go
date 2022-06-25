package domain

import "time"

// StrToTime ...
func StrToTime(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fallback, _ := time.Parse(time.RFC3339, FallBackTime)
		return fallback
	}
	return t
}
