package routes

import (
	"net/http"
	"server/http/bill"
	"server/storage/userstore"
)

func BillRoutes(store *userstore.MongoStore) http.Handler{
	router := http.NewServeMux()
	billHandler := bill.NewHandle(store)
	// user Handler
	router.HandleFunc("GET /",billHandler.Find)

	return router
	
}