create-migrate:
	migrate create -ext sql -dir db/migrations -seq create_student_table

migrateup:
	migrate -path db/migrations -database "postgres://docker:password@localhost:5433/school?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://docker:password@localhost:5433/school?sslmode=disable" -verbose down
.PHONY: create-migrate
