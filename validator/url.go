package validator

import (
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"log"
	"net/url"
)

// Url checks if an url can be parsed and returns a validationerror if not
func Url(rawurl, name string) *models.FieldValidationError {
	parsed, err := url.ParseRequestURI(rawurl)
	if parsed == nil {
		log.Printf("%s is a NOT valid url: %v", name, err.Error())
		return &models.FieldValidationError{
			Name:   name,
			Code:   "invalid",
			Reason: "Opgegeven URL is niet valide."}

	}

	if parsed.Host == "" || parsed.Scheme == "" {
		log.Printf("%s is a NOT valid url missing a scheme or host: %s", name, rawurl)
		return &models.FieldValidationError{
			Name:   name,
			Code:   "invalid",
			Reason: "Opgegeven URL is niet valide."}
	}

	if !(parsed.Scheme == "http" || parsed.Scheme == "https") {
		log.Printf("%s is a NOT valid url scheme not http or https (not supported): %s", name, rawurl)
		return &models.FieldValidationError{
			Name:   name,
			Code:   "invalid",
			Reason: "Opgegeven URL is niet valide."}
	}

	return nil
}
