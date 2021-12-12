package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//Just an example test to see what we can achieve. Should ideally have tests for all handlers involved.
func TestLoginFailure(t *testing.T) {
	data := url.Values{}
	data.Set("name", "lemma")
	data.Set("password", "wrongPassword")
	req, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 401 {
		t.Errorf("Wrong status code: got %v wanted %v",
			status, 401)
	}

	expected := "Invalid User\n"
	if rr.Body.String() != expected {
		t.Errorf("response: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

func TestLoginSuccess(t *testing.T) {
	data := url.Values{}
	data.Set("name", "lemma")
	data.Set("password", "@dm1n")
	req, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 302 {
		t.Errorf("Wrong status code: got %v wanted %v",
			status, 302)
	}
}
