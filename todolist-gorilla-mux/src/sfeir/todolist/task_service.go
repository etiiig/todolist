package main

import (

)

var tasks map[int]Task = make(map[int]Task)
var currentId int = 0

func getTasks() Tasks {
	values := make(Tasks, 0, len(tasks))
	for v := range tasks {
		values = append(values, tasks[v])
	}
	return values
}

func addTask(task Task) int {
	task.Id = currentId
	tasks[currentId] = task
	currentId++
	return currentId - 1
}

func getTaskById(id int) Task {
	return tasks[id]
}

func deleteTaskById(id int) {
	delete(tasks, id);
}
