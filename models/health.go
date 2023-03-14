package models

// swagger:model
type Health struct {
	// Als de api online is is dit altijd true
	//
	// required: true
	Healthy bool `json:"healthy"`
	// Tijd waarop de message is gegenereerd
	//
	// required: true
	Time string `json:"time"`
}
