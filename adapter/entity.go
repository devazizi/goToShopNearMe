package adapter

import "time"

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Address struct {
	ID        uint
	Address   string
	Latitude  float32
	Longitude float32
	CreatedAt time.Time
	UpdatedAt time.Time
}
