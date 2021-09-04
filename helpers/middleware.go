package helpers

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/patrickoliveros/bookings/config"
)

var AppConfig *config.AppConfig

func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   AppConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return AppConfig.Session.LoadAndSave(next)
}
