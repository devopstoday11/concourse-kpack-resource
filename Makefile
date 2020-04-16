docker_registry = gcr.io/cf-build-service-public/concourse-kpack-resource

docker:
	docker build -t $(docker_registry) .

publish: docker
	docker push $(docker_registry)

test:
	go test -v ./...

fmt:
	find . -name '*.go' | while read -r f; do \
		gofmt -w -s "$$f"; \
	done

.DEFAULT_GOAL := docker

.PHONY: go-mod docker-build docker-push docker test fmt
