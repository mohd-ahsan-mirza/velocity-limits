all: clean build run
build:
	docker-compose build --no-cache
run:
	docker-compose up -d
	sleep 10
	docker run -it --network=app-network --rm cmd_app
testrun:
	docker-compose up -d
	sleep 10
	docker run -it --network=app-network --rm cmd_test
test: clean build testrun
clean:
	docker-compose down


