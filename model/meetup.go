package model

type Meetups struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}
