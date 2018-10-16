package models

import (
	"time"
)

type Profile struct {
	UID      int       `json:"uid"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	DOB      time.Time `json:"dob"`
}

var ProfilesMock []Profile

// for mocking data to DB
// func MockData() {
// 	ProfilesMock := make([]Profile, 0)
// 	for i := 0; i < 5; i++ {
// 		profile := Profile{
// 			UID:      i,
// 			Username: RandomString(2) + strconv.Itoa(i),
// 			Name:     RandomString(5) + " " + RandomString(5),
// 			Age:      15 + rand.Intn(20),
// 			DOB:      time.Now(),
// 		}
// 		ProfilesMock = append(ProfilesMock, profile)
// 	}
// }

// func RandomString(len int) string {
// 	bytes := make([]byte, len)
// 	for i := 0; i < len; i++ {
// 		bytes[i] = byte(65 + rand.Intn(25))
// 	}
// 	return string(bytes)
// }
