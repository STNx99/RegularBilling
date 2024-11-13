package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}


func (h *Handler) Login(w http.ResponseWriter, r *http.Request){

}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request){

}
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request){
	var newUser User 
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) 
	}


}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request){

}