package main

import (
	"log"
	"os"

	"sample_layered/router"

	"github.com/joho/godotenv"
)

// const location = "Asia/Tokyo"

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)

	// loc, err := time.LoadLocation(location)
	// if err != nil {
	// 	loc = time.FixedZone(location, 9*60*60)
	// }
	// time.Local = loc

	// router
	router := router.Init()

	router.Logger.Fatal(router.Start(":" + port))

}
