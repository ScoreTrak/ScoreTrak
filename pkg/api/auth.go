package api

import (
	"ScoreTrak/pkg/config"
	"encoding/json"
	"net/http"
	"strings"
)

//TokenVerify verifies authentication token and passes on the response if token found successfully
func TokenVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header != config.Token() {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing or incorrect auth token")
			return
		}
		next.ServeHTTP(w, r)
	})
}
