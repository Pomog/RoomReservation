package repository

import (
	"time"
	"udemyCourse1/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	AllRooms() ([]models.Room, error)

	InserReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomById(id int) (models.Room, error)

	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Autenticate(email, testPassword string) (int, string, error)

	AllReservations() ([]models.Reservation, error)
	AllNewReservations() ([]models.Reservation, error)
	GetReservationByID(id int) (models.Reservation, error)
	UpdateReservation(r models.Reservation) error
	DeleteReservation(id int) error
	UpdateProcessedForReservation(id, processed int) error
	GetRestictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)
	DeleteBlockForRoom(id int) error
	InsertBlockForRoom(id int, startDate time.Time) error
}
