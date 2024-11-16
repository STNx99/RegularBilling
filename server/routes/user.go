package routes

import (
	"net/http"
	"server/http/user"
	"server/storage/userstore"
)

func UserRoutes(store *userstore.MongoStore) http.Handler{
	router := http.NewServeMux()
	handler := user.NewHandler(*store)
	// user Handler
	router.HandleFunc("GET /user", handler.Login)
	router.HandleFunc("POST /user", handler.SignIn)
	router.HandleFunc("PUT /user", handler.UpdateUser)
	router.HandleFunc("DELETE /user", handler.Logout)
	return router
	
}