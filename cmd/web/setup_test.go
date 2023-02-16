package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

// http handler interface
type handler struct{}

func (handler *handler) ServeHTTP(Res http.ResponseWriter, Req *http.Request) {

}
