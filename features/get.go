package features

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CiaranAshton/features/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// GetFeatures returns all features in the db
func (fa FeatureAPI) GetFeatures(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fa.info.Println("[MongoDB] Fetching all features from mongoDB")
	fs, err := fa.db.GetAllFeatures(fa, w)

	if err != nil {
		fa.err.Println(err)
	}

	fsj, err := json.Marshal(fs)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", fsj)
}

// GetFeature returns status 200 and JSON of the desired feature
func (fa FeatureAPI) GetFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid := bson.ObjectIdHex(id)

	f := models.Feature{}

	err := fa.db.GetFeature(fa, oid, &f)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	fj, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", fj)
}
