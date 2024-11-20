package service

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/storage/servicestore"
	"server/storage/userstore"
)

type Handler struct {
	serviceStore servicestore.MongoStore
	userStore userstore.MongoStore
}

func NewHandler(serviceStore servicestore.MongoStore, userStore userstore.MongoStore) *Handler {
	return &Handler{
		serviceStore: serviceStore,
		userStore: userStore,
	}
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	services, err := h.serviceStore.FindAll()
	if err != nil {
		http.Error(w, "Error finding services"+err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(&services)
	if err != nil {
		http.Error(w, "Error encoding services" + err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var service models.AddUserService
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, "Invalid input:"+err.Error(), http.StatusBadRequest)
	}
	newService, err := h.serviceStore.FindService(service)
	if newService == (models.Service{}) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = h.userStore.UpdateUserServices(service.Username, newService)
	if err != nil{
		http.Error(w, "Error adding service" + err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request){
	var s  models.DeleteUserService
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil{
		http.Error(w, "Invalid input:"+err.Error(), http.StatusBadRequest)
	}
	err = h.userStore.DeleteUSerServices(s.Username, s.Service)
	if err != nil{
		http.Error(w, "Error deleting service:"+err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) FindUserService(w http.ResponseWriter, r *http.Request) {
	var user models.User
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	services, err := h.userStore.FindUserServices(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if services == nil {
		services = []models.Service{}
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(services)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
