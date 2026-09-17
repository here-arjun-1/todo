package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"todo1/internal/models"
	"todo1/internal/store"
)

type TodoHandler struct {
	store *store.Store
}

func NewTodoHandler(s *store.Store) *TodoHandler {
	return &TodoHandler{store: s}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var in models.TodoInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, h.store.Create(in.Title))
}

func (h *TodoHandler) List(c *gin.Context) {
	c.JSON(200, h.store.List())
}

func (h *TodoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	var in models.TodoInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	t, ok := h.store.Update(id, in)
	if !ok {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}
	c.JSON(200, t)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if !h.store.Delete(id) {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}
	c.Status(204)
}
