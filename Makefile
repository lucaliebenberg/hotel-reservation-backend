build:
	@go build -o bin/api

run:
	@./bin/api

seed:
	@go run scripts/seed.go

docker:
	echo "building docker file"
	@docker build -t api .
	echo "running API inside docker container"
	@docker run -p 3000:3000 api

test:
	@go test -v ./... 