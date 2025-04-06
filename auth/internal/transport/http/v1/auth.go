package v1

import "net/http"

type AuthService interface{
	Register(name string, password string) (string, error)
	Login(name string, password string) (string, error)
}

type AuthHandler struct{
	as AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler{
	return &AuthHandler{as: service}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}
	
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}
	
}