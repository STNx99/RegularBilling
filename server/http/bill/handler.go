package bill

import (
	"encoding/json"
	"net/http"
	"server/middleware"
	"server/models"
	"server/storage/userstore"
)

type Handler struct {
	store *userstore.MongoStore
}

func NewHandle(store *userstore.MongoStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	username, err := middleware.GetUsernameFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}
	user := models.User{UserName: username}
	bills, err := h.store.FindUserBill(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	yearTotal, err := YearTotal(bills)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result := CreateNewBill(&bills, yearTotal)

	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
