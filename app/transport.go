package app

import "net/http"

func AppRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/dude", func(w http.ResponseWriter, r *http.Request) {})

	return router
}
