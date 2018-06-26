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
	fs := []models.Feature{}

	fa.info.Println("[MongoDB] Fetching all features from mongoDB")
	if err := fa.session.DB("cjla").C("features").Find(nil).All(&fs); err != nil {
		w.WriteHeader(404)
		return
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

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	f := models.Feature{}

	if err := fa.session.DB("cjla").C("features").FindId(oid).One(&f); err != nil {
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
