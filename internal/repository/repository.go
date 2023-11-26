package repository

import "udemyCourse1/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InserReservation(res models.Reservation) error
}
