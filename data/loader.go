package data

import (
	"embed"
	"encoding/json"
	"github.com/VNG-Realisatie/referentielijsten-api/models"
	"log"
	"path"
)

//go:embed json
var jsonFiles embed.FS

func LoadData() (*models.Loader, error) {
	var loader models.Loader
	filePath := "json"
	rootDir, err := jsonFiles.ReadDir(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, dir := range rootDir {
		log.Printf("loading %s\n", dir.Name())
		switch dir.Name() {
		case "resultaattypeomschrijvingen.json":
			plan, _ := jsonFiles.ReadFile(path.Join(filePath, dir.Name()))
			var results models.ResultaattypeOmschrijvingen
			err := json.Unmarshal(plan, &results)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			loader.ResultaattypenOmscrhijvingen = results

		case "procestypen.json":
			plan, _ := jsonFiles.ReadFile(path.Join(filePath, dir.Name()))
			var results models.ProcesTypen
			err := json.Unmarshal(plan, &results)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			loader.ProcesTypen = results

		case "resultaten.json":
			plan, _ := jsonFiles.ReadFile(path.Join(filePath, dir.Name()))
			var results []models.Result
			err := json.Unmarshal(plan, &results)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			loader.Resultaten = results
		}

	}

	return &loader, nil
}
