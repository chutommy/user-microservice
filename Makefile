.PHONY: build, run

build:
	docker-compose -f docker-compose.yml -p booking-terminal build

run:
	docker-compose -f docker-compose.yml -p booking-terminal up
