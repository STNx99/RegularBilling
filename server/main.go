package main

import (
	"fmt"
	"net/http"
	"server/middleware"
	"server/routes"
)

func main(){
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", routes.UserRoutes()))

	server := http.Server{
		Addr: ":8080",
		Handler: middleware.Logging(v1),
	}
	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}