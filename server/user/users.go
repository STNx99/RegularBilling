package user

import "time"

type User struct {
	userId    string 
	userName  string
	email     string
	password  string
	credits  float64
	createdAt time.Time
	updatedAt time.Time
}