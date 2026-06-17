# go-microservices-course-auth-service

Go authentication service built as part of a microservices learning project.

## Project purpose

This repository contains a small gRPC-based authentication/user service created for coursework and microservices practice. It demonstrates how to structure a Go service, expose a protobuf API, connect to PostgreSQL, and run local infrastructure with Docker Compose.

## Learning context

The project is an archived educational repository. It is kept public to show implementation progress and course exercises rather than to provide a production-ready authentication platform.

## Tech stack

- Go 1.21
- gRPC and Protocol Buffers
- PostgreSQL 14
- `pgx` for PostgreSQL access
- `squirrel` for SQL query building
- `goose` for database migrations
- Docker and Docker Compose
- GitHub Actions workflow configuration

## Service responsibilities

The service owns basic user records and exposes gRPC operations to:

- create a user with name, email, password confirmation, and role;
- fetch a user by ID;
- update a user's name and/or email;
- delete a user by ID.

> Note: the current educational implementation validates password confirmation but does not persist password hashes. Do not use it as-is for production authentication.

## Repository rename

The Go module and repository references use:

```text
github.com/HpPpL/go-microservices-course-auth-service
```

## Local run instructions

1. Install Go 1.21+, Docker, Docker Compose, `make`, and Protocol Buffers tooling if you plan to regenerate protobuf files.
2. Create a local environment file:

   ```bash
   cp .env.example .env
   ```

3. Review `.env` and change placeholder values for your local machine.
4. Start PostgreSQL:

   ```bash
   docker compose up -d pg
   ```

5. Install local development tools when needed:

   ```bash
   make install-deps
   ```

6. Run migrations:

   ```bash
   make local-migration-up
   ```

7. Run the gRPC server:

   ```bash
   go run ./cmd/grpc_server --config-path .env
   ```

8. In another terminal, run the sample client if desired:

   ```bash
   go run ./cmd/grpc_client --config-path .env
   ```

## Environment variables

| Variable | Purpose | Example |
| --- | --- | --- |
| `PG_DATABASE_NAME` | PostgreSQL database name used by Docker Compose | `auth` |
| `PG_USER` | PostgreSQL user used by Docker Compose | `auth_user` |
| `PG_PASSWORD` | PostgreSQL password for local development | `change_me_local_password` |
| `PG_PORT` | Host port mapped to PostgreSQL | `54321` |
| `PG_DSN` | PostgreSQL DSN used by the Go service | `host=localhost port=54321 dbname=auth user=auth_user password=change_me_local_password sslmode=disable` |
| `MIGRATION_DIR` | Directory containing Goose migrations | `./migrations` |
| `MIGRATION_DSN` | PostgreSQL DSN used by migration scripts | `host=pg port=5432 dbname=auth user=auth_user password=change_me_local_password sslmode=disable` |
| `GRPC_HOST` | gRPC bind/connect host | `localhost` |
| `GRPC_PORT` | gRPC port | `50051` |

Only `.env.example` is tracked. Local `.env`, `prod.env`, and other environment files are intentionally ignored.

## API overview

The protobuf definition lives in `api/auth_v1/auth.proto` and defines the `AuthV1` service:

- `Create(CreateRequest) returns (CreateResponse)`
- `Get(GetRequest) returns (GetResponse)`
- `Update(UpdateRequest) returns (google.protobuf.Empty)`
- `Delete(DeleteRequest) returns (google.protobuf.Empty)`

Generated Go stubs are committed under `pkg/auth_v1`.

## Repository status and limitations

- Archived learning/coursework project.
- Not production-ready.
- No real credentials should be stored in Git.
- Historical environment files contained placeholder local values only during this cleanup review; rotate any values before reuse if they were ever replaced with real credentials in another clone or deployment.
- Repository description suggestion: `Go authentication service built as part of a microservices learning project.`
- Suggested topics: `go`, `golang`, `microservices`, `auth-service`, `authentication`, `jwt`, `postgresql`, `docker`, `coursework`, `learning`.
