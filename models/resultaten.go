package models

// swagger:model Resultaten
type Resultaten struct {
	// Totaal aantal resultaten
	//
	// required: false
	// example: 356
	Count int `json:"count"`
	// Volgende pagina indien beschikbaar
	//
	// required: false
	// example: http://localhost:8000/api/v1/resultaten?page=2
	// Extensions:
	// ---
	// x-nullable: true
	// swagger:strfmt uri
	Next string `json:"next"`
	// Vorige pagina indien beschikbaar
	//
	// required: false
	// example: http://localhost:8000/api/v1/resultaten?page=1
	// Extensions:
	// ---
	// x-nullable: true
	// swagger:strfmt uri
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

// swagger:model Resultaat
type Result struct {
	// Interne referentie naar de url
	//
	// read only: true
	// required: true
	// example: http://referentielijsten-api.internal/api/v1/651a1b5b-f84f-4c73-9151-4d485c7dcb99
	URL string `json:"url"`
	// procestype
	//
	// read only: true
	//
	// required: true
	// swagger:strfmt uri
	ProcesType string `json:"procesType"`
	//  Nummer van het resultaat. Dit wordt samengesteld met het procestype en generiek resultaat indien van toepassing.
	//
	// max: 2147483647
	// min: -2147483648
	//
	// required: true
	// example: 6
	Nummer int64 `json:"nummer"`
	// volledigNummer
	//
	// read only: true
	//
	// required: true
	VolledigNummer string `json:"volledigNummer"`
	// generiek
	//
	// read only: true
	//
	// required: true
	Generiek bool `json:"generiek"`
	// specifiek
	//
	// read only: true
	//
	// required: true
	Specifiek bool `json:"specifiek"`
	//  Benaming van het procestype
	//
	//  max length: 40
	//
	// required: true
	Naam string `json:"naam"`
	//  Omschrijving van het specifieke resultaat
	//
	//  max length: 150
	//
	// required: false
	Omschrijving string `json:"omschrijving"`
	//  "Voorbeeld: 'Risicoanalyse', 'Systeemanalyse' of verwijzing
	//  naar Wet- en regelgeving"
	//
	//  max length: 200
	//
	// required: true
	Herkomst string `json:"herkomst"`
	// swagger:allOf
	Waardering WaarderingEnum `json:"waardering"`
	// Uitleg bij mogelijke waarden:
	// `nihil` - Nihil
	// `bestaansduur_procesobject` - De bestaans- of geldigheidsduur van het procesobject.
	// `ingeschatte_bestaansduur_procesobject` - De ingeschatte maximale bestaans- of geldigheidsduur van het procesobject.
	// `vast_te_leggen_datum` - De tijdens het proces vast te leggen datum waarop de geldigheid van het procesobject komt te vervallen.
	// `samengevoegd_met_bewaartermijn` - De procestermijn is samengevoegd met de bewaartermijn."
	//
	// required: false
	Procestermijn ProcestermijnEnum `json:"procestermijn"`
	//  procestermijnWeergave
	//
	// read only: true
	//
	// required: true
	ProcestermijnWeergave string `json:"procestermijnWeergave"`
	//  bewaartermijn
	//
	// Extensions:
	// ---
	// x-nullable: true
	// required: false
	Bewaartermijn *string `json:"bewaartermijn"`
	//  Toelichting
	//
	//  max length: 150
	//
	// required: false
	Toelichting string `json:"toelichting"`
	// algemeen bestuur en inrichting organisatie
	//
	// read only: true
	//
	// required: false
	AlgemeenBestuurEnInrichtingOrganisatie bool `json:"algemeenBestuurEnInrichtingOrganisatie"`
	// bedrijfsvoering en personeel
	//
	// read only: true
	//
	// required: false
	BedrijfsvoeringEnPersoneel bool `json:"bedrijfsvoeringEnPersoneel"`
	// publieke informatie en registratie
	//
	// read only: true
	//
	// required: false
	PubliekeInformatieEnRegistratie bool `json:"publiekeInformatieEnRegistratie"`
	// burgerzaken
	//
	// read only: true
	//
	// required: false
	Burgerzaken bool `json:"burgerzaken"`
	// veiligheid
	//
	// read only: true
	//
	// required: false
	Veiligheid bool `json:"veiligheid"`
	// verkeer en vervoer
	//
	// read only: true
	//
	// required: false
	VerkeerEnVervoer bool `json:"verkeerEnVervoer"`
	// economie
	//
	// read only: true
	//
	// required: false
	Economie bool `json:"economie"`
	// onderwijs
	//
	// read only: true
	//
	// required: false
	Onderwijs bool `json:"onderwijs"`
	// sport, cultuur en recreatie
	//
	// read only: true
	//
	// required: false
	SportCultuurEnRecreatie bool `json:"sportCultuurEnRecreatie"`
	// sociaal domein
	//
	// read only: true
	//
	// required: false
	SociaalDomein bool `json:"sociaalDomein"`
	// volksgezondheid en milieu
	//
	// read only: true
	//
	// required: false
	VolksgezonheidEnMilieu bool `json:"volksgezonheidEnMilieu"`
	// VHROSV
	//
	// read only: true
	//
	// required: false
	Vhrosv bool `json:"vhrosv"`
	// Heffen belastingen etc.
	//
	// read only: true
	//
	// required: false
	HeffenBelastingen bool `json:"heffenBelastingen"`
	// alle taakgebieden
	//
	// read only: true
	//
	// required: false
	AlleTaakgebieden bool `json:"alleTaakgebieden"`
	//  Voorbeeld: '25 jaar', '30 jaar, '5 of 10 jaar'
	//
	//  max length: 20
	//
	// required: false
	ProcestermijnOpmerking string `json:"procestermijnOpmerking"`
}

// swagger:enum ProcestermijnEnum
type ProcestermijnEnum string

const (
	BestaansduurProcesobject     ProcestermijnEnum = "bestaansduur_procesobject"
	Nihil                        ProcestermijnEnum = "nihil"
	Procestermijn                ProcestermijnEnum = ""
	SamengevoegdMetBewaartermijn ProcestermijnEnum = "samengevoegd_met_bewaartermijn"
	VastTeLeggenDatum            ProcestermijnEnum = "vast_te_leggen_datum"
)

// swagger:enum WaarderingEnum
type WaarderingEnum string

const (
	BlijvendBewaren WaarderingEnum = "blijvend_bewaren"
	Vernietigen     WaarderingEnum = "vernietigen"
)
