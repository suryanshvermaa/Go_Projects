package main

import (
	"log"
	"net/http"

	"github.com/Go_Projects/golang-MySQL/pkg/routes"
	"github.com/Go_Projects/golang-MySQL/pkg/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	http.HandleFunc("Get /", func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, 200, "server is healty✅", nil)
	})
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
