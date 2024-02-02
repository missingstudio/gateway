package utils

import "net/http"

func Custom404handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "custom 404 error", http.StatusNotFound)
	})
}
