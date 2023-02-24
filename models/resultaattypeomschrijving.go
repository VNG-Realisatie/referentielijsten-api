package models

type ResultaattypeOmschrijvingen []ResultaattypeOmschrijvingenElement

// swagger:model ResultaattypeOmschrijvingGeneriek
type ResultaattypeOmschrijvingenElement struct {
	//  Naam van het veld met ongeldige gegevens
	//
	// required: true
	URL string `json:"url"`
	//  Algemeen gehanteerde omschrijvingen van de aard van het resultaat
	//
	// required: true
	Omschrijving string `json:"omschrijving"`
	//  Nauwkeurige beschrijving van het generieke type resultaat.
	//
	// required: true
	Definitie string `json:"definitie"`
	//  Zinvolle toelichting bij de waarde van de generieke omschrijving
	//
	// required: false
	Opmerking string `json:"opmerking"`
}
