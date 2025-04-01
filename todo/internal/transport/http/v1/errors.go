package v1

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error, code int){
	response := map[string]string{"error": err.Error()}
	info, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(info)
}