package v1

import "github.com/VNG-Realisatie/referentielijsten-api/models"

const (
	UUIDFromParam       string = "uuid"
	JaarFromParam       string = "jaar"
	ProcesTypeFromParam string = "proces_type"
	PageFromParam       string = "page"
)

type ReferentielijstenHandler struct {
	Data *models.Loader
}
