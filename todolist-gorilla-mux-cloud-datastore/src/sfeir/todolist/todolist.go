package main

import (
	"html/template"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}
    t.Execute(w, nil)
}

func getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := getIdFromPath(r)
	task := getTaskById(id, r)
	if task.Id < 0 {
		writeJson(w, nil)
		return
	}
	writeJson(w, task)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := getTasks(r)
	writeJson(w, tasks)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	var body []byte
	var err error
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 4096))
	if err != nil {
		panic(err)
	}
	err = r.Body.Close()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &task)
	if err != nil {
		panic(err)
	}

	id := addTask(task, r)
	writeJson(w, ResponseStatus{Status: strconv.Itoa(id)})
}

func deleteTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := getIdFromPath(r)
	deleteTaskById(id, r)
	writeJson(w, ResponseStatus{Status: "ok"})
}

func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todolist/", indexHandler).Methods("GET")
	router.HandleFunc("/todolist/tasks/{id:[0-9]+}", getTaskByIdHandler).Methods("GET")
	router.HandleFunc("/todolist/tasks", getTasksHandler).Methods("GET")
	router.HandleFunc("/todolist/tasks", addTaskHandler).Methods("POST")
	router.HandleFunc("/todolist/tasks/{id:[0-9]+}", deleteTaskByIdHandler).Methods("DELETE")
	return router
}

//for app engine
func init() {
	router := initRouter()
    http.Handle("/", router)
}

func main() {
	router := initRouter()
	http.ListenAndServe(":8080", router)
}