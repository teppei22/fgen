package main

import (
	"log"
	"os"
	"time"

	"github.com/teppei22/fji-codegen/sample/router"
)

const location = "Asia/Tokyo"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	// ctx := context.Background()

	// サービス（DB）挿入

	// コントローラ（Controller）挿入

	// ルーター定義
	router := router.Init()

	router.Logger.Fatal(router.Start(":" + port))

}
