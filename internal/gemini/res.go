package gemini

import "encoding/json"

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
		FinishReason string `json:"finishReason"`
	} `json:"candidates"`
}

type GeminiErrorResponse struct {
	Error struct {
		Details json.RawMessage `json:"details"`
		Message string          `json:"message"`
		Status  string          `json:"status"`
		Code    int             `json:"code"`
	} `json:"error"`
}
