package main

import (
	"fmt"
	"net/http"
	"strconv"
	"appengine"
	"appengine/datastore"
)

func getTasks(r *http.Request) Tasks {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Task").Order("Id")
	var tasks Tasks
	_, err := q.GetAll(c, &tasks)
	if err != nil {
		panic(err)
	}
	return tasks
}

func getAndIncrementId(c appengine.Context) int {
	var id int
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		key := datastore.NewKey(c, "TaskCounter", "counter", 0, nil)
		var counter Counter
		err := datastore.Get(c, key, &counter)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		id = counter.Count
		fmt.Printf(strconv.Itoa(id))
		counter.Count++
		_, err = datastore.Put(c, key, &counter)
		return err
	}, nil)
	if err != nil {
		panic(err)
	}
	return id
}

func addTask(task Task, r *http.Request) int {
	c := appengine.NewContext(r)
	task.Id = getAndIncrementId(c)
	key := datastore.NewKey(c, "Task", strconv.Itoa(task.Id), 0, nil)
	_, err := datastore.Put(c, key, &task)
	if err != nil {
		panic(err)
	}
	return task.Id
}

func getTaskById(id int, r *http.Request) Task {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Task").Filter("Id =", id)
	var tasks Tasks
	_, err := q.GetAll(c, &tasks)
	if err != nil {
		panic(err)
	}
	return tasks[0]
}

func deleteTaskById(id int, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Task", strconv.Itoa(id), 0, nil)
	err := datastore.Delete(c, key)
	if err != nil {
		panic(err)
	}
}
