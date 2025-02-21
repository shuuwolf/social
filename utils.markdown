### Force version ###
migrate -path="./cmd/migrate/migrations" -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" force 1