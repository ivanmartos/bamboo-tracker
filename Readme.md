# Bamboo tracker
![Test](https://github.com/ivanmartos/bamboo-tracker/workflows/Test/badge.svg)


## Overview
This project contains serverless application deployable to AWS for automatic tracking of worked hours in [BambooHR](https://www.bamboohr.com/)
Project is implemented using [GoLang](https://golang.org/) programming language and [serverless framework](https://serverless.com/).

## Content
Current project creates CloudFormation stack that contains one lambda. This lambda is executed periodically via 
CloudWatch Events. Upon invocation it:
1. Parses timesheet entries from config (environment variable)
2. If there is current day in parsed config it continues. Otherwise it terminates
3. Logs in to BambooHR using credentials in environment variables
4. Uploads timesheet entries for current day to bambooHR
5. Terminates

## Tutorial
After cloning repository it is require to install dependencies.
```
make install
```
This is required one time only.

### How to deploy
Before you deploy, you need to pass your own configuration to file [serverless.yml](serverless.yml)

What needs to be modified:
1. `BAMBOO_HOST` - environment variable containing hostname of your company bambooHR
2. `BAMBOO_USERNAME` - your username used for logging in to BambooHR
3. `BAMBOO_PASSWORD` - your password used for logging in to BambooHR
4. `TIMESHEETS` - yml formatted string that contains timesheet entries to be uploaded. **Important** - name of the weekdays must be lowercase

Other defaults (can be overriden in [serverless.yml](serverless.yml))
- lambda will be executed every working day at 5pm UTC
- CloudFormation stack will be deployed to `eu-west-1` region

```
make deploy STAGE=YOUR_STAGE //STAGE is optional, be default it is "dev"
```

### How to run locally
To start as local serverless project execute
```
make offline
```

To just execute timesheetUploader functionality without serverless framework start program [locally](cmd/timesheetUploaderLocal)

### How to run tests
Execute command 
```
make test
```


### Disclaimer
Content of this project was created only for research purposes. 
Definitely not for automatic tracking of working hours due to laziness.
