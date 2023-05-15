package v1

import (
	"fmt"
	"github.com/VNG-Realisatie/referentielijsten-api/middleware"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
)

// ListProcesTypen returns a list of procestypen taken from the data directory
func (r *ReferentielijstenHandler) ListProcesTypen(w http.ResponseWriter, req *http.Request) {

	// swagger:route GET /procestypen procestypen procestypen_list
	//
	//  Ontsluit de selectielijst procestypen.
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
	//     + name: jaar
	//       in: query
	//       description: jaar
	//       required: false
	//       type: string
	//
	//	Schemes: http, https
	//
	//	Responses:
	//	  200: ProcesType
	//	  400: ValidatieFout
	// 	  405: Fout

	year := req.URL.Query().Get(JaarFromParam)
	var parsedYear int32

	var response models.ProcesTypen

	scheme := getScheme(req)

	if year != "" {
		convertedYear, err := strconv.Atoi(year)
		if err != nil {
			ref, _ := url.JoinPath(scheme, req.Host, "ref/fouten/ValidationError/")
			invalidParams := []models.FieldValidationError{
				{
					Name:   "jaar",
					Code:   "invalid",
					Reason: fmt.Sprintf("Jaar: %s is geen valide input", year),
				},
			}

			validation := models.ValidationError{
				Type:          ref,
				Code:          "invalid",
				Title:         "Invalid input",
				Status:        http.StatusBadRequest,
				Detail:        "",
				Instance:      middleware.CreateUUID(),
				InvalidParams: invalidParams,
			}

			middleware.ResponseWithCode(w, http.StatusBadRequest, validation)
			return
		}
		parsedYear = int32(convertedYear)

	}

	for _, result := range r.Data.ProcesTypen {
		u, _ := url.JoinPath(scheme, req.Host, req.URL.Path, result.URL)
		result.URL = u
		if year != "" {
			if parsedYear == result.Jaar {
				response = append(response, result)
			}
		} else {
			response = append(response, result)
		}
	}

	middleware.ResponseWithCode(w, http.StatusOK, response)
}

// GetProcesType returns a specific procestypen taken from the data directory
func (r *ReferentielijstenHandler) GetProcesType(w http.ResponseWriter, req *http.Request) {

	// swagger:route GET /procestypen/{uuid} procestypen procestypen_retrieve
	//
	//  Ontsluit de selectielijst procestypen.
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
	//	  200: ProcesType
	//	  400: ValidatieFout
	//	  404: Fout
	// 	  405: Fout

	pathParams := mux.Vars(req)
	uuid := pathParams[UUIDFromParam]
	scheme := getScheme(req)

	for _, result := range r.Data.ProcesTypen {
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
		Title:    "ProcesType niet gevonden",
		Status:   http.StatusNotFound,
		Detail:   fmt.Sprintf("UUID: %s niet gevonden", uuid),
		Instance: middleware.CreateUUID(),
	}
	middleware.ResponseWithCode(w, http.StatusNotFound, notFound)
}
