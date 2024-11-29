package user

import (
	"encoding/json"
	"net/http"
	"server/middleware"
	"server/models"
	"server/storage/userstore"
	"time"
)

type Handler struct {
	store userstore.MongoStore
}

func NewHandler(store userstore.MongoStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := h.store.FindUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, "No user found"+err.Error(), http.StatusUnauthorized)
		return
	}

	tokenString, err := middleware.IssuesToken(*foundUser)
	if err != nil {
		http.Error(w, "Error Issuing token:"+err.Error(), http.StatusInternalServerError)
	}
	setCookieHandler(w, tokenString, time.Now().Add(time.Hour*24))
	// cookie := map[string]interface{}{
	// 	"success": true,
	// 	"token":   tokenString,
	// }

	// err = json.NewEncoder(w).Encode(&cookie)
	// if err != nil {
	// 	http.Error(w, "Err encoding: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }
	w.WriteHeader(http.StatusOK)

}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	setCookieHandler(w, "", time.Now().Add(-time.Hour))
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	err, newUser := CreateNewUser(&user)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
	}

	if err := h.store.CreateUser(&newUser); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	tokenString, err := middleware.IssuesToken(newUser)
	if err != nil {
		http.Error(w, "Error Issuing token:"+err.Error(), http.StatusInternalServerError)
	}
	setCookieHandler(w, tokenString, time.Now().Add(time.Hour*24))
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	username, err := middleware.GetUsernameFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	currUser, err := h.store.LoggedInUser(username)
	if currUser == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(&currUser)
	if err != nil {
		http.Error(w, "Err encoding: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func setCookieHandler(w http.ResponseWriter, tokenString string, expTime time.Time) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expTime,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write([]byte("Cookie has been set"))
}
