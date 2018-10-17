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
