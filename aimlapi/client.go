package aimlapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`  // "user", "system", "assistant"
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string     `json:"model"`    // "gpt-4o", "gpt-3.5-turbo", etc.
	Messages []Message `json:"messages"` // История диалога
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type AIMLAPIClient struct {
	BaseURL string
	APIKey  string
}

func NewAIMLAPIClient(apiKey string) *AIMLAPIClient {
	return &AIMLAPIClient{
		BaseURL: "https://api.aimlapi.com/v1",
		APIKey:  apiKey,
	}
}

func (c *AIMLAPIClient) Chat(model string, messages []Message) (string, error) {
	requestBody := ChatRequest{
		Model:    model,
		Messages: messages,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("ошибка кодирования запроса: %v", err)
	}

	req, err := http.NewRequest(
		"POST",
		c.BaseURL+"/chat/completions",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка запроса к AIMLAPI: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        return "", fmt.Errorf("AIMLAPI вернул ошибку: %s", resp.Status)
    }

	var response ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("пустой ответ от API")
	}

	return response.Choices[0].Message.Content, nil
}