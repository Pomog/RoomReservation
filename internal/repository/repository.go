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
}
