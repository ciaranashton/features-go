package features

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/CiaranAshton/features/models"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

// FeatureAPI structure
type FeatureAPI struct {
	session *mgo.Session
}

// New method for creating a new instance of a FeatureAPI
func New(s *mgo.Session) *FeatureAPI {
	return &FeatureAPI{s}
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

// CreateFeature creates a new feature and stores it in the database
func (fa FeatureAPI) CreateFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f := models.Feature{}

	json.NewDecoder(r.Body).Decode(&f)

	f.Id = bson.NewObjectId()

	fa.session.DB("cjla").C("features").Insert(f)

	fj, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", fj)
}

// DeleteFeature removes function from database
func (fa FeatureAPI) DeleteFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := fa.session.DB("cjla").C("features").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted Feature", oid, "\n")
}
