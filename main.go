package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/patrickoliveros/bookings/config"
	"github.com/patrickoliveros/bookings/handlers"
	"github.com/patrickoliveros/bookings/helpers"
	"github.com/patrickoliveros/bookings/models"
	"github.com/patrickoliveros/bookings/pages"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	// we need to tell the application the types that we are going to store in the session
	gob.Register(models.Reservation{})

	runApplication()

	log.Printf("Starting application on port %s", app.PortNumber)

	srv := &http.Server{
		Addr:    app.PortNumber,
		Handler: helpers.Routes(&app),
	}

	err := srv.ListenAndServe()
	log.Println(err)
}

func runApplication() error {
	setupApplicationConfig()
	setupDependencies()

	err := setupApplicationTemplates()

	setupRepo()
	setupSession()

	return err
}

func setupApplicationConfig() {
	app.UseSecure = false
	app.InProduction = false
	app.UseCache = false
	app.PortNumber = ":8080"

	infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
}

func setupApplicationTemplates() error {
	tc, err := handlers.CreateAllTemplatesCache()
	helpers.HandleFatalError(err, "cannot create template cache")

	app.TemplateCache = tc

	handlers.NewTemplates(&app)

	return err
}

func setupDependencies() {
	helpers.AppConfig = &app
}

func setupRepo() {
	repo := pages.NewRepo(&app)
	pages.NewHandlers(repo)
}

func setupSession() {
	session = scs.New()
	session.Lifetime = 23 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
}
