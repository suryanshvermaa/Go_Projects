package handlers

import (
	"net/http"

	"github.com/suryanshvermaa/Go_Projects/chi/utils"
)

func HanderReadiness(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJson(w, 200, "healty server", map[string]string{"server": "healty"})
}
