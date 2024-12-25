package transport

import (
	"encoding/json"
	"github.com/Parside01/dev11/internal/entity"
	"io"
	"net/http"
	"strconv"
	"time"
)

type GetRequestData struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

func GetDataFromRequest(r *http.Request) (*GetRequestData, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(r.PostForm.Get("user_id"))
	if err != nil {
		return nil, err
	}

	return &GetRequestData{
		UserID: id,
		Date:   r.PostForm.Get("date"),
	}, nil
}

type PostRequestData struct {
	UserID      int    `json:"user_id,omitempty"`
	Date        string `json:"date,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
}

func PostDataFromRequest(r *http.Request) (*PostRequestData, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	data := &PostRequestData{}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func EventFromPostData(data *PostRequestData) (*entity.Event, error) {
	date, err := time.Parse(time.RFC3339, data.Date)
	if err != nil {
		return nil, err
	}

	event := &entity.Event{
		Date:        date,
		Title:       data.Title,
		Description: data.Description,
	}
	return event, nil
}
