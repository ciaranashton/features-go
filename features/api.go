package features

import (
	"log"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/urfave/negroni"

	"github.com/julienschmidt/httprouter"
)

// FeatureAPI structure
type FeatureAPI struct {
	db    DB
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger
}

// New function for creating an instance of FeatureAPI
func New(db DB, i, w, e *log.Logger) *FeatureAPI {
	return &FeatureAPI{db, i, w, e}
}

// API defines the api routes for the service
func (fa FeatureAPI) API() *negroni.Negroni {
	// Create Router
	mux := httprouter.New()

	//middleware
	n := negroni.New()

	// Routes
	mux.GET("/features", fa.GetFeatures)
	mux.GET("/features/:id", fa.GetFeature)
	mux.POST("/features", fa.CreateFeature)
	mux.DELETE("/features/:id", fa.DeleteFeature)

	rl := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(mux, w, r)
		fa.info.Printf("%s %s | %d", r.Method, r.URL, m.Code)
	})

	n.UseHandler(rl)

	return n
}
