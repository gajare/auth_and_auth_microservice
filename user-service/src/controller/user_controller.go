package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"user-service/src/model"
	"user-service/src/service"

	"github.com/gorilla/mux"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var creds model.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = service.SignupUser(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds model.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	token, err := service.LoginUser(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Could not login user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := service.ParseToken(tokenStr)
	if err != nil {
		http.Error(w, "Could not parse token", http.StatusUnauthorized)
	}
	w.Write([]byte("Welcome " + claims.Username + "!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	err = service.UpdateUser(id, user.Username, user.Password)
	if err != nil {
		http.Error(w, "could not update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := service.DeleteUser(id)
	if err != nil {
		http.Error(w, "could not delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
