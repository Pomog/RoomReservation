package models

import "time"

// User is the model for the user table in the database
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room is the model for the rooms table in the database
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction is the model for the restrictions table in the database
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation is the model for the reservations table in the database
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

// RoomRestriction is the model for the room_restrictions table in the database
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}
