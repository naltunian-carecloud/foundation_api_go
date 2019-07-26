package bundle

import (
	"log"
	"encoding/json"
	"github.com/CareCloud/gofhir/models"
	"github.com/gorilla/mux"
	"net/http"
)

//Routes - Account
func Routes(v1 *mux.Router) {
	v1.HandleFunc("/bundle", initialize).Methods("POST")
}

type BundleMappingData struct {
	Bundle models.Bundle `json:"bundle"`
}

func initialize(w http.ResponseWriter, r *http.Request) {
	var (
		bundleMappingData BundleMappingData
		err error
	)
	err = json.NewDecoder(r.Body).Decode(&bundleMappingData)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(&bundleMappingData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}