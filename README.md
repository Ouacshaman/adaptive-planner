# Adaptive Planner

A lightweight, offline-first daily planner built with Go and PostgreSQL.

Adaptive Planner is a work-in-progress personal task manager designed to run on your own machine and be accessed through a browser. The goal is a focused, responsive planning experience for desktop and mobile without cloud accounts, subscription services, or a heavy frontend framework.

> **Project status:** Early development. The Go server, PostgreSQL-backed task API, and initial HTML page are present. The planning interface and several navigation views are not yet functional.

## Goals

- **Local-first:** Keep application data in a PostgreSQL database you control.
- **Cross-device interface:** Use a browser-based UI that can work on desktop and mobile devices on an accessible local network.
- **Simple stack:** Build with Go, SQL, server-rendered HTML, and progressively enhanced interactions.
- **Maintainable architecture:** Separate HTTP handlers, database access, SQL queries, and templates.

## Current functionality

- Start an HTTP server on port `8080`.
- Store tasks in PostgreSQL.
- Create tasks with `POST /api/tasks`.
- List tasks in creation order with `GET /api/tasks`.
- Render an initial HTML page at `GET /`.
- Run Go tests for response helpers and handlers.

The current homepage contains placeholder task entries. It is **not yet connected** to the task API, and the sidebar links do not yet have corresponding routes.

## Tech stack

| Layer | Technology |
| --- | --- |
| Backend | Go 1.26.4 (version declared in `go.mod`) |
| HTTP | Go standard library (`net/http`) |
| Database | PostgreSQL |
| SQL access | `database/sql`, `lib/pq`, and sqlc-generated queries |
| Rendering | Go `html/template` |
| Frontend | HTML, CSS, and minimal vanilla JavaScript |
| Planned interactivity | HTMX |
| Configuration | Environment variables, optionally loaded from `.env` via `godotenv` |
| Database migration format | Goose-style SQL migrations |

The project intentionally avoids React, Vue, Svelte, TypeScript, Node.js, npm, and frontend build tooling.

## Getting started

### Prerequisites

- Go 1.26.4 or a compatible newer version
- PostgreSQL
- A PostgreSQL database and user with permission to create tables

The repository includes a Goose-format migration but does not currently include a migration runner. You can apply the table definition manually, or use [Goose](https://github.com/pressly/goose).

### 1. Clone the repository

```bash
git clone https://github.com/Ouacshaman/adaptive-planner.git
cd adaptive-planner
```

### 2. Create the database

For example, using `psql` with a PostgreSQL role that has database creation privileges:

```sql
CREATE DATABASE planner;
```

Connect to the new database and create the task table by running the SQL under the `-- +goose Up` marker in `sql/schema/001_plans.sql`:

```sql
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    task TEXT UNIQUE NOT NULL
);
```

Alternatively, with the Goose CLI installed and a suitable PostgreSQL connection string:

```bash
goose -dir sql/schema postgres "$DB_URL" up
```

### 3. Configure environment variables

Create a `.env` file at the repository root:

```dotenv
DB_URL=postgres://YOUR_USER:YOUR_PASSWORD@localhost:5432/planner?sslmode=disable
PLATFORM=dev
```

Replace the placeholders with your local PostgreSQL credentials. Both `DB_URL` and `PLATFORM` are required by the server. The value of `PLATFORM` is currently stored in application configuration but not used to change server behavior.

Do not commit real database passwords or your `.env` file.

### 4. Install dependencies and run the server

From the repository root:

```bash
go mod download
go run ./cmd/server
```

The server starts at:

```text
http://localhost:8080
```

Run the application from the repository root: the current server loads its HTML template using a relative file path.

## API reference

### Create a task

`POST /api/tasks`

```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"task":"Review Go documentation"}'
```

A successful response has HTTP status `200 OK` and includes the new task:

```json
{
  "id": "<generated-uuid>",
  "created_at": "<timestamp>",
  "updated_at": "<timestamp>",
  "task": "Review Go documentation"
}
```

The database currently enforces uniqueness on the `task` text.

### List tasks

`GET /api/tasks`

```bash
curl http://localhost:8080/api/tasks
```

Returns a JSON array of tasks ordered by creation time. Each task includes an ID, creation timestamp, update timestamp, and task description.

**Note:** Update, complete, and delete endpoints are not yet implemented.

## Project structure

```text
adaptive-planner/
├── cmd/
│   └── server/
│       └── main.go              # Server setup, configuration, and routes
├── internal/
│   ├── app/
│   │   └── app.go               # Application configuration
│   ├── database/                # sqlc-generated database code
│   ├── handlers/                # Page and task HTTP handlers
│   └── response/                # JSON request/response helpers
├── sql/
│   ├── queries/
│   │   └── tasks.sql            # Named SQL queries
│   └── schema/
│       └── 001_plans.sql        # Task-table migration
├── web/
│   ├── static/
│   │   ├── css/
│   │   └── js/
│   └── templates/
│       └── pages/
│           └── index.html       # Initial planning page
├── AGENTS.md                    # Development and architecture conventions
├── go.mod
├── go.sum
└── sqlc.yaml
```

## Architecture

The Go server handles HTTP requests using `net/http`. Task handlers call sqlc-generated query methods via an application configuration struct. PostgreSQL persists the data, while the page handler renders HTML with `html/template`.

The intended frontend architecture uses server-rendered templates and HTMX for incremental updates. HTMX interactions are a **design direction**, not an implemented feature of the current repository.

## Development

Run the Go test suite:

```bash
go test ./...
```

Regenerate database access code after updating the SQL schema or queries (requires the sqlc CLI):

```bash
sqlc generate
```

**Development note:** Some tests and error-handling paths are still under development. The commands above describe how to run them, not a guarantee that every test currently passes.

## Roadmap

- [x] Establish the Go HTTP server
- [x] Configure PostgreSQL access
- [x] Add SQL schema and sqlc queries
- [x] Implement task creation and listing APIs
- [x] Add the initial server-rendered homepage
- [ ] Connect the homepage to real task data
- [ ] Add forms and HTMX-powered task interactions
- [ ] Implement task editing, completion, and deletion
- [ ] Build functional Today, Upcoming, Tasks, and Completed views
- [ ] Implement scheduling and adaptive daily planning
- [ ] Make the interface responsive and accessible on mobile
- [ ] Improve test coverage and error handling
- [ ] Document secure local-network access from another device

## Design philosophy

Adaptive Planner is an exploration of building a useful everyday application with a small, understandable web stack. The emphasis is on backend fundamentals, explicit SQL, server-side rendering, and local ownership of data rather than adopting additional frameworks by default.

## License

This project is licensed under the [MIT License](LICENSE). Copyright (c) 2026 Ouacshaman.
