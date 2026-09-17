package router

import (
	"github.com/gin-gonic/gin"

	"todo1/internal/handlers"
	"todo1/internal/store"
)

func New(s *store.Store) *gin.Engine {
	r := gin.Default()

	todoHandler := handlers.NewTodoHandler(s)

	r.GET("/", handlers.Root)
	r.GET("/health", handlers.Health)

	r.POST("/todos", todoHandler.Create)
	r.GET("/todos", todoHandler.List)
	r.PUT("/todos/:id", todoHandler.Update)
	r.DELETE("/todos/:id", todoHandler.Delete)

	return r
}
