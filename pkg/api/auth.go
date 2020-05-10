package api

import (
	"ScoreTrak/pkg/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//TokenVerify verifies authentication token and passes on the response if token found successfully
func TokenVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header != config.GetToken() {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing or incorrect auth token")
			return
		}

		json.NewEncoder(w).Encode(fmt.Sprintf("Token found. Value %s", header))
		next.ServeHTTP(w, r)
	})
}
