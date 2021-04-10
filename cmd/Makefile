all: clean build run
build:
	docker-compose build --no-cache
run:
	docker-compose up -d
	docker run -it --rm cmd_app
clean:
	docker-compose down


