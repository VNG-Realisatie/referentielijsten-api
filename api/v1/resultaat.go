package v1

import (
	"fmt"
	"github.com/VNG-Realisatie/referentielijsten-api/middleware"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"github.com/VNG-Realisatie/referentielijsten-api/validator"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ListResultaten returns a list of resultaten taken from the data directory
func (r *ReferentielijstenHandler) ListResultaten(w http.ResponseWriter, req *http.Request) {

	// swagger:route GET /resultaten resultaten resultaten_list
	//
	//  Ontsluit de selectielijst resultaten.
	//  Procestypen worden gerefereerd in zaaktypecatalogi - bij het configureren
	//  van een zaaktype wordt aangegeven welk procestype van toepassing is, zodat
	//  het archiefregime van zaken bepaald kan worden.
	//  Zie https://vng.nl/files/vng/20170706-selectielijst-gemeenten-intergemeentelijke-organen-2017.pdf
	//  voor de bron van de inhoud.
	//
	//	Consumes:
	//	- application/json
	//
	//	Produces:
	//	- application/json
	//
	//   Parameters:
	//     + name: proces_type
	//       in: query
	//       description: proces_type
	//       required: false
	//       type: string
	//     + name: page
	//       in: query
	//       description: Een pagina binnen de gepagineerde set resultaten.
	//       required: false
	//       type: integer
	//
	//	Schemes: http, https
	//
	//	Responses:
	//	  200: Resultaten
	//	  400: ValidatieFout
	//	  404: Fout
	// 	  405: Fout

	scheme := "http://"
	if req.TLS != nil || req.Host == productionHostName || req.Host == testHostName {
		scheme = "https://"
	}

	log.Printf("req.TLS: %v", req.TLS)

	pageId := req.URL.Query().Get(PageFromParam)
	procesType := req.URL.Query().Get(ProcesTypeFromParam)

	if procesType != "" {
		validation := validator.Url(procesType, ProcesTypeFromParam)
		if validation != nil {
			ref, _ := url.JoinPath(scheme, req.Host, "ref/fouten/ValidationError/")

			validation := models.ValidationError{
				Type:     ref,
				Code:     "invalid",
				Title:    "Invalid input",
				Status:   http.StatusBadRequest,
				Detail:   "",
				Instance: middleware.CreateUUID(),
				InvalidParams: []models.FieldValidationError{
					*validation,
				},
			}

			middleware.ResponseWithCode(w, http.StatusBadRequest, validation)
			return
		}
	}

	if pageId == "" {
		pageId = "1"
	}

	parsedId, err := strconv.Atoi(pageId)
	if err != nil {
		ref, _ := url.JoinPath(scheme, req.Host, "ref/fouten/ValidationError/")

		field := []models.FieldValidationError{
			{
				Name:   "page",
				Code:   "bad_request",
				Reason: "page cannot be parsed",
			},
		}
		validation := models.ValidationError{
			Type:          ref,
			Code:          "invalid",
			Title:         "Invalid input",
			Status:        http.StatusBadRequest,
			Detail:        "",
			Instance:      middleware.CreateUUID(),
			InvalidParams: field,
		}

		middleware.ResponseWithCode(w, http.StatusBadRequest, validation)
		return
	}

	var response models.Resultaten
	var start int
	finish := parsedId*100 - 1

	response.Count = len(r.Data.Resultaten)
	if parsedId*100 > response.Count {
		finish = response.Count - 1
	} else {
		u, _ := url.JoinPath(scheme, req.Host, req.URL.Path, fmt.Sprintf("?page=%v", parsedId+1))
		response.Next = u
	}

	if parsedId != 1 {
		u, _ := url.JoinPath(scheme, req.Host, req.URL.Path, fmt.Sprintf("?page=%v", parsedId-1))
		response.Previous = u
		start = parsedId*100 - 100
	} else {
		start = 0
	}

	if start > finish {
		u, _ := url.JoinPath(scheme, req.Host, "ref/fouten/NotFound/")
		notFound := models.HTTPError{
			Type:     u,
			Code:     "not_found",
			Title:    "Pagina niet gevonden",
			Status:   http.StatusNotFound,
			Detail:   "Ongeldige pagina.",
			Instance: middleware.CreateUUID(),
		}
		middleware.ResponseWithCode(w, http.StatusNotFound, notFound)
		return
	}

	resultsToBeUsed := r.Data.Resultaten[start:finish]

	for _, result := range resultsToBeUsed {
		u, _ := url.JoinPath(scheme, req.Host, req.URL.Path, result.URL)
		result.URL = u
		procesTypeUUID := result.ProcesType
		p, _ := url.JoinPath(scheme, req.Host, "api/v1/procestypen", result.ProcesType)
		result.ProcesType = p

		if procesType != "" {
			if strings.Contains(procesType, procesTypeUUID) {
				response.Results = append(response.Results, result)
			}
		} else {
			response.Results = append(response.Results, result)
		}
	}

	middleware.ResponseWithCode(w, http.StatusOK, response)
}

// GetResultaat returns a specific resultaat taken from the data directory
func (r *ReferentielijstenHandler) GetResultaat(w http.ResponseWriter, req *http.Request) {

	// swagger:route GET /resultaten/{uuid} resultaten resultaten_retrieve
	//
	//  Ontsluit de selectielijst resultaten.
	//  Procestypen worden gerefereerd in zaaktypecatalogi - bij het configureren
	//  van een zaaktype wordt aangegeven welk procestype van toepassing is, zodat
	//  het archiefregime van zaken bepaald kan worden.
	//  Zie https://vng.nl/files/vng/20170706-selectielijst-gemeenten-intergemeentelijke-organen-2017.pdf
	//  voor de bron van de inhoud.
	//
	//	Consumes:
	//	- application/json
	//
	//	Produces:
	//	- application/json
	//
	//   Parameters:
	//     + name: uuid
	//       in: path
	//       description: uuid van de resultaattypeomschrijving
	//       required: true
	//       type: string
	//       format: uuid
	//		 title: uuid
	//
	//	Schemes: http, https
	//
	//	Responses:
	//	  200: Resultaat
	//	  404: Fout
	// 	  405: Fout

	pathParams := mux.Vars(req)
	uuid := pathParams[UUIDFromParam]
	scheme := "https://"
	if req.TLS == nil {
		scheme = "http://"
	}

	for _, result := range r.Data.Resultaten {
		if result.URL == uuid {
			u, _ := url.JoinPath(scheme, req.Host, req.RequestURI)
			result.URL = u
			p, _ := url.JoinPath(scheme, req.Host, "api/v1/procestypen", result.ProcesType)
			result.ProcesType = p
			middleware.ResponseWithCode(w, http.StatusOK, result)
			return
		}
	}

	u, _ := url.JoinPath(scheme, req.Host, "ref/fouten/NotFound/")
	notFound := models.HTTPError{
		Type:     u,
		Code:     "not_found",
		Title:    "Resultaat niet gevonden",
		Status:   http.StatusNotFound,
		Detail:   fmt.Sprintf("UUID: %s niet gevonden", uuid),
		Instance: middleware.CreateUUID(),
	}
	middleware.ResponseWithCode(w, http.StatusNotFound, notFound)
}
