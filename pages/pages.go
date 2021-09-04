package pages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/patrickoliveros/bookings/config"
	"github.com/patrickoliveros/bookings/forms"
	"github.com/patrickoliveros/bookings/handlers"
	"github.com/patrickoliveros/bookings/logging"
	"github.com/patrickoliveros/bookings/models"
)

//region repo

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func (r *Repository) AddSessionError(req *http.Request, message string) {
	fmt.Println("I am called")
	r.App.Session.Put(req.Context(), "error", message)
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//endregion

func (m *Repository) HomePage(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	handlers.RenderPageWithTemplate(w, r, "home", &models.TemplateData{
		StringMap: stringMap,
		CSRFToken: "token value",
	})
}

func (m *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	handlers.RenderPageWithTemplate(w, r, "about", &models.TemplateData{})
}

func (m *Repository) RoomsGeneral(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	handlers.RenderPageWithTemplate(w, r, "generals-quarters", &models.TemplateData{
		StringMap: stringMap,
		PageTitle: "General's Quarters",
	})
}

func (m *Repository) RoomsMajor(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	handlers.RenderPageWithTemplate(w, r, "majors-suite", &models.TemplateData{
		StringMap: stringMap,
		CSRFToken: "token value",
		PageTitle: "Major's Suite",
	})
}

func (m *Repository) ContactPage(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	handlers.RenderPageWithTemplate(w, r, "contact", &models.TemplateData{
		StringMap: stringMap,
		PageTitle: "Contact Page",
	})
}

func (m *Repository) ReservationsPage(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	handlers.RenderPageWithTemplate(w, r, "reservations", &models.TemplateData{
		StringMap: stringMap,
		PageTitle: "Reservations",
	})
}

func (m *Repository) PostReservationsPage(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")

	w.Write([]byte(fmt.Sprintf("Posted to search reservations. Start: %s, End: %s", start, end)))
}

func (m *Repository) AvailabilityReservationsPage(w http.ResponseWriter, r *http.Request) {
	resp := models.JsonResponse{
		OK:      true,
		Message: "some message for you",
	}

	start := r.Form.Get("start_date")

	if start == "" {
		start = r.Form.Get("start")
	}

	end := r.Form.Get("end_date")

	if end == "" {
		end = r.Form.Get("end")
	}

	log.Println(r.Form)

	resp.Body = start + " / " + end

	out, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) MakeReservationPage(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	handlers.RenderPageWithTemplate(w, r, "make-reservation", &models.TemplateData{
		Data:      data,
		PageTitle: "Make Reservation",
		Form:      forms.New(nil),
	})
}

func (m *Repository) PostMakeReservationPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	// err = errors.New("this is an error message")

	if err != nil {
		logging.ServerError(w, err)
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 10)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		handlers.RenderPageWithTemplate(w, r, "make-reservation", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) SummaryMakeReservationPage(w http.ResponseWriter, r *http.Request) {

	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("Cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		// TODO: validate this
		// m.AddSessionError(r, "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	handlers.RenderPageWithTemplate(w, r, "reservation-summary", &models.TemplateData{
		PageTitle: "Reservation Summary",
		Data:      data,
	})
}
