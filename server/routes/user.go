package routes

import (
	"net/http"
	"server/http/user"
	"server/storage/userstore"
)

func UserRoutes(store *userstore.MongoStore) http.Handler {
	router := http.NewServeMux()
	userHandler := user.NewHandler(*store)
	// user Handler
	router.HandleFunc("/user/login", userHandler.Login)
	router.HandleFunc("/user/signin", userHandler.SignIn)
	router.HandleFunc("/user/update", userHandler.UpdateUser)
	router.HandleFunc("/user/logout", userHandler.Logout)

	return router

}
