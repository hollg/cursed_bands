.PHONY: build deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/tweet-name main.go

deploy: build
	sls deploy --verbose
