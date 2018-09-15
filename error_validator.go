package valgo

import (
	"encoding/json"
	"fmt"
)

type ErrorItem struct {
	Name     string   `json:"name"`
	Messages []string `json:"messages"`
}

type ErrorValidator struct {
	items        []*ErrorItem
	currentError *ErrorItem
}

func (e *ErrorValidator) Error() string {
	count := len(e.items)
	if count == 1 {
		return fmt.Sprintf("There is 1 error")
	} else {
		return fmt.Sprintf("There are %v errors", count)
	}
}

func (e *ErrorValidator) Items() []ErrorItem {
	items := []ErrorItem{}
	for _, item := range e.items {
		items = append(items, *item)
	}
	if e.currentError != nil {
		items = append(items, *e.currentError)
	}
	return items
}

func (e *ErrorValidator) MarshalJSON() ([]byte, error) {
	items := map[string][]string{}
	for _, item := range e.items {
		items[item.Name] = item.Messages
	}
	if e.currentError != nil {
		items[e.currentError.Name] = e.currentError.Messages
	}
	return json.Marshal(struct {
		Items map[string][]string `json:"errors"`
	}{
		Items: items,
	})
}
