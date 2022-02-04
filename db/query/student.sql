-- name: CreateStudent :exec
INSERT INTO students (name, nrc, age, gender, password) VALUES ($1, $2, $3, $4, $5);

-- name: ListStudents :many
SELECT * FROM students;

-- name: FindById :one
SELECT * FROM students WHERE id = $1;
