package basic

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CareCloud/gofhir/models"
	"github.com/gorilla/mux"
)

//Routes - Account
func Routes(v1 *mux.Router) {
	v1.HandleFunc("/basic/initialize", initializeBasic).Methods("POST")
}

type BasicMappingData struct {
	Basic models.Basic `json:"basic"`
}

func initializeBasic(w http.ResponseWriter, r *http.Request) {
	var (
		basicMappingData BasicMappingData
		err              error
	)

	err = json.NewDecoder(r.Body).Decode(&basicMappingData)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basic := basicMappingData.Basic.New()

	err = json.NewEncoder(w).Encode(&basic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
