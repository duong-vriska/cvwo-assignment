package utils

import (
	"encoding/json"
	"net/http"
)

func Parse(w http.ResponseWriter, r *http.Request, params interface{}) {
    err := json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
}
