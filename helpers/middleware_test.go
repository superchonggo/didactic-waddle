package helpers

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var tstHandler testHandler

	h := NoSurf(&tstHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not httpHandler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var tstHandler testHandler
	h := SessionLoad(&tstHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler but is %T", v))
	}
}
