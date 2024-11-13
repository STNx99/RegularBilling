package routes

import (
	"net/http"
	"server/user"
)

func UserRoutes() http.Handler{
	router := http.NewServeMux()
	handler := &user.Handler{}
	// user Handler
	router.HandleFunc("GET /", handler.Login)
	router.HandleFunc("POST /", handler.SignIn)
	router.HandleFunc("PUT /", handler.UpdateUser)
	router.HandleFunc("DELETE /", handler.Logout)
	return router
	
}