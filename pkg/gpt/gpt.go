package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Prompt struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type GPT3Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func getGPT3Response(promptText string) (string, error) {
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"
	prompt := &Prompt{
		Prompt:    promptText,
		MaxTokens: 40,
	}
	jsonData, err := json.Marshal(prompt)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_API_KEY")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response GPT3Response
	json.Unmarshal(body, &response)

	return response.Choices[0].Text, nil
}
