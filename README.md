# Notes API (`notes-api`)

A production-structured RESTful API for managing notes, built in Go using the standard library (`net/http`, `database/sql`, `context`, `os/signal`) and PostgreSQL (`pgx/v5/stdlib`).

## Features

- **Full CRUD REST Endpoints:** Create, list, retrieve, partially update (`PATCH`), and delete notes.
- **Layered Architecture:** Clean separation between HTTP transport, business logic, and storage interfaces.
- **PostgreSQL Persistence:** Parameterized queries (`$1`, `$2`), `INSERT ... RETURNING`, `COALESCE`-based partial updates, and tuned `*sql.DB` connection pooling.
- **Context Propagation:** Request-scoped `context.Context` passed from HTTP handlers through the application layer down to database queries.
- **Production HTTP Hardening:** Explicit `ReadTimeout`, `WriteTimeout`, and `IdleTimeout` limits, paired with `SIGINT`/`SIGTERM` graceful shutdown via `signal.NotifyContext`.
- **Unit & Integration Test Suite:** 41 race-tested unit and HTTP routing tests backed by a thread-safe `InMemoryStorage` implementation.

---

## Architecture

The application follows a strict unidirectional dependency flow where outer layers depend on inner abstractions:

```text
HTTP Request
     │
     ▼
┌─────────────┐
│   Router    │  net/http ServeMux (method & path routing)
└──────┬──────┘
       │
       ▼
┌─────────────┐
│   Handler   │  JSON decoding/encoding, path parsing, HTTP status mapping
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ Application │  Input validation, timestamps, business rules, domain errors
└──────┬──────┘
       │  (Storage Interface)
       ├──────────────────────────────┐
       ▼                              ▼
┌───────────────────┐        ┌──────────────────┐
│  PostgresStorage  │        │ InMemoryStorage  │
│ (internal/storage)│        │    (storage)     │
└────────┬──────────┘        └──────────────────┘
         │
         ▼
   PostgreSQL DB
```

### Error Mapping Across Layers

| Storage / Domain Error                                       | HTTP Status                 | JSON Response                                              |
| :----------------------------------------------------------- | :-------------------------- | :--------------------------------------------------------- |
| `application.ErrInvalidInput` / Malformed JSON               | `400 Bad Request`           | `{"error": "invalid input"}` / `{"error": "invalid JSON"}` |
| `sql.ErrNoRows` -> `ErrStorageNotFound` -> `ErrNoteNotFound` | `404 Not Found`             | `{"error": "note not found"}`                              |
| Unsupported HTTP Method                                      | `405 Method Not Allowed`    | `{"error": "method not allowed"}`                          |
| Unexpected Internal / Database Failure                       | `500 Internal Server Error` | `{"error": "internal server error"}`                       |

---

## Project Structure

```text
notes-api/
├── main.go                  # Entry point, dependency wiring, http.Server & graceful shutdown
├── router.go                # HTTP route definitions (Go 1.22+ ServeMux patterns)
├── router_test.go           # End-to-end HTTP routing tests
├── go.mod
├── go.sum
├── note/
│   └── note.go              # Core Note domain struct
├── handler/
│   ├── handler.go           # HTTP handlers, request parsing, error-to-status mapping
│   └── handler_test.go      # Handler unit tests using httptest
├── application/
│   ├── application.go       # Business logic, validation, Storage interface definition
│   └── application_test.go  # Application unit tests with FakeStorage
├── storage/
│   ├── inmemory.go          # Thread-safe in-memory Storage implementation (sync.RWMutex)
│   └── inmemory_test.go     # InMemoryStorage unit tests
└── internal/
    ├── config/
    │   └── config.go        # Environment variable loading & PostgreSQL DSN builder
    ├── database/
    │   └── database.go      # *sql.DB initialization, ping check & connection pool tuning
    └── storage/
        └── postgres.go      # PostgreSQL implementation of application.Storage
```

---

## Prerequisites & Database Setup

- **Go:** 1.24 or newer
- **PostgreSQL:** 14 or newer

Create a PostgreSQL database and run the following schema migration to create the `notes` table:

```sql
CREATE TABLE IF NOT EXISTS notes (
    id         INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title      TEXT        NOT NULL,
    content    TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
```

---

## Running the Application

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Mayankkaira/notes-api.git
   cd notes-api
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Export required PostgreSQL environment variables:**

   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=your_password
   export DB_NAME=notes_db
   ```

4. **Start the server:**
   ```bash
   go run .
   ```
   The server listens on `:8080`. Press `Ctrl+C` to trigger a graceful shutdown (drains active requests for up to 10 seconds before closing the database pool).

---

## Running Tests

Run the full test suite with Go's race detector enabled (uses `InMemoryStorage`, no live PostgreSQL instance required):

```bash
go test -race -v ./...
```

---

## API Reference

| Method   | Endpoint      | Description                               | Success Status   |
| :------- | :------------ | :---------------------------------------- | :--------------- |
| `POST`   | `/notes`      | Create a new note                         | `201 Created`    |
| `GET`    | `/notes`      | Retrieve all notes                        | `200 OK`         |
| `GET`    | `/notes/{id}` | Retrieve a note by ID                     | `200 OK`         |
| `PATCH`  | `/notes/{id}` | Partially update `title` and/or `content` | `200 OK`         |
| `DELETE` | `/notes/{id}` | Delete a note by ID                       | `204 No Content` |

### Example `curl` Requests

**Create a note:**

```bash
curl -i -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "First Note", "content": "Stored in PostgreSQL"}'
```

**List all notes:**

```bash
curl -i http://localhost:8080/notes
```

**Get a single note:**

```bash
curl -i http://localhost:8080/notes/1
```

**Partially update a note (`PATCH`):**

```bash
curl -i -X PATCH http://localhost:8080/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated Title Only"}'
```

**Delete a note:**

```bash
curl -i -X DELETE http://localhost:8080/notes/1
```
