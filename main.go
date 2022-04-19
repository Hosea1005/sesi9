package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi9/controller"
	"sesi9/middleware"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8889"

	fmt.Printf("Running %s\n", server.Addr)
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(w, r) {
		return
	}

	if !middleware.AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); (id) != "" {
		OutputJSON(w, controller.SelectStudent(id))
		return
	}

	OutputJSON(w, controller.GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
