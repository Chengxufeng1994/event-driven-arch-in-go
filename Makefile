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

.PHONY: check-outdated
check-outdated:
	@echo check outdated
	@go list -u -m -json all | go run github.com/psampaz/go-mod-outdated@latest -update -direct

.PHONY: docker-compose-up
docker-compose-up:
	@docker compose up -d --remove-orphans

.PHONY: docker-compose-down
docker-compose-down:
	@docker compose down -v

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
