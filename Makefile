.PHONY: db redis dev clean gen-proto deploy rmi build rpc

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
	@docker rmi executor

build:
	@docker build -t executor:latest .\executor\.

rpc:
	@docker run --name executor -d -p 50001:50001 executor

dev: db redis
	@echo "[Dev ready]"

gen-proto:
	@protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./executor/proto/*.proto

deploy:
	@docker compose --env-file .env.production up -d --build