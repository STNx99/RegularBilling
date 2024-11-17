package routes

import (
	"net/http"
	"server/http/service"
	"server/storage/servicestore"
	"server/storage/userstore"
)

func ServiceRoutes(serviceStore *servicestore.MongoStore, userStore *userstore.MongoStore) http.Handler{
	router := http.NewServeMux()
	serviceHandler := service.NewHandler(*serviceStore, *userStore)
	// service handler
	router.HandleFunc("GET /", serviceHandler.FindAll)
	router.HandleFunc("POST /", serviceHandler.Add)

	return router
	
}