package entity

import (
	"encoding/json"
	"time"
)

type Event struct {
	Date        time.Time
	Title       string `json:"title,omitempty"` // В задании ничего не было про id для event, так что пусть уникальным будет заголовок.
	Description string `json:"description"`     // Может быть пустым.
}

// UnmarshalJSON Та самая вспомогательная функция для десериализации, заодно валидации.
// Не используются(.
func (e *Event) UnmarshalJSON(data []byte) error {
	alias := &struct {
		Data string `json:"data,omitempty"`
		*Event
	}{
		Event: e,
	}

	if err := json.Unmarshal(data, alias); err != nil {
		return err
	}

	if alias.Data == "" {
		e.Date = time.Unix(0, 0)
		return nil
	}

	parsed, err := time.Parse(time.RFC3339, alias.Data)
	if err != nil {
		return err
	}

	e.Date = parsed
	return nil
}

func (e *Event) MarshalJSON() ([]byte, error) {
	alias := &struct {
		Data string `json:"data"`
		*Event
	}{
		Event: e,
		Data:  e.Date.Format(time.RFC3339),
	}
	return json.Marshal(alias)
}
