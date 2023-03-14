package models

// Fout represents an error
//
// Formaat van HTTP 4xx en 5xx fouten.
//
// swagger:model Fout
type HTTPError struct {
	// URI referentie naar het type fout, bedoeld voor developers
	//
	// required: false
	Type string `json:"type"`
	// Systeemcode die het type fout aangeeft
	//
	// required: true
	Code string `json:"code"`
	// Generieke titel voor het type fout
	//
	// required: true
	Title string `json:"title"`
	// De HTTP status code
	//
	// required: true
	Status int64 `json:"status"`
	// Extra informatie bij de fout, indien beschikbaar
	//
	// required: true
	Detail string `json:"detail"`
	// URI met referentie naar dit specifiek voorkomen van de fout. Deze kan gebruikt worden in combinatie met server logs, bijvoorbeeld.
	//
	// required: true
	Instance string `json:"instance"`
}

// ValidatieFout represents an error
//
// Formaat van HTTP 4xx en 5xx fouten.
//
// swagger:model ValidatieFout
type ValidationError struct {
	// URI referentie naar het type fout, bedoeld voor developers
	//
	// required: false
	Type string `json:"type,omitempty"`
	// Systeemcode die het type fout aangeeft
	//
	// required: true
	Code string `json:"code"`
	// Generieke titel voor het type fout
	//
	// required: true
	Title string `json:"title"`
	//  De HTTP status code
	//
	// required: true
	Status int32 `json:"status"`
	// Extra informatie bij de fout, indien beschikbaar
	//
	// required: true
	Detail string `json:"detail"`
	// URI met referentie naar dit specifiek voorkomen van de fout. Deze kan gebruikt worden in combinatie met server logs, bijvoorbeeld.
	//
	// required: true
	Instance string `json:"instance"`
	// required: true
	InvalidParams []FieldValidationError `json:"invalidParams"`
}

// FieldValidationError represents an error
//
// Formaat van validatiefouten.
//
// swagger:model FieldValidationError
type FieldValidationError struct {
	//  Naam van het veld met ongeldige gegevens
	//
	// required: true
	Name string `json:"name"`
	// Systeemcode die het type fout aangeeft
	//
	// required: true
	Code string `json:"code"`

	//  Uitleg wat er precies fout is met de gegevens
	//
	// required: true
	Reason string `json:"reason"`
}
