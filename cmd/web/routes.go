package main

import (
	"net/http"
	"udemyCourse1/internal/config"
	"udemyCourse1/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/test", handlers.Repo.TestResults)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	//not implemented yet
	mux.Get("/contact", handlers.Repo.Contacts)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user-login", handlers.Repo.ShowLogin)
	mux.Get("/user-logout", handlers.Repo.Logout)
	mux.Post("/user-login", handlers.Repo.PostShowLogin)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", handlers.Repo.AdminDashBoard)

		mux.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		mux.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		mux.Get("/reservations-calendar", handlers.Repo.AdminCalendarReservations)
		mux.Post("/reservations-calendar", handlers.Repo.AdminPostCalendarReservations)
		mux.Get("/processes-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservations)
		mux.Get("/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservations)
		
	
		mux.Get("/reservations/{src}/{id}/show",handlers.Repo.AdminShowReservation )
		mux.Post("/reservations/{src}/{id}",handlers.Repo.AdminPostShowReservation )
	})

	return mux
}
