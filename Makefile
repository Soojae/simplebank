postgres:
	docker-compose up -d

createdb:
	docker exec -it postgres_container createdb psql -U your_username -d simple_bank
#-c "CREATE DATABASE simple_bank;"

dropdb:
	docker exec -it postgres_container drobdb psql -U your_username -d simple_bank
#-c "DROP DATABASE simple_bank;"

migrateup:
	migrate -path db/migration -database "postgres://your_username:your_password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://your_username:your_password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination=db/mock/store.go github.com/soojae/simplebank/db/sqlc Store

#
#migrate
#
#rollback:
#	migrate -database postgres://your_username:your_password@localhost:5432/simple_bank?sslmode=disable -path migration down

.PHONY: postgres createdb dropdb migrateup migratedown test server mock
#migrate rollback

# command in gogo folder
# make migrate