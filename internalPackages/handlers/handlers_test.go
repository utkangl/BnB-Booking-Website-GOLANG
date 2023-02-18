package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/utkangl/GoWEB/pkg"
)

type PostData struct {
	key   string
	value string
}

var TheTests = []struct {
	name               string
	url                string
	method             string
	parameters         []PostData
	expectedStatusCode int
}{
	{"home", "/", "GET", []PostData{}, http.StatusOK},
	{"book", "/book", "GET", []PostData{}, http.StatusOK},
	{"about", "/about", "GET", []PostData{}, http.StatusOK},
	{"contact", "/contact", "GET", []PostData{}, http.StatusOK},
	{"kingsuit", "/kings_suit ", "GET", []PostData{}, http.StatusOK},
	{"regularroom", "/regular-room", "GET", []PostData{}, http.StatusOK},
	{"makeres", "/make-reservation", "GET", []PostData{}, http.StatusOK},
	{"resSum", "/reservation-summary", "GET", []PostData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {

	routes := getRoutes()
	TestServer := httptest.NewTLSServer(routes) // create a tlss server to test the website without running the application
	defer TestServer.Close()                    // close the server

	for _, e := range TheTests {
		if e.method == "GET" {
			response, err := TestServer.Client().Get(TestServer.URL + e.url)
			pkg.ErrorNilCheckPrint(err)
			pkg.ErrorNilCheckFatal(err)

			if response.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d, got %d", e.name, e.expectedStatusCode, response.StatusCode)
			}

		}
	}

}
