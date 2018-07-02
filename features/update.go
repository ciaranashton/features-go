package features

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CiaranAshton/features-go/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// UpdateFeature updates a given feature
func (fa FeatureAPI) UpdateFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	f := models.Feature{}

	if !bson.IsObjectIdHex(id) {
		fa.l.Err.Printf("Id %s is not a valid Id \n", id)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid Id: %s", id)
		return
	}

	oid := bson.ObjectIdHex(p.ByName("id"))
	f.Id = oid

	json.NewDecoder(r.Body).Decode(&f)

	err := fa.db.UpdateFeature(fa.l, oid, &f)

	if err != nil {
		fa.l.Err.Println("Unable to find feature:", id)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Feature not found: %s", id)
		return
	}

	fj, _ := json.Marshal(f)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", fj)
}
