package v1

import (
	"auth/pkg/logger"
	"context"
	"encoding/json"
	"net/http"
)

type AuthService interface{
	Register(ctx context.Context, name string, password string) (string, error)
	Login(ctx context.Context, name string, password string) (string, error)
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
		return
	}

	type Data struct{
		Name string `json:"name"`
		Password string `json:"password"`
	}

	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
		WriteError(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	token, err := ah.as.Register(r.Context(), data.Name, data.Password)
	if err != nil{
		WriteError(w, err, http.StatusBadRequest)
		return
	}
	
	cookie := &http.Cookie{
		Name: "jwt_id",
		Value: token,
		Path:     "/", 
        Domain:   "", 
        MaxAge:   86400,
        HttpOnly: true,
        SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusAccepted)
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		WriteError(w, nil, http.StatusMethodNotAllowed)
		return
	}

	logger.GetFromCtx(r.Context()).InfoContext(r.Context(), "called")

	type Data struct{
		Name string `json:"name"`
		Password string `json:"password"`
	}

	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
		WriteError(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	logger.GetFromCtx(r.Context()).InfoContext(r.Context(), "called")
	token, err := ah.as.Login(r.Context(), data.Name, data.Password)
	if err != nil{
		WriteError(w, err, http.StatusBadRequest)
		return
	}
	
	cookie := &http.Cookie{
		Name: "jwt_id",
		Value: token,
		Path:     "/", 
        Domain:   "", 
        MaxAge:   86400,
        HttpOnly: true,
        SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusAccepted)
}

func (ah *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
        Name:   "jwt_id",
        Value:  "",
        MaxAge: -1,
    }
    http.SetCookie(w, cookie)
    w.WriteHeader(http.StatusAccepted)
}