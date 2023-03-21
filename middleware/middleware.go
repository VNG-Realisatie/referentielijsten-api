package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"gopkg.in/oauth2.v3/utils/uuid"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Adapter func(http.HandlerFunc) http.HandlerFunc

// Adapt Iterate over adapters and run them one by one
func Adapt(h http.HandlerFunc, adapters ...Adapter) http.HandlerFunc {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// LogRequestDetails logs the incoming requests
func LogRequestDetails() Adapter {
	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s route %s", r.Method, r.URL.Path)
			f(w, r)
		}
	}
}

// ValidateRestMethod middleware to validate proper methods
func ValidateRestMethod(method string) Adapter {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				scheme := "https://"
				if r.TLS == nil {
					scheme = "http://"
				}

				u, _ := url.JoinPath(scheme, r.Host, "ref/fouten/MethodNotAllowed/")
				notFound := models.HTTPError{
					Type:     u,
					Code:     "not_allowed",
					Title:    "Niet toegestaan",
					Status:   http.StatusNotFound,
					Detail:   fmt.Sprintf("Method: %s niet toegestaan", method),
					Instance: CreateUUID(),
				}
				ResponseWithCode(w, http.StatusMethodNotAllowed, notFound)
				return
			}
			f(w, r)
		}
	}
}

// SetCorsHeaders allows all methods (for testing purposes)
func SetCorsHeaders() Adapter {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			//allow all CORS
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			if (*r).Method == "OPTIONS" {
				return
			}
			f(w, r)
		}
	}
}

// ResponseWithCode returns formed JSON
func ResponseWithCode(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	resp := string(response)

	if code != 200 {
		log.Printf("responseCode: %s body: %s", strconv.Itoa(code), resp)
	} else {
		log.Printf("responseCode: %s", strconv.Itoa(code))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(resp))
}

// CreateUUID creates a Guid for error tracing
func CreateUUID() string {
	b, _ := uuid.NewRandom()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}
