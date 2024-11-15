package user

import (
	"encoding/json"
	"net/http"
	"server/components/provider/jwt"
	"server/models"
	"server/storage/userstore"
	"time"
)

type Handler struct{
	store userstore.MongoStore
}


func NewHandler(store userstore.MongoStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request){
	var user models.User

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil { 
		http.Error(w, "Invalid input: "+ err.Error(), http.StatusBadRequest)
		return 
	}
	
	if err = h.store.FindUser(user.Email, user.Password); err != nil{
		http.Error(w, "No user found" + err.Error(),  http.StatusUnauthorized)
		return
	}

	tokenString, err  := jwt.IssuesToken(user)
	if err != nil{
		http.Error(w,  "Error Issuing token:" + err.Error(), http.StatusInternalServerError)
	}
	setCookieHandler(w, tokenString)
	w.WriteHeader(http.StatusOK)
	
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request){

}
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request){
	var user models.User 
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil { 
		http.Error(w, "Invalid input: "+ err.Error(), http.StatusBadRequest)
		return 
	}
	err, newUser := CreateNewUser(&user)
	if err != nil{
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
	}
	
	if err := h.store.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&newUser)
	if err != nil {
		http.Error(w, "Failed to encode response: "+ err.Error(), http.StatusInternalServerError) 
	}
}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request){

}

func setCookieHandler(w http.ResponseWriter, tokenString string){
	cookie := http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(time.Hour * 24),
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write([]byte("Cookie has been set"))
}