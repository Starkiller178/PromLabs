package main

import (
	"fmt"
	"lab2/entity"
)

func main() {
	h := entity.New()

	h.Add("https://google.com", true)
	h.Add("https://github.com", true)
	h.Add("https://stackoverflow.com", false)
	fmt.Printf("Current = %v, Entries = %v\n", *h.GetCurrent(), *h.GetEntries())

	e, _ := h.Back()
	fmt.Printf("Current = %v, entry = %v\n", *h.GetCurrent(), e)

	e, _ = h.Forward()
	fmt.Printf("Current = %v, entry = %v\n", *h.GetCurrent(), e)

	h.SaveToFile("History")

	h.Clear()
	fmt.Printf("Current = %v, Entries = %v\n", *h.GetCurrent(), *h.GetEntries())

	h.LoadFromFile("History")
	fmt.Printf("Current = %v, Entries = %v\n", *h.GetCurrent(), *h.GetEntries())

}
