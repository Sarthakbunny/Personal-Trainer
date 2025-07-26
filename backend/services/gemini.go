package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var GeminiAPIKey = os.Getenv("GEMINI_API_KEY")

type GeminiRequest struct {
	Contents []map[string]string `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func GenerateSuggestions(prompt string) (string, error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=" + GeminiAPIKey

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"role": "user",
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var gemResp GeminiResponse
	json.NewDecoder(res.Body).Decode(&gemResp)

	if len(gemResp.Candidates) > 0 {
		return gemResp.Candidates[0].Content.Parts[0].Text, nil
	}
	return "", fmt.Errorf("no suggestions found")
}
