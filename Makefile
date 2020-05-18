.PHONY: install build clean deploy test uploadTimesheet

export STAGE=dev
export PROFILE=default

install:
	go get ./...
	npm install

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/timesheetUploader ./cmd/timesheetUploader/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/timeTrackingChecker ./cmd/timeTrackingChecker/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --verbose -s $(STAGE) --aws-profile $(PROFILE)

offline: clean build
	serverless offline --useDocker -s local

test:
	go test ./internal/**

uploadTimesheet:
	aws s3 cp ./timesheet.yml s3://bamboo-tracker-timesheets-$(STAGE)/ --profile $(PROFILE)
