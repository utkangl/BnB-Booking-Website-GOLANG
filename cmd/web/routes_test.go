package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/utkangl/GoWEB/internalPackages/config"
)

// tests if the app variable's type is valid
func TestRoutes(t *testing.T) {

	var app config.AppConfig

	mux := routes(&app)

	switch typ := mux.(type) {
	case *chi.Mux:
	default:
		t.Errorf(fmt.Sprintf("type is not *chi.Mux, but is %T", typ))
	}
}
