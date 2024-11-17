package routes

import (
	"net/http"
	"server/http/user"
	"server/storage/userstore"
)

func UserRoutes(store *userstore.MongoStore) http.Handler{
	router := http.NewServeMux()
	userHandler := user.NewHandler(*store)
	// user Handler
	router.HandleFunc("GET /", userHandler.Login)
	router.HandleFunc("POST /", userHandler.SignIn)
	router.HandleFunc("PUT /", userHandler.UpdateUser)
	router.HandleFunc("DELETE /", userHandler.Logout)

	return router
	
}