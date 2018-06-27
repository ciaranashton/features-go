package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/CiaranAshton/features/features"
	"github.com/joho/godotenv"
)

var (
	info  *log.Logger
	debug *log.Logger
	er    *log.Logger
)

func init() {
	info = log.New(os.Stdout,
		"\033[1;32m[Info]:\033[0m ",
		log.Ldate|log.Ltime|log.Lshortfile)

	debug = log.New(os.Stdout,
		"\033[1;33m[Debug]:\033[0m ",
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

	// Create new instance of the Features API package
	api := features.New(getSession(), info, debug, er).API()

	// Listen and serve API
	p := os.Getenv("PORT")
	info.Println("Server listening on port:", p)
	er.Fatal(http.ListenAndServe("localhost:"+p, api))
}

// Setup or MongoDB session. Currently, hitting a local instance of mongo.
func getSession() *mgo.Session {
	info.Println("[MongoDB] Connecting to MongoDB...")
	s, err := mgo.Dial(os.Getenv("MONGO_CONNECT"))

	if err != nil {
		er.Fatalln(err)
	}
	return s
}
