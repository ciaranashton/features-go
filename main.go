package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/CiaranAshton/features/features"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

var (
	info *log.Logger
	warn *log.Logger
	er   *log.Logger
)

func init() {
	info = log.New(os.Stdout,
		"\033[1;32m[Info]:\033[0m ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warn = log.New(os.Stdout,
		"\033[1;33m[Warning]:\033[0m ",
		log.Ldate|log.Ltime|log.Lshortfile)

	er = log.New(os.Stderr,
		"\033[1;31m[Error]:\033[0m ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Get Environment variables
	err := godotenv.Load()
	if err != nil {
		er.Fatalln("Error loading environment variables", err)
	}
	p := os.Getenv("PORT")

	// Create new instance of the Features API package
	fa := features.New(getSession(), info, warn, er)

	// Create Router
	r := httprouter.New()

	// Routes
	r.GET("/features", fa.GetFeatures)
	r.GET("/features/:id", fa.GetFeature)
	r.POST("/features", fa.CreateFeature)
	r.DELETE("/features/:id", fa.DeleteFeature)

	// Listen on port 8080
	info.Println("Server listening on port:", p)
	er.Fatal(http.ListenAndServe("localhost:"+p, r))
}

// Setup or MongoDB session. Currently, hitting a local instance of mongo.
func getSession() *mgo.Session {
	info.Println("[MongoDB] Connecting to MongoDB...")
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		er.Fatalln(err)
	}
	return s
}
