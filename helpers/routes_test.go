package helpers

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/patrickoliveros/bookings/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := Routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}
