package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {

	var myHandler handler

	handler := NoSurf(&myHandler)

	switch v := handler.(type) {

	case http.Handler:
	default:
		t.Errorf(fmt.Sprintf("type is not http.Handler, it is type %T", v))
	}

}

func TestSessionLoad(t *testing.T) {

	var myHandler handler

	handler := SessionLoad(&myHandler)

	switch v := handler.(type) {

	case http.Handler:
	default:
		t.Errorf(fmt.Sprintf("type is not http.Handler, it is type %T", v))
	}

}
