package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/phaicom/golang-docker-example/models"
)

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{
		UID:      1,
		Username: RandomString(2) + strconv.Itoa(1),
		Name:     RandomString(5) + " " + RandomString(5),
		Age:      15 + rand.Intn(20),
		DOB:      time.Now(),
	}
	RenderTemplate(w, "./templates/myprofile.html", profile)
}
