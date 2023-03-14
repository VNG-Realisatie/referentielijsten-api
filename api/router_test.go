package api

import (
	"encoding/json"
	"github.com/VNG-Realisatie/referentielijsten-api/data"
	"github.com/VNG-Realisatie/referentielijsten-api/middleware"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestHealthEndPoint(t *testing.T) {
	testConfig, err := data.LoadData()
	assert.Nil(t, err)
	router := InitRoutes(testConfig)

	t.Run("Healthy", func(t *testing.T) {
		response := performGetRequest(router, "/api/v1/health")
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.Health
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.True(t, sut.Healthy)
	})
}

func TestResultaattypenomschrijvingenEndPoint(t *testing.T) {
	testConfig, err := data.LoadData()
	assert.Nil(t, err)
	router := InitRoutes(testConfig)
	resulutsLen := len(testConfig.ResultaattypenOmscrhijvingen)

	t.Run("ListValidInput", func(t *testing.T) {
		response := performGetRequest(router, "/api/v1/resultaattypeomschrijvingen")
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.ResultaattypeOmschrijvingen
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, resulutsLen, len(sut))
	})

	t.Run("ListWrongMethod", func(t *testing.T) {
		response := performPostRequest(router, "/api/v1/resultaattypeomschrijvingen", nil)
		assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	})

	t.Run("GetValidInput", func(t *testing.T) {
		resultaatOmschrijving := testConfig.ResultaattypenOmscrhijvingen[0]
		response := performGetRequest(router, "/api/v1/resultaattypeomschrijvingen/"+resultaatOmschrijving.URL)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.ResultaattypeOmschrijvingenElement
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, resultaatOmschrijving.Omschrijving, sut.Omschrijving)
	})

	t.Run("GetNoneExistingId", func(t *testing.T) {
		randomUuid := middleware.CreateUUID()
		response := performGetRequest(router, "/api/v1/resultaattypeomschrijvingen/"+randomUuid)
		assert.Equal(t, http.StatusNotFound, response.Code)

		var sut models.HTTPError
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "not_found", sut.Code)
	})
}

func TestProcesTypen(t *testing.T) {
	testConfig, err := data.LoadData()
	assert.Nil(t, err)
	router := InitRoutes(testConfig)
	resulutsLen := len(testConfig.ProcesTypen)

	t.Run("ListValidInput", func(t *testing.T) {
		response := performGetRequest(router, "/api/v1/procestypen")
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.ProcesTypen
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, resulutsLen, len(sut))
	})

	t.Run("ListWrongMethod", func(t *testing.T) {
		response := performPostRequest(router, "/api/v1/resultaattypeomschrijvingen", nil)
		assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	})

	t.Run("GetValidInput", func(t *testing.T) {
		procesType := testConfig.ProcesTypen[0]
		response := performGetRequest(router, "/api/v1/procestypen/"+procesType.URL)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.ProcesTypeElement
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, procesType.Omschrijving, sut.Omschrijving)
	})

	t.Run("ListValidYear", func(t *testing.T) {
		procesType := testConfig.ProcesTypen[0]
		year := strconv.Itoa(int(procesType.Jaar))
		response := performGetRequest(router, "/api/v1/procestypen?jaar="+year)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.ProcesTypen
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		for _, result := range sut {
			assert.Equal(t, procesType.Jaar, result.Jaar)
		}

	})

	t.Run("GetNoneExistingId", func(t *testing.T) {
		randomUuid := middleware.CreateUUID()
		response := performGetRequest(router, "/api/v1/procestypen/"+randomUuid)
		assert.Equal(t, http.StatusNotFound, response.Code)

		var sut models.HTTPError
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "not_found", sut.Code)
	})

	t.Run("ListNoneExistingYear", func(t *testing.T) {
		year := "235235523"
		response := performGetRequest(router, "/api/v1/procestypen?jaar="+year)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.ProcesTypen
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, 0, len(sut))
	})

	t.Run("ListInvalidYear", func(t *testing.T) {
		year := "ICANNOTBEPARSED"
		response := performGetRequest(router, "/api/v1/procestypen?jaar="+year)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		var sut models.ValidationError
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "jaar", sut.InvalidParams[0].Name)
	})
}

