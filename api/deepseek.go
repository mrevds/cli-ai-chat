package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string     `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type DeepSeekClient struct {
	apiKey string
}

func NewDeepSeekClient(apiKey string) *DeepSeekClient {
	return &DeepSeekClient{apiKey: apiKey}
}

func (c *DeepSeekClient) SendMessage(userInput string) (string, error) {
	url := "https://api.deepseek.com/chat/completions"

	messages := []Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: userInput},
	}

	requestBody := ChatRequest{
		Model:    "deepseek-chat",
		Messages: messages,
		Stream:   false,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("ошибка маршалинга запроса: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API вернуло ошибку: %s, тело ответа: %s", resp.Status, string(body))
	}

	var chatResponse ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResponse); err != nil {
		return "", fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	if len(chatResponse.Choices) == 0 {
		return "", fmt.Errorf("пустой ответ от API")
	}

	return chatResponse.Choices[0].Message.Content, nil
}