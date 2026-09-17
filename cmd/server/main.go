package main

import (
	"todo1/internal/router"
	"todo1/internal/store"
)

func main() {
	s := store.New()
	r := router.New(s)
	r.Run(":8080")
}
