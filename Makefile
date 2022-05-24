.PHONY: install build clean deploy test uploadTimesheet package
LAMBDAS_DIR = ./lambda

export STAGE=dev
export PROFILE=default

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

prepare-offline-env:
	cp .env.offline .env

install:
	go get ./...
	npm ci

build:
	export GO111MODULE=on
	@for f in $(shell ls ${LAMBDAS_DIR}); do env GOOS=linux go build -ldflags="-s -w" -o bin/lambda/$$f ./lambda/$$f/main.go; done

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	npx serverless deploy --verbose -s $(STAGE) --aws-profile $(PROFILE)

test:
	go test ./internal/**

package: clean build
	npx serverless package --verbose -s $(STAGE)

uploadTimesheet:
	aws s3 cp ./timesheet.yml s3://bamboo-tracker-timesheets-$(STAGE)/ --profile $(PROFILE)
