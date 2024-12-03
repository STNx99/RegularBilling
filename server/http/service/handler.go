package service

import (
	"encoding/json"
	"net/http"
	"server/middleware"
	"server/models"
	"server/storage/servicestore"
	"server/storage/userstore"
)

type Handler struct {
	serviceStore servicestore.MongoStore
	userStore    userstore.MongoStore
}

func NewHandler(serviceStore servicestore.MongoStore, userStore userstore.MongoStore) *Handler {
	return &Handler{
		serviceStore: serviceStore,
		userStore:    userStore,
	}
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	services, err := h.serviceStore.FindAll()
	if err != nil {
		http.Error(w, "Error finding services"+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&services)
	if err != nil {
		http.Error(w, "Error encoding services"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var service models.AddUserService
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, "Invalid input:"+err.Error(), http.StatusBadRequest)
		return
	}
	err = h.userStore.AddUserServices(service.UserId, service.Service)
	if err != nil {
		http.Error(w, "Error adding service"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var s models.DeleteUserService
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid input:"+err.Error(), http.StatusBadRequest)
		return
	}
	err = h.userStore.DeleteUSerServices(s.UserId, s.Service)
	if err != nil {
		http.Error(w, "Error deleting service:"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) FindUserService(w http.ResponseWriter, r *http.Request) {
	username, err := middleware.GetUsernameFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}
	user := models.User{UserName: username}

	services, err := h.userStore.FindUserServices(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if services == nil {
		services = []models.Service{}
	}

	total := CalculateServiceTotal(services)
	servicesData := models.ServicesData{
		Services:     services,
		ServiceTotal: total,
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&servicesData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var user models.AddUserService
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	newService := CreateNewService(&user.Service)

	err = h.userStore.UpdateUserServices(user.UserId, *newService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
