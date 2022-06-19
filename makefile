localmigrateup:
	migrate -path pkg/db/migration -database "postgresql://postgres@localhost:5432/loggator?sslmode=disable" -verbose up

localmigratedown:
	migrate -path pkg/db/migration -database "postgresql://postgres@localhost:5432/loggator?sslmode=disable" -verbose down
