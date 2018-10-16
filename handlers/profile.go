package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/phaicom/golang-docker-example/models"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	profiles := make([]models.Profile, 0)
	for i := 0; i < 5; i++ {
		profile := models.Profile{
			UID:      i,
			Username: RandomString(2) + strconv.Itoa(i),
			Name:     RandomString(5) + " " + RandomString(5),
			Age:      15 + rand.Intn(20),
			DOB:      time.Now(),
		}
		profiles = append(profiles, profile)
	}
	RenderTemplate(w, "./templates/profile.html", profiles)
}
