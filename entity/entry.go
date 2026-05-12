package entity

import (
	"strings"
	"time"
)

type Entry struct {
	URL       string
	VisitedAt time.Time
	BookMark  bool
}

func GetDomain(url string) string {
	start := 0
	if ind := strings.Index(url, "://"); ind != -1 {
		start = ind + 3
	}

	end := len(url)
	if ind := strings.Index(url[start:], "/"); ind != -1 {
		end = start + ind
	}
	return url[start:end]
}
