package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func writeJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}

func getIdFromPath(r *http.Request) int {
	vars := mux.Vars(r)
	var id int
	var err error
	id, err = strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}