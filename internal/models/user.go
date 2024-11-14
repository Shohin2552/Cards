package models

import "time"

type User struct {
	Id          int       `json:"id"`
	Name        string    ` json:"name"`
	Number      int       `json:"number"`
	Create_Date time.Time `json:"create_date"`
}

type UserRole struct {
	User User `json:"user"`
}
