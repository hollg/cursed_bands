.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/tweet-name main.go

deploy: clean build
	sls deploy --verbose
