package main

import (
	"net/http"

	"github.com/urfave/negroni"

	"gopkg.in/mgo.v2"

	"github.com/CiaranAshton/features/features"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Create new instance of the Features API package
	fa := features.New(getSession())

	// Create Router
	r := httprouter.New()

	// Routes
	r.GET("/features/:id", fa.GetFeature)
	r.POST("/features", fa.CreateFeature)
	r.DELETE("/features/:id", fa.DeleteFeature)

	// Add logger middleware that logs each incoming request and response.
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	// Listen on port 8080
	http.ListenAndServe("localhost:8080", n)
}

// Setup or MongoDB session. Currently, hitting a local instance of mongo.
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
