package handlers

import (
	"net/http"

	"github.com/phaicom/golang-docker-example/mocks"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	profiles := mocks.OpenJson()
	RenderTemplate(w, "./templates/profile.html", profiles)
}
