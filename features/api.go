package features

import (
	"log"

	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

// FeatureAPI structure
type FeatureAPI struct {
	session *mgo.Session
	info    *log.Logger
	warn    *log.Logger
	err     *log.Logger
}

// New function for creating an instance of FeatureAPI
func New(s *mgo.Session, i, w, e *log.Logger) *FeatureAPI {
	return &FeatureAPI{s, i, w, e}
}

// API defines the api routes for the service
func (fa FeatureAPI) API() *httprouter.Router {
	// Create Router
	mux := httprouter.New()

	// Routes
	mux.GET("/features", fa.GetFeatures)
	mux.GET("/features/:id", fa.GetFeature)
	mux.POST("/features", fa.CreateFeature)
	mux.DELETE("/features/:id", fa.DeleteFeature)

	return mux
}
