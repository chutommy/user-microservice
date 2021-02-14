.PHONY: postgres
postgres:
	docker run -p 10521:5432 --env POSTGRES_PASSWORD=secret --env POSTGRES_DB=user_service  -d  --name user-svc-db postgres:12-alpine

.PHONY: migrate-new
migrate-new:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate create -ext sql -dir /migrations -seq $(FILENAME)

.PHONY: migrate-up
migrate-up:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate -path /migrations -database $(USER_DB_CONN) -verbose up $(NUM)

.PHONY: migrate-down
migrate-down:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate -path /migrations -database $(USER_DB_CONN) -verbose down $(NUM)

.PHONY: sqlc
sqlc:
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate

.PHONY: test
test:
	go test -v -cover ./...
