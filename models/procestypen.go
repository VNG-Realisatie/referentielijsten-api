package models

type ProcesTypen []ProcesTypeElement

// swagger:model ProcesType
type ProcesTypeElement struct {
	// Interne referentie naar de url
	//
	// read only: true
	// required: true
	// example: http://referentielijsten-api.internal/api/v1/651a1b5b-f84f-4c73-9151-4d485c7dcb99
	URL string `json:"url"`
	//  Het jaartal waartoe dit ProcesType behoort
	//
	// max: 2147483647
	// min: -2147483648
	//
	// required: true
	// example: 6
	Nummer int32 `json:"nummer"`
	//  Nummer van de selectielijstcategorie
	//
	// max: 2147483647
	// min: -2147483648
	//
	// required: true
	// example: 2020
	Jaar int32 `json:"jaar"`
	//  Benaming van het procestype
	//
	//  max length: 100
	//
	// required: true
	// example: Producten en diensten leveren
	Naam string `json:"naam"`
	//  Omschrijving van het procestype
	//
	//  max length: 300
	//
	// required: true
	// example: Het door het orgaan leveren van (publieke) producten of diensten
	Omschrijving string `json:"omschrijving"`
	//  Toelichting van het procestype
	//
	// required: true
	// example: Voor het verstrekken van een product of dienst hoeft de aanvrager niet aan een bepaalde situatie te voldoen, want in dat geval is het procestype 'Voorzieningen verstrekken' van toepassing. Het product of de dienst houdt niet het doen of laten van het orgaan in, want in dat geval valt het onder het procestype 'Verzoeken behandelen'.
	Toelichting string `json:"toelichting"`
	//  Object waar de uitvoering van het proces op van toepassing
	//is en waarvan de bestaans- of geldigheidsduur eventueel van belang is
	//bij het bepalen van de start van de bewaartermijn
	//
	//  max length: 80
	//
	// required: true
	// example: Het geleverde product of de geleverde dienst
	Procesobject string `json:"procesobject"`
}
