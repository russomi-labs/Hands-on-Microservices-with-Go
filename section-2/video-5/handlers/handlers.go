package handlers

import (
	"fmt"
	"net/http"
)

// MyHandler returns StatusOk and then prints success
func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprint(w, "Request was successful")
}
