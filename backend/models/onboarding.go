package models

// OnboardingData represents the data structure for user onboarding information
type OnboardingData struct {
	Goal     string `json:"goal"`
	Diet     string `json:"diet"`
	Level    string `json:"level"`
	Allergy  string `json:"allergy"`
	Timezone string `json:"timezone"`
}
