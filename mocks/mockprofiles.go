package mocks

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/phaicom/golang-docker-example/models"
)

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}

func CreateJson() {
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

	data, err := json.Marshal(profiles)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("./profiles.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Profiles data written to ", file.Name())
}

func OpenJson() []models.Profile {
	profiles := make([]models.Profile, 0)

	file, err := os.Open("./profiles.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&profiles)
	if err != nil {
		log.Fatal(err)
	}
	return profiles
}

func FindByUsername(username string) models.Profile {
	profile := models.Profile{}
	profiles := OpenJson()
	for _, p := range profiles {
		if p.Username == username {
			profile = p
			break
		}
	}
	return profile
}
