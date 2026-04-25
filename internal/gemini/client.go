package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	APIKey string
	URL    string
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		URL:    "https://generativelanguage.googleapis.com/v1beta/models/gemini-flash-latest:generateContent",
	}
}

func (c *Client) LinkedInify(inputText string) (string, error) {
	reqBody := GeminiRequest{
		SystemInstruction: &SystemInstruction{
			Parts: []Part{
				{Text: "Kamu adalah mesin penerjemah LinkedIn Speak yang mengubah kalimat sederhana menjadi istilah profesional mewah. JANGAN memberikan kalimat pengantar, langsung berikan hasilnya."},
			},
		},
		Contents: []Content{
			{
				Parts: []Part{
					{Text: inputText},
				},
			},
		},
		GenerationConfig: GenerationConfig{
			Temperature: 0.7,
			// TopK:            40,
			TopP:            0.95,
			MaxOutputTokens: 300,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	fullURL := fmt.Sprintf("%s?key=%s", c.URL, c.APIKey)
	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respBody GeminiErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			return "", fmt.Errorf("gemini api error: status %d", resp.StatusCode)
		}
		return respBody.Error.Message, fmt.Errorf("gemini api error. status=%d, message=%s", resp.StatusCode, respBody.Error.Message)
	}

	var resBody GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return "", err
	}

	if len(resBody.Candidates) > 0 && len(resBody.Candidates[0].Content.Parts) > 0 {
		return resBody.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("empty response from gemini")
}
