package features

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CiaranAshton/features/models"
	"github.com/julienschmidt/httprouter"
)

// GetFeatures returns all features in the db
func (fa FeatureAPI) GetFeatures(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fs := []models.Feature{}

	err := fa.db.GetAllFeatures(fa.debug, &fs)

	if err != nil {
		fa.err.Println("Unable to find features")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Features not found\n")
		return
	}

	fsj, _ := json.Marshal(fs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", fsj)
}

// GetFeature returns status 200 and JSON of the desired feature
func (fa FeatureAPI) GetFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	f := models.Feature{}

	err := fa.db.GetFeature(fa.debug, id, &f)

	if err != nil {
		fa.err.Println("Unable to find feature:", id)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Feature not found: %s\n", id)
		return
	}

	fj, _ := json.Marshal(f)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", fj)
}
