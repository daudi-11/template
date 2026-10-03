# Backend

This Go service provides account registration, login, logout, and session validation for the template.

## Setup

Configure `AUTH_DB_URL` with a PostgreSQL URL and set a unique `JWT_SECRET` of at least 32 characters in `.env`. Start the service from this directory so the `.env` file is loaded:

```powershell
if (!(Test-Path .env)) { Copy-Item .env.example .env }
go mod download
go run .
```

`PORT` defaults to `3030`; `JWT_EXPIRATION_HOURS` defaults to `1`. Account credentials and hashed session tokens are stored in the authentication database.

## Auth entry points

- `routes/routes.go` registers the public auth endpoints and protected logout endpoint.
- `controllers/authController.go` handles registration, login, and logout.
- `middleware/auth.go` verifies JWTs and active sessions.
- `config/db.go` opens the database and prepares the auth tables.
