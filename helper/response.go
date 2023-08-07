package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-tyoe", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
