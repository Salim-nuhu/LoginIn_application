# Login Form API

A simple authentication REST API built with Go.

## Endpoints
- `POST /register` — Register a new user
- `POST /login` — Login and receive a JWT token

## Setup
1. Clone the repository
2. Create a `.env` file with your database credentials
3. Run `go mod tidy` to install dependencies
4. Run `swag init` to generate Swagger docs
5. Run `go run main.go`

## Documentation
Visit `http://localhost:8080/swagger/index.html`
