package v1

import (
	"fmt"
	"github.com/VNG-Realisatie/referentielijsten-api/middleware"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

// ListResultaattypeomschrijvingen returns a list of resultaattypeomschrijvingen taken from the data directory
func (r *ReferentielijstenHandler) ListResultaattypeomschrijvingen(w http.ResponseWriter, req *http.Request) {
	// swagger:route GET /resultaattypeomschrijvingen resultaattypeomschrijvingen resultaattypeomschrijvinggeneriek_list
	//
	// Raadpleeg de generieke resultaattypeomschrijvingen.
	//
	//	Consumes:
	//	- application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Schemes: http, https
	//
	//	Responses:
	//	  200: ResultaattypeOmschrijvingGeneriek
	// 	  405: Fout

	scheme := getScheme(req)

	//TODO remove
	log.Print(req)
	log.Printf("TLS: %v", req.TLS)

	var response models.ResultaattypeOmschrijvingen

	for _, result := range r.Data.ResultaattypenOmscrhijvingen {

		u, _ := url.JoinPath(scheme, req.Host, req.URL.Path, result.URL)
		result.URL = u
		response = append(response, result)
	}

	middleware.ResponseWithCode(w, http.StatusOK, response)
}

// GetResultaattypeomschrijving returns a specific resultaattypeomschrijving taken from the data directory
func (r *ReferentielijstenHandler) GetResultaattypeomschrijving(w http.ResponseWriter, req *http.Request) {

	// swagger:route GET /resultaattypeomschrijvingen/{uuid} resultaattypeomschrijvingen resultaattypeomschrijvinggeneriek_retrieve
	//
	// Raadpleeg de generieke resultaattypeomschrijvingen.
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
	//	  200: ResultaattypeOmschrijvingGeneriek
	//	  400: FieldValidationError
	//	  404: Fout
	// 	  405: Fout

	pathParams := mux.Vars(req)
	uuid := pathParams[UUIDFromParam]
	scheme := getScheme(req)

	for _, result := range r.Data.ResultaattypenOmscrhijvingen {
		if result.URL == uuid {
			u, _ := url.JoinPath(scheme, req.Host, req.URL.Path)
			result.URL = u
			middleware.ResponseWithCode(w, http.StatusOK, result)
			return
		}
	}

	u, _ := url.JoinPath(scheme, req.Host, "ref/fouten/NotFound/")
	notFound := models.HTTPError{
		Type:     u,
		Code:     "not_found",
		Title:    "Resultaatomschrijving niet gevonden",
		Status:   http.StatusNotFound,
		Detail:   fmt.Sprintf("UUID: %s niet gevonden", uuid),
		Instance: middleware.CreateUUID(),
	}
	middleware.ResponseWithCode(w, http.StatusNotFound, notFound)
}
