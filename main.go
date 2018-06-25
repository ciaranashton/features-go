package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/CiaranAshton/features/features"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	fa := features.New(getSession())

	r.GET("/features/:id", fa.GetFeature)
	r.POST("/features", fa.CreateFeature)
	r.DELETE("/features/:id", fa.DeleteFeature)

	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
