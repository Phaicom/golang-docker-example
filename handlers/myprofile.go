package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phaicom/golang-docker-example/mocks"
)

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	profile := mocks.FindByUsername(username)
	RenderTemplate(w, "./templates/myprofile.html", profile)
}
