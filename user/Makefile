.PHONY: proto-gen
proto-gen:
	protoc --proto_path=pkg/grpc/proto \
	--go_out=pkg/grpc/userpb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/grpc/userpb --go-grpc_opt=paths=source_relative \
	pkg/grpc/proto/*.proto

.PHONY: migrate-new
migrate-new:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate create -ext sql -dir /migrations -seq $(NAME)

.PHONY: migrate-up
migrate-up:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate -path /migrations -database $(USER_DB_CONN) -verbose up $(UP)

.PHONY: migrate-down
migrate-down:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate -path /migrations -database $(USER_DB_CONN) -verbose down $(DOWN)

.PHONY: sqlc
sqlc:
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate

.PHONY: mockery-querier
mockery-querier:
	docker run --rm -v $(PWD):/pkg -w /pkg vektra/mockery --case camel --dir pkg/repo --outpkg mocks --output pkg/mocks --name Querier

.PHONY: run
run:
	go run cmd/main.go -debug -db_url=$(USER_DB_CONN) -grpc-port="8082"

.PHONY: postgres
postgres:
	docker run -p 10521:5432 --env POSTGRES_PASSWORD=secret --env POSTGRES_DB=user_service  -d  --name user_service_db postgres:12-alpine
