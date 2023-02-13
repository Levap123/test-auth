package transport

import (
	"encoding/json"
	"net/http"
)

func sendJSON(w http.ResponseWriter, status int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}

func readJSON(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(dest)
}
