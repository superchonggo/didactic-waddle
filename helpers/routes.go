package helpers

import (
	"net/http"

	"github.com/patrickoliveros/bookings/config"
	"github.com/patrickoliveros/bookings/pages"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	setMiddlewares(mux)
	getRequests(mux)
	postRequests(mux)

	staticFileHandler := enableStaticFiles()
	mux.Handle("/static/*", http.StripPrefix("/static", staticFileHandler))

	return mux
}

func setMiddlewares(mux *chi.Mux) {
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
}

func getRequests(mux *chi.Mux) {
	mux.Get("/", pages.Repo.HomePage)
	mux.Get("/about", pages.Repo.AboutPage)
	mux.Get("/reservations", pages.Repo.ReservationsPage)
	mux.Get("/make-reservation", pages.Repo.MakeReservationPage)
	mux.Get("/reservation-summary", pages.Repo.SummaryMakeReservationPage)
	mux.Get("/contact", pages.Repo.ContactPage)

	mux.Get("/rooms/generals-quarters", pages.Repo.RoomsGeneral)
	mux.Get("/rooms/majors-suite", pages.Repo.RoomsMajor)

	mux.Post("/availability", pages.Repo.AvailabilityReservationsPage)
}

func postRequests(mux *chi.Mux) {
	mux.Post("/reservations", pages.Repo.PostReservationsPage)
	mux.Post("/make-reservation", pages.Repo.PostMakeReservationPage)
}

func enableStaticFiles() http.Handler {
	fileServer := http.FileServer(http.Dir("./static"))

	return fileServer
}
