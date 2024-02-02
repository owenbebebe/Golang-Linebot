package gpt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Prompt struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// GetGPT3Response is a function to get response from GPT-3.5-turbo model
func GetGPT3Response(promptText string) (string, error) {
	reqBody := fmt.Sprintf(`{
        "model": "gpt-3.5-turbo",
        "messages": [
            {
                "role": "system",
                "content": "You are lovely Taiwanese boyfriend."
            },
            {
                "role": "user",
                "content": "%s"
            }
        ],
		"max_tokens": 50
    }`, promptText)
	url := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	// replace <YOUR API KEY> with your OpenAI API key
	req.Header.Set("Authorization", "Bearer <YOUR_API_KEY>")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
	var response GPTResponse
	json.Unmarshal(body, &response)
	return response.Choices[0].Message.Content, nil
}
