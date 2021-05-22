# Bamboo tracker
![Test](https://github.com/ivanmartos/bamboo-tracker/workflows/Test/badge.svg)


## Overview
This project contains serverless application deployable to AWS for automatic tracking of worked hours in [BambooHR](https://www.bamboohr.com/).
Project also contains functionality for notifying about current tracked timesheets per day / week.
Project is implemented using [GoLang](https://golang.org/) programming language and [serverless framework](https://serverless.com/).

## Content
Current project creates CloudFormation stack that contains 
- AWS StepFunction that orchestrates lambda steps for uploading timesheet and notifying about emails
- s3 bucket for storing timesheets
- EventBridge for triggering StepFunction

Project also contains definition of terraform stack
- verification of SES email identity (it's not part of CloudFormation stack since this is not supported by CloudFormation).
This is for sending notification emails about current time tracking values. This terraform stack should be applied before deploying 
CloudFormation stack.

### Upon invocation lambda for uploading timesheet:
1. Parses timesheet entries from yaml file in s3 bucket (by default it retrieves file by key `timesheet.yml`)
2. If there is current day in parsed config it continues. Otherwise it terminates
3. Logs in to BambooHR using credentials in environment variables
4. Uploads timesheet entries for current day to bambooHR
5. Terminates

### Upon invocation lambda for notifying about current timesheet values:
1. Logs in to BambooHR using credentials in environment variables
2. Retrieves logged working hours for current day and week
3. Sends an email using AWS SES to/from emails specified as environment variables with HTML content of logged hours.
4. Terminates

## Tutorial
After cloning repository it is require to install dependencies.
```
make install
```
This is required one time only.

Also create a copy of timesheet entries
```
cp timesheet.dst.yml timesheet.yml
```
And modify these timesheet according to your need. **Important** - name of the weekdays must be lowercase

Update your custom values in [terraform](./tf-infrastructure) directory and execute in that directory
```
terraform apply
```

### How to deploy
Environment variables needed before deployment
1. `BAMBOO_HOST` - environment variable containing hostname of your company bambooHR
2. `BAMBOO_USERNAME` - your username used for logging in to BambooHR
3. `BAMBOO_PASSWORD` - your password used for logging in to BambooHR
4. `DAILY_TIME_TRACKING_SENDER_EMAIL` - email address to send time tracking notifications from
%. `DAILY_TIME_TRACKING_RECIPIENT_EMAIL` - email address to receive time tracking notifications from

```
make deploy STAGE=YOUR_STAGE //STAGE is optional, be default it is "dev"
```

You can copy content of your local `timesheet.yml` to s3 by running
```
make uploadTimesheet
```

### How to run locally

To just execute timesheetUploader functionality without serverless framework start program [locally](cmd/timesheetUploaderLocal)

### How to run tests
Execute command 
```
make test
```


### Disclaimer
Content of this project was created only for research purposes. 
Definitely not for automatic tracking of working hours due to laziness.
