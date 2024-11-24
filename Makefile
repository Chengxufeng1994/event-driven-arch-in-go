install-tools:
	@echo installing tools
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@go install github.com/cucumber/godog/cmd/godog@latest
	@go install github.com/pact-foundation/pact-go/v2
	@echo done

# prequisites:
# 	@pact-go -l DEBUG install --libDir /tmp

generate:
	@echo code generation
	@go generate ./...
	@echo done

.PHONY: lint
lint:
	@echo code lint
	@golangci-lint run ./...

.PHONY: test
test:
	@echo code test
	@go test -v -short -race `go list ./... | grep -v /vendor/`

.PHONY: test_e2e
test_e2e:
	@echo code e2e test
	@go test -v -short -race -tags=e2e ./testing/e2e --mono

.PHONY: check-outdated
check-outdated:
	@echo check outdated
	@go list -u -m -json all | go run github.com/psampaz/go-mod-outdated@latest -update -direct

.PHONY: docker-compose-up-mono
docker-compose-up-mono:
	@docker compose \
		--profile monolith \
		--profile testing \
		up -d --remove-orphans

.PHONY: docker-compose-up-microservices
docker-compose-up-microservices:
	@docker compose \
		--profile microservices \
		--profile testing \
		up -d --remove-orphans

.PHONY: docker-compose-down-mono
docker-compose-down-mono:
	@docker compose \
		--profile monolith \
		--profile testing \
		down -v

docker-compose-down-microservices:
	@docker compose \
		--profile microservices \
		--profile testing \
		down -v

build: build-monolith build-microservices

rebuild: clean-monolith clean-microservices build

.PHONY: build-monolith
build-monolith:
	docker build -t mallbots-monolith --file build/docker/Dockerfile .

.PHONY: build-microservices
build-microservices:
	docker build -t mallbots-baskets --file build/docker/Dockerfile.microservices --build-arg=service=basket .
	docker build -t mallbots-cosec --file build/docker/Dockerfile.microservices --build-arg=service=cosec .
	docker build -t mallbots-customers --file build/docker/Dockerfile.microservices --build-arg=service=customer .
	docker build -t mallbots-depot --file build/docker/Dockerfile.microservices --build-arg=service=depot .
	docker build -t mallbots-notifications --file build/docker/Dockerfile.microservices --build-arg=service=notification .
	docker build -t mallbots-ordering --file build/docker/Dockerfile.microservices --build-arg=service=ordering .
	docker build -t mallbots-payments --file build/docker/Dockerfile.microservices --build-arg=service=payment .
	docker build -t mallbots-search --file build/docker/Dockerfile.microservices --build-arg=service=search .
	docker build -t mallbots-stores --file build/docker/Dockerfile.microservices --build-arg=service=store .

.PHONY: clean-monolith
clean-monolith:
	docker image rm mallbots-monolith

.PHONY: clean-microservices
clean-microservices:
	docker image rm mallbots-baskets mallbots-cosec mallbots-customers mallbots-depot mallbots-notifications mallbots-ordering mallbots-payments mallbots-search mallbots-stores

.PHONY: busywork
busywork:
	@go run ./cmd/busywork/ -otlp=localhost:4317 -clients=10

.PHONY: prepare-configmap
prepare-configmap:
	@kubectl delete cm --namespace mallbots initdb
	@kubectl create cm --namespace mallbots initdb --from-file=build/docker/database

.PHONY: k8s-postgres
k8s-postgres:
	@helm upgrade --install postgresql bitnami/postgresql \
		--namespace mallbots \
		--create-namespace \
		--set global.storageClass=nfs-client \
		--set global.postgresql.auth.postgresPassword="postgres" \
		--set architecture=standalone \
		--set primary.service.type=NodePort \
		--set primary.service.nodePorts.postgresql=32345 \
		--set primary.persistence.enabled=true \
		--set primary.persistence.storageClass=nfs-client \
		--set primary.initdb.scriptsConfigMap=initdb \

.PHONY: k8s-nats
k8s-nats:
	@helm upgrade --install nats bitnami/nats \
		--namespace mallbots \
		--create-namespace \
		--set service.type=NodePort \
		--set service.nodePorts.client=32224 \
		--set jetstream.enabled=true \
		--set persistence.enabled=true \
		--set persistence.storageClass=nfs-client
