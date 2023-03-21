package v1

import (
	"github.com/VNG-Realisatie/referentielijsten-api/middleware"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"net/http"
	"time"
)

func (r *ReferentielijstenHandler) Health(w http.ResponseWriter, req *http.Request) {
	// swagger:route GET /health health health_retrieve
	//
	// Controleer of de api online is.
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
	//	  200: Health
	currentTime := time.Now()

	health := models.Health{
		Healthy: true,
		Time:    currentTime.String(),
	}

	middleware.ResponseWithCode(w, http.StatusOK, health)
}
