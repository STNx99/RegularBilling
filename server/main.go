package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"server/components/appContext"
	"server/middleware"
	"server/routes"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	appCtx := appContext.InitAppContext(uri)
	defer appCtx.CloseDB()

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", routes.UserRoutes()))

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(v1),
	}
	fmt.Println("Server is running on port 8080")

	_ = server.ListenAndServe()
}
