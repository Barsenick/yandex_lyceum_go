package main

import "time"

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() (is bool) {
	return time.Now().After(t.deadline)
}

func (t Task) IsTopPriority() (is bool) {
	return t.priority >= 4
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (list ToDoList) TasksCount() int {
	return len(list.tasks)
}

func (list ToDoList) NotesCount() int {
	return len(list.notes)
}

func (list ToDoList) CountTopPrioritiesTasks() int {
	count := 0

	for _, task := range list.tasks {
		if task.IsTopPriority() {
			count++
		}
	}

	return count
}

func (list ToDoList) CountOverdueTasks() int {
	count := 0

	for _, task := range list.tasks {
		if task.IsOverdue() {
			count++
		}
	}

	return count
}
