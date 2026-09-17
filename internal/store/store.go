package store

import (
	"sync"
	"time"

	"todo1/internal/models"
)

type Store struct {
	mu     sync.Mutex
	nextID int
	todos  map[int]*models.Todo
}

func New() *Store {
	return &Store{nextID: 1, todos: make(map[int]*models.Todo)}
}

func (s *Store) Create(title string) *models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	t := &models.Todo{ID: s.nextID, Title: title, Done: false, CreatedAt: time.Now()}
	s.todos[t.ID] = t
	s.nextID++
	return t
}

func (s *Store) List() []*models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]*models.Todo, 0, len(s.todos))
	for _, t := range s.todos {
		out = append(out, t)
	}
	return out
}

func (s *Store) Update(id int, in models.TodoInput) (*models.Todo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.todos[id]
	if !ok {
		return nil, false
	}
	t.Title = in.Title
	if in.Done != nil {
		t.Done = *in.Done
	}
	return t, true
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.todos[id]; !ok {
		return false
	}
	delete(s.todos, id)
	return true
}
