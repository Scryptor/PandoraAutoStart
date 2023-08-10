package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Session struct {
	MessageId string `json:"message_id"`
	SessionId string `json:"session_id"`
	SkillId   string `json:"skill_id"`
}
type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
	Version    string      `json:"version"`
	Session    Session     `json:"session"`
	Response   response    `json:"response"`
}
type response struct {
	Text       string `json:"text"`
	EndSession string `json:"end_session"`
}

func Handler(ctx context.Context) (*Response, error) {
	var StatusMessage string
	StatusMessage = "Гриня запущен. Шершавого асфальта босс"
	if err := SendTG(); err != nil {
		StatusMessage = "Не удалось запустить Гриню"
	}
	return &Response{
		StatusCode: 200,
		Body:       "Hello, world!",
		Response: response{
			Text:       StatusMessage,
			EndSession: "true",
		},
		Version: "1.0",
	}, nil
}
func SendTG() error {
	ClientH := http.Client{Timeout: 10 * time.Second}
	data := url.Values{}

	req, err := http.NewRequest("GET", "http://ссылка на ваш url", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	resp, err := ClientH.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if !json.Valid(body) {
		return err
	}
	return nil
}
