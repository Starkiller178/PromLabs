package tests

import (
	"lab2/entity"
	"testing"
)

func TestAdd(t *testing.T) {
	h := entity.New()
	h.Add("https://google.com", true)
	h.Add("https://github.com", true)
	h.Add("https://stackoverflow.com", false)
	if len(*h.GetEntries()) != 3 {
		t.Error("Add method is incorrect")
	}
}

func TestBack(t *testing.T) {
	h := entity.New()

	h.Add("https://google.com", true)
	h.Add("https://github.com", true)
	h.Add("https://stackoverflow.com", false)

	e, ok := h.Back()
	if !ok || e.URL != "https://github.com" {
		t.Errorf("Back method is incorrect, expected https://github.com, received %v", e.URL)
	}

	e, ok = h.Back()
	if !ok || e.URL != "https://google.com" {
		t.Errorf("Back method is incorrect, expected https://google.com, received %v", e.URL)
	}

	_, ok = h.Back()
	if ok {
		t.Error("should not go back")
	}
}

func TestForward(t *testing.T) {
	h := entity.New()

	h.Add("https://google.com", true)
	h.Add("https://github.com", true)
	h.Add("https://stackoverflow.com", false)

	h.Back()
	h.Back()

	e, ok := h.Forward()
	if !ok || e.URL != "https://github.com" {
		t.Errorf("Forward method is incorrect, expected https://github.com, received %v", e.URL)
	}

	e, ok = h.Forward()
	if !ok || e.URL != "https://stackoverflow.com" {
		t.Errorf("Forward method is incorrect, expected https://stackoverflow.com, received %v", e.URL)
	}

	_, ok = h.Forward()
	if ok {
		t.Error("should not go forward")
	}
}

func TestClear(t *testing.T) {
	h := entity.New()

	h.Add("https://google.com", true)
	h.Add("https://github.com", true)
	h.Add("https://stackoverflow.com", false)

	h.Clear()

	if len(*h.GetEntries()) != 0 || *h.GetCurrent() != -1 {
		t.Error("Clear method is incorrect")
	}

}

func TestGetDomain(t *testing.T) {
	domain := entity.GetDomain("https://github.com")
	if domain != "github.com" {
		t.Errorf("GetDomain method is incorrect, expected github.com, received %v", domain)
	}

	domain = entity.GetDomain("https://github.com/solutions")
	if domain != "github.com" {
		t.Errorf("GetDomain method is incorrect, expected github.com, received %v", domain)
	}
}

func TestSearchByDomain(t *testing.T) {
	h := entity.New()

	h.Add("https://google.com", true)
	h.Add("https://github.com", true)
	h.Add("https://google.com", false)

	res := h.SearchByDomain("google.com")

	cnt := 0
	for _, entry := range res {
		if entry.URL == "https://google.com" {
			cnt++
		}
	}

	if cnt != 2 {
		t.Error("SearchByDomain method is incorrect")
	}
}

func TestSaveLoad(t *testing.T) {
	h1 := entity.New()
	h1.Add("https://example.com", true)
	h1.Add("https://golang.org", false)

	err := h1.SaveToFile("test.history")
	if err != nil {
		t.Fatal(err)
	}

	h2 := entity.New()
	err = h2.LoadFromFile("test.history")
	if err != nil {
		t.Fatal(err)
	}

	if len(*h2.GetEntries()) != 2 {
		t.Error("Save/Load method is incorrect")
	}

}
