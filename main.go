package main

import (
	"log"
	"net/http"
	"os"

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
		log.Ldate|log.Ltime)

	debug = log.New(os.Stdout,
		"\033[1;34m[Debug]:\033[0m ",
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
	db := features.NewDatabase()
	api := features.New(db, info, debug, er).API()

	// Listen and serve API
	p := os.Getenv("PORT")
	info.Println("Server listening on port:", p)
	er.Fatal(http.ListenAndServe("localhost:"+p, api))
}
