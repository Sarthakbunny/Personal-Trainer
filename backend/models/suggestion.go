package models

// SuggestionRequest represents the request structure for generating suggestions
type SuggestionRequest struct {
	Goal      string `json:"goal"`
	DietType  string `json:"diet_type"`
	Intensity string `json:"intensity"`
}
