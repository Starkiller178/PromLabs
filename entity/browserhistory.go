package entity

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
	"time"
)

type BrowserHistory struct {
	entries []Entry
	current int
}

func New() *BrowserHistory {
	return &BrowserHistory{
		entries: []Entry{},
		current: -1,
	}
}

func (h *BrowserHistory) GetEntries() *[]Entry {
	return &h.entries
}

func (h *BrowserHistory) GetCurrent() *int {
	return &h.current
}

func (h *BrowserHistory) Add(url string, bookmark bool) {
	if h.current >= 0 && h.current < len(h.entries)-1 {
		h.entries = h.entries[:h.current+1]
	}

	entry := Entry{
		URL:       url,
		VisitedAt: time.Now(),
		BookMark:  bookmark,
	}

	h.entries = append(h.entries, entry)
	h.current = len(h.entries) - 1
}

func (h *BrowserHistory) Back() (Entry, bool) {
	if h.current == -1 || h.current == 0 {
		return Entry{}, false
	}
	h.current--
	return h.entries[h.current], true
}

func (h *BrowserHistory) Forward() (Entry, bool) {
	if h.current == len(h.entries)-1 {
		return Entry{}, false
	}
	h.current++
	return h.entries[h.current], true
}

func (h *BrowserHistory) Clear() {
	h.entries = []Entry{}
	h.current = -1
}

func (h *BrowserHistory) SearchByDomain(Domain string) []Entry {
	EntriesList := make([]Entry, 0, 10)

	for _, e := range h.entries {
		if GetDomain(e.URL) == Domain {
			EntriesList = append(EntriesList, e)
		}
	}

	return EntriesList
}

func (h *BrowserHistory) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, entry := range h.entries {

		jsonData, err := json.Marshal(entry)
		if err != nil {
			return err
		}

		line := base64.StdEncoding.EncodeToString(jsonData)

		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (h *BrowserHistory) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r\n")
		if line == "" {
			continue
		}

		jsonData, err := base64.StdEncoding.DecodeString(line)
		if err != nil {
			return err
		}
		var entry Entry
		if err := json.Unmarshal(jsonData, &entry); err != nil {
			return err
		}

		h.entries = append(h.entries, entry)
		h.current = len(h.entries) - 1
	}
	return scanner.Err()
}