func TestResultaatTypen(t *testing.T) {
	testConfig, err := data.LoadData()
	assert.Nil(t, err)
	router := InitRoutes(testConfig)
	maxPage := len(testConfig.Resultaten) / 100

	t.Run("ListValidInput", func(t *testing.T) {
		response := performGetRequest(router, "/api/v1/resultaten")
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.Resultaten
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, 99, len(sut.Results))
	})

	t.Run("ListWrongMethod", func(t *testing.T) {
		response := performPostRequest(router, "/api/v1/resultaten", nil)
		assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	})

	t.Run("GetValidInput", func(t *testing.T) {
		resultaatType := testConfig.Resultaten[0]
		response := performGetRequest(router, "/api/v1/resultaten/"+resultaatType.URL)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.Result
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, resultaatType.Omschrijving, sut.Omschrijving)
	})

	t.Run("ListFirstPage", func(t *testing.T) {
		firstPage := 1
		parsedPage := strconv.Itoa(firstPage)
		nextPage := strconv.Itoa(firstPage + 1)
		response := performGetRequest(router, "/api/v1/resultaten?page="+parsedPage)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.Resultaten
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.True(t, strings.Contains(sut.Next, nextPage))
		assert.Nil(t, sut.Previous)
	})

	t.Run("ListLastPage", func(t *testing.T) {
		parsedPage := strconv.Itoa(maxPage + 1)
		response := performGetRequest(router, "/api/v1/resultaten?page="+parsedPage)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.Resultaten
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "", sut.Next)
		assert.NotNil(t, sut.Previous)
	})

	t.Run("ListPageOutOfBound", func(t *testing.T) {
		nonExistingPage := maxPage + 2
		parsedPage := strconv.Itoa(nonExistingPage)
		response := performGetRequest(router, "/api/v1/resultaten?page="+parsedPage)
		assert.Equal(t, http.StatusNotFound, response.Code)

		var sut models.HTTPError
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "not_found", sut.Code)
	})

	t.Run("ListInvalidPage", func(t *testing.T) {
		invalidPage := "swkjendhfs"
		response := performGetRequest(router, "/api/v1/resultaten?page="+invalidPage)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		var sut models.ValidationError
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "page", sut.InvalidParams[0].Name)
	})

	t.Run("ListValidProcesType", func(t *testing.T) {
		resultaatType := testConfig.Resultaten[0]
		response := performGetRequest(router, "/api/v1/resultaten?proces_type=http://localhost:8000/api/v1/"+resultaatType.ProcesType)
		assert.Equal(t, http.StatusOK, response.Code)

		var sut models.Resultaten
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)

		for _, result := range sut.Results {
			assert.Contains(t, result.ProcesType, resultaatType.ProcesType)
		}
	})

	t.Run("ListInvalidProcesType", func(t *testing.T) {
		invalidUrl := "skljdhnfskjdnfsdf"
		response := performGetRequest(router, "/api/v1/resultaten?proces_type="+invalidUrl)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		var sut models.ValidationError
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "proces_type", sut.InvalidParams[0].Name)
	})

	t.Run("GetNoneExistingId", func(t *testing.T) {
		randomUuid := middleware.CreateUUID()
		response := performGetRequest(router, "/api/v1/resultaten/"+randomUuid)
		assert.Equal(t, http.StatusNotFound, response.Code)

		var sut models.HTTPError
		err = json.NewDecoder(response.Body).Decode(&sut)

		assert.Nil(t, err)
		assert.Equal(t, "not_found", sut.Code)
	})

}

func performGetRequest(r http.Handler, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performPostRequest(r http.Handler, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
