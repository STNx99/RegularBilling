package routes

import (
	"net/http"
	"server/http/service"
	"server/storage/servicestore"
	"server/storage/userstore"
)

func ServiceRoutes(serviceStore *servicestore.MongoStore, userStore *userstore.MongoStore) http.Handler {
	router := http.NewServeMux()
	serviceHandler := service.NewHandler(*serviceStore, *userStore)
	// service handler
	router.HandleFunc("GET /", serviceHandler.FindUserService)
	router.HandleFunc("POST /", serviceHandler.Add)
	router.HandleFunc("DELETE /", serviceHandler.Delete)

	return router

}
