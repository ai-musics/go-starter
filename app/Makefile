.PHONY:

build:
	docker build -f Dockerfile -t starter:v1 .

run:
	docker run --name starter -it --rm -p 3001:3001  starter:v1

stop:
	docker stop starter
