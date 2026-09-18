# Todo API

A minimal REST API for managing todos, built in Go with Gin and an in-memory store.

## Run

```bash
go mod download
go run ./cmd/server
```

Server runs on `http://localhost:8080`.

## Endpoints

- `GET /` - API status
- `GET /health` - Health check
- `GET /todos` - List todos
- `POST /todos` - Create todo
- `PUT /todos/:id` - Update todo
- `DELETE /todos/:id` - Delete todo

## Notes

- Data resets on server restart.
- `title` is required for create/update.
