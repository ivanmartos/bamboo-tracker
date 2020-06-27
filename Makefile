.PHONY: install build clean deploy test uploadTimesheet
LAMBDAS_DIR = ./lambda

export STAGE=dev
export PROFILE=default

install:
	go get ./...
	npm install

build:
	export GO111MODULE=on
	@for f in $(shell ls ${LAMBDAS_DIR}); do env GOOS=linux go build -ldflags="-s -w" -o bin/lambda/$$f ./lambda/$$f/main.go; done

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
