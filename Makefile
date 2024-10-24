pgadmin:
	docker run --name paadmin \
		-p 5050:80 -d \
    -e "PGADMIN_DEFAULT_EMAIL=admin@gmail.com" \
    -e "PGADMIN_DEFAULT_PASSWORD=admin" \
    -d dpage/pgadmin4

postgres:
	docker run --name postgres \
		-p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d \
		-v pgdata:/var/lib/postgresql/data \
		-v $(PWD)/build/docker/database:/docker-entrypoint-initdb.d/ \
		postgres:14-alpine

generate:
	@echo ==> running code generation
	@go generate ./...
	@echo done
