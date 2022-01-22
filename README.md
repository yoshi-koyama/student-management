# student-management
This is a project to learn Golang and Gin

# Prerequiste
- Installed Go 1.17
- Installed Docker
- Installed golang-migrate
- Installed sqlc (necessary for development)

# Getting Started
- Clone this repository
- Up a postgres docker container
  - `docker-compose up -d`
- Create table
  - `make migrateup`
- Run the command below
  - `go run main.go`

# API Specification
- Move to `/api_dos` directory
- Up an aglio container
  - `docker-compose up -d`
- Navigate to `http://localhost:3000/`
