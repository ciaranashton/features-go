package features

import (
	"github.com/CiaranAshton/features-go/logger"
	"github.com/julienschmidt/httprouter"

	"github.com/unrolled/secure"
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

	sec := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.cjla.com"},
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},
		SSLRedirect:             true,
		SSLTemporaryRedirect:    false,
		SSLHost:                 "ssl.cjla.com",
		SSLHostFunc:             nil,
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:              315360000,
		STSIncludeSubdomains:    true,
		STSPreload:              true,
		ForceSTSHeader:          false,
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		ContentSecurityPolicy:   "default-src 'self'",
		ReferrerPolicy:          "same-origin",

		IsDevelopment: true,
	})

	// 	Middlewares
	n := negroni.New()
	n.Use(negroni.HandlerFunc(sec.HandlerFuncWithNext))
	n.UseHandler(logger.ResponseLogger(mux))

	// Routes
	mux.GET("/features", fa.GetFeatures)
	mux.GET("/features/:id", fa.GetFeature)
	mux.POST("/features", fa.CreateFeature)
	mux.PUT("/features/:id", fa.UpdateFeature)
	mux.DELETE("/features/:id", fa.DeleteFeature)

	return n
}
