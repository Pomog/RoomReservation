package repository

import (
	"time"
	"udemyCourse1/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InserReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomById(id int) (models.Room, error)

	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) (error)
	Autenticate(email, testPassword string) (int, string, error)
	
	AllReservations() ([]models.Reservation, error)
	AllNewReservations() ([]models.Reservation, error)
}
