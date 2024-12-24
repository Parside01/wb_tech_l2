package transport

import (
	"encoding/json"
	"io"
	"net/http"
)

type GetRequestData struct {
	UserID string `json:"user_id"`
	Date   string `json:"date"`
}

func GetDataFromRequest(r *http.Request) (*GetRequestData, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	return &GetRequestData{
		UserID: r.PostForm.Get("user_id"),
		Date:   r.PostForm.Get("date"),
	}, nil
}

type PostRequestData struct {
	UserID      string `json:"user_id"`
	Date        string `json:"date"`
	Title       string `json:"title"`
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
