This is a simple Golang CRUD project using Gin, SQLC with PostgreSQL database.

To run migration
`migrate -path db/migration -database  "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose up`

To revert migration
`migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose down`

Make sure to set environment variables in the `app.env` at root path

```
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@go_postgres:5432/contact_db?sslmode=disable
POSTGRES_USER=root
POSTGRES_PASSWORD=secret
POSTGRES_DB=contact_db
SERVER_ADDRESS=8000
```
