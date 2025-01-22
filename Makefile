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

migrateup1:
	migrate -path db/migration -database "postgres://your_username:your_password@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://your_username:your_password@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://your_username:your_password@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination=db/mock/store.go github.com/soojae/simplebank/db/sqlc Store

proto:
	rm -rf pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simplebank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl


redis:
	docker run --name redis -p 6379:6379 -d redis:7.4.2-alpine
#
#migrate
#
#rollback:
#	migrate -database postgres://your_username:your_password@localhost:5432/simple_bank?sslmode=disable -path migration down

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 test server mock proto evans redis
#migrate rollback

# command in gogo folder
# make migrate