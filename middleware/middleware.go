package middleware

import (
	"net/http"
)

const (
	__USERNAME__ = "hosea12"
	__PASSWORD__ = "hosea"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("something went wrong"))
		return false
	}

	isValid := (username == __USERNAME__) && (password == __PASSWORD__)
	if !isValid {
		w.Write([]byte("wrong username/password"))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}
