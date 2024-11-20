package bill

import (
	"net/http"
	"server/models"
	"server/storage/userstore"
)

type Handler struct {
	store *userstore.MongoStore
}

func(h *Handler) Find(w http.ResponseWriter, r *http.Request){
	var user models.User

	bills, err := h.store.FindUserBill(&user)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}