package features

import (
	"github.com/CiaranAshton/features/logger"
	"github.com/julienschmidt/httprouter"

	"github.com/urfave/negroni"
)

// FeatureAPI structure
type FeatureAPI struct {
	db DB
	l  *logger.Logger
}

// New function for creating an instance of FeatureAPI
func New(db DB, l *logger.Logger) *FeatureAPI {
	return &FeatureAPI{db, l}
}

// API defines the api routes for the service
func (fa FeatureAPI) API() *negroni.Negroni {
	// Create Router
	mux := httprouter.New()

	// 	Middlewares
	n := negroni.New()
	n.UseHandler(logger.ResponseLogger(mux))

	// Routes
	mux.GET("/features", fa.GetFeatures)
	mux.GET("/features/:id", fa.GetFeature)
	mux.POST("/features", fa.CreateFeature)
	mux.DELETE("/features/:id", fa.DeleteFeature)

	return n
}
