package features

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

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