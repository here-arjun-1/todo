package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var todos []Todo
var id = 1

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/add", addTodo)
	http.HandleFunc("/delete", deleteTodo)
	http.HandleFunc("/update", updateTodo)
	http.HandleFunc("/complete", completeTodo)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	todo := Todo{
		ID:        id,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, todo)
	id++
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID, _ := strconv.Atoi(r.FormValue("id"))
	for i, todo := range todos {
		if todo.ID == todoID {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	todoID, _ := strconv.Atoi(r.FormValue("id"))
	title := r.FormValue("title")
	for i := range todos {
		if todos[i].ID == todoID {
			todos[i].Title = title
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func completeTodo(w http.ResponseWriter, r *http.Request) {
	todoID, _ := strconv.Atoi(r.FormValue("id"))
	for i := range todos {
		if todos[i].ID == todoID {
			todos[i].Completed = true
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
