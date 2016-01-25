# Cloud GorillaMux Implementation
Back : Go, Gorilla Mux, Google App Engine, Google Datastore.
Front : HTML and jQuery.

## Run
```
go get github.com/gorilla/mux
go install github.com/gorilla/mux

goapp serve src/sfeir/todolist
```
Go to : http://127.0.0.1:8080/todolist/

## On cloud
```
goapp deploy -application todolistgo src/sfeir/todolist
```
**Demo** : http://todolistgo.appspot.com/todolist/