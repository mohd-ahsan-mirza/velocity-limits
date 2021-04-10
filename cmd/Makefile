all: clean build run
build:
	docker-compose build --no-cache
run:
	docker-compose up -d
	sleep 10
	docker run -it --network=cmd_default --rm cmd_app
clean:
	docker-compose down


