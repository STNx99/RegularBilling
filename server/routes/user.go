package routes

import (
	"net/http"
	"server/http/user"
	"server/middleware"
	"server/storage/userstore"
)

func UserRoutes(store *userstore.MongoStore) http.Handler {
	router := http.NewServeMux()
	userHandler := user.NewHandler(*store)

	//User Handler
	router.Handle("/user/find", middleware.VerifyJWT(http.HandlerFunc(userHandler.Find)))
	router.HandleFunc("/user/login", userHandler.Login)
	router.HandleFunc("/user/signin", userHandler.SignIn)
	router.Handle("/user/update", middleware.VerifyJWT(http.HandlerFunc(userHandler.UpdateUser)))
	router.Handle("/user/logout", middleware.VerifyJWT(http.HandlerFunc(userHandler.Logout)))

	return router

}
