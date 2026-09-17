package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Todo")
	http.ListenAndServe(":8000", nil)
}

type Todo struct {
	ID     int
	Title  string
	Status bool
}

var todos []Todo
