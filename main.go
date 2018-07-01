package main

import (
	"net/http"
	"os"

	"github.com/CiaranAshton/features/logger"

	"github.com/CiaranAshton/features/features"
	"github.com/joho/godotenv"
)

func main() {
	// Import logger
	l := logger.NewLogger()

	// Get Environment variables
	err := godotenv.Load()
	if err != nil {
		l.Err.Fatalln("Error loading environment variables", err)
	}

	// Create new instance of the Features API package
	db := features.NewDatabase()
	api := features.New(db, l).API()

	// Listen and serve API
	p := os.Getenv("PORT")
	l.Info.Println("Server listening on port:", p)
	l.Err.Fatal(http.ListenAndServe("localhost:"+p, api))
}
