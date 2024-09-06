qa: analyze test

analyze:
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

build: qa
	@go build -o ./build/voting .

build-docker:
	@docker build -t programmierigel/voting .

coverage: test
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/coverage.out ./...
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	@open ./coverage/coverage.html

docker-push: build-docker
	@docker push programmierigel/voting
	@docker system prune --all --volumes --force

docker-run: docker-push build-docker
	@docker pull programmierigel/voting
	docker run -d -p 3000:3000 -e PASSWORD=123 programmierigel/voting

test:
	@go test -cover ./...

.PHONY: analyze \
	build \
	build-docker \
	coverage \
  docker-push \
  docker-run \
	qa \
	test
