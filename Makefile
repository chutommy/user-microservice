.PHONY:

postgres:
	docker run -p 5432:5432 --env POSTGRES_PASSWORD=secret --env POSTGRES_DB=booking_terminal --name postgres12 -d postgres:12-alpine

migrate-new:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate create -ext sql -dir /migrations -seq $(FILENAME)

migrate-up:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate -path /migrations -database $(BOOKING_TERMINAL_DB_CONN) -verbose up

migrate-down:
	docker run --rm -v $(PWD)/db/schema:/migrations --network host migrate/migrate -path /migrations -database $(BOOKING_TERMINAL_DB_CONN) -verbose drop -f

sqlc:
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate

#mock:
#	docker run --rm -v $(PWD):/pkg -w /pkg vektra/mockery --case camel --dir db/sqlc --name Store --outpkg mocks --output db/mocks

test:
	go test -v -cover ./...

server:
	go run main.go
