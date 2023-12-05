package dbrepo

import (
	"time"
	"udemyCourse1/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InserReservation inserts a reservation into the database
func (m *testDBRepo) InserReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomById gets a room by ID
func (m *testDBRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	return room, nil
}