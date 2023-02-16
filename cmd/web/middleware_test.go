package main

import (
	"fmt"
	"net/http"
	"testing"
)

// tests if the parameter is handler type
func TestNoSurf(t *testing.T) {

	var myHandler handler

	handler := NoSurf(&myHandler)

	switch v := handler.(type) {

	case http.Handler:
	default:
		t.Errorf(fmt.Sprintf("type is not http.Handler, it is type %T", v))
	}

}

// tests if the parameter is handler type
func TestSessionLoad(t *testing.T) {

	var myHandler handler

	handler := SessionLoad(&myHandler)

	switch v := handler.(type) {

	case http.Handler:
	default:
		t.Errorf(fmt.Sprintf("type is not http.Handler, it is type %T", v))
	}

}
