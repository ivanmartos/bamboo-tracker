.PHONY: build clean deploy test

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/timesheetUploader ./cmd/timesheetUploader/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --verbose

offline: clean build
	serverless offline --useDocker

test:
	go test ./internal/**