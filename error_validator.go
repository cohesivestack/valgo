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

func (ev *ErrorValidator) Error() string {
	count := len(ev.items)
	if count == 1 {
		return fmt.Sprintf("There is 1 error")
	} else {
		return fmt.Sprintf("There are %v errors", count)
	}
}

func (ev *ErrorValidator) Items() []ErrorItem {
	items := []ErrorItem{}
	for _, item := range ev.items {
		items = append(items, *item)
	}
	if ev.currentError != nil {
		items = append(items, *ev.currentError)
	}
	return items
}

func (ev *ErrorValidator) MarshalJSON() ([]byte, error) {
	items := map[string][]string{}
	for _, item := range ev.items {
		items[item.Name] = item.Messages
	}
	if ev.currentError != nil {
		items[ev.currentError.Name] = ev.currentError.Messages
	}
	return json.Marshal(struct {
		Items map[string][]string `json:"errors"`
	}{
		Items: items,
	})
}
