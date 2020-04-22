.PHONY: install build clean deploy test

install:
	go get ./...
	npm install

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/timesheetUploader ./cmd/timesheetUploader/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --verbose -s $(STAGE)

offline: clean build
	serverless offline --useDocker -s local

test:
	go test ./internal/**
