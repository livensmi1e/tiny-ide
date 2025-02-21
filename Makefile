.PHONY: db redis dev clean gen-proto

db:	
	@docker run --name postgres -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:17.2-alpine3.21

redis:
	@docker run --name redis -d -p 6379:6379 redis:7.4.2-alpine

clean:
	@docker stop postgres && \
	docker rm postgres -v && \
	docker stop redis && \
	docker rm redis -v

rmi:
	@docker rmi sandbox

build:
	@docker build -t sandbox:latest .\pkg\cee\.

dev: db redis
	@echo "[Dev ready]"

gen-proto:
	@protoc -I. --go_out=. --go-grpc_out=. ./executor/proto/*.proto