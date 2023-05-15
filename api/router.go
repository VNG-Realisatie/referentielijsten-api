package api

import (
	"github.com/VNG-Realisatie/referentielijsten-api/api/v1"
	"github.com/VNG-Realisatie/referentielijsten-api/middleware"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"github.com/gorilla/mux"
)

// InitRoutes to start up a mux router and return the routes
func InitRoutes(loader *models.Loader) *mux.Router {
	serveMux := mux.NewRouter()

	v1Handler := v1.ReferentielijstenHandler{
		Data: loader,
	}

	serveMux.HandleFunc("/api/v1/health", middleware.Adapt(v1Handler.Health, middleware.ValidateRestMethod("GET"), middleware.SetCorsHeaders()))

	//resultaten
	serveMux.HandleFunc("/api/v1/resultaten", middleware.Adapt(v1Handler.ListResultaten, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))
	serveMux.HandleFunc("/api/v1/resultaten/{uuid}", middleware.Adapt(v1Handler.GetResultaat, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))

	//resultaattypeomschrijvingen
	serveMux.HandleFunc("/api/v1/resultaattypeomschrijvingen", middleware.Adapt(v1Handler.ListResultaattypeomschrijvingen, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))
	serveMux.HandleFunc("/api/v1/resultaattypeomschrijvingen/{uuid}", middleware.Adapt(v1Handler.GetResultaattypeomschrijving, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))

	//procestypen
	serveMux.HandleFunc("/api/v1/procestypen", middleware.Adapt(v1Handler.ListProcesTypen, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))
	serveMux.HandleFunc("/api/v1/procestypen/{uuid}", middleware.Adapt(v1Handler.GetProcesType, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))

	//redoc
	serveMux.HandleFunc("/api/v1/schema", middleware.Adapt(serveRedoc, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))
	serveMux.HandleFunc("/api/v1/openapi.yaml", middleware.Adapt(serveOpenApi, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))
	serveMux.HandleFunc("/api/v1/schema/openapi.yaml", middleware.Adapt(serveOpenApi, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails(), middleware.SetCorsHeaders()))

	return serveMux
}
