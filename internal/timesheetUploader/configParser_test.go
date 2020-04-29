package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestGetTimesheetEntries(t *testing.T) {
	verifyTimesheetEntry := func(t *testing.T, parser TimesheetParser, weekday time.Weekday, expectedEntries []TimesheetEntry) {
		result := parser.GetTimesheetEntries(weekday)

		if len(expectedEntries) == 0 {
			if len(result) != 0 {
				t.Errorf("Expected len of result array for %v to be 0, got %v", weekday, len(result))
			}
		} else {
			if !reflect.DeepEqual(result, expectedEntries) {
				t.Errorf("Expected result of %v entries was %v, got %v", weekday, expectedEntries, result)

			}
		}
	}

	verifyS3Call := func(t *testing.T, repository mocks.MockS3Repository, expectedKey string, expectedBucket string) {
		if repository.GetS3FileContentBucketArg != expectedBucket {
			t.Errorf("Expected s3 bucket argument in repository call %v, got %v", expectedBucket, repository.GetS3FileContentBucketArg)
		}

		if repository.GetS3FileContentKeyArg != expectedKey {
			t.Errorf("Expected s3 key argument in repository call %v, got %v", expectedKey, repository.GetS3FileContentKeyArg)
		}
	}

	mockS3Repo := &mocks.MockS3Repository{}
	mockS3Repo.GetS3FileContentFunc = func(key string, bucket string) string {
		return `
monday:
 - start: 11:00
   end: 12:00
saturday:
 - start: 11:10
   end: 12:10
 - start: 13:10
   end: 14:10
Sunday:
 - start: 11:10
   end: 12:10
`
	}
	parser := InitTimesheetParser(mockS3Repo)
	s3Key := "anyS3Key"
	s3Bucket := "anyS3Bucket"
	_ = os.Setenv("TIMESHEET_S3_KEY", s3Key)
	_ = os.Setenv("TIMESHEET_S3_BUCKET", s3Bucket)

	currentDate := time.Now().Format(timesheetEntryDateLayout)

	verifyTimesheetEntry(t, parser, time.Monday, []TimesheetEntry{{currentDate, "11:00", "12:00"}})
	verifyS3Call(t, *mockS3Repo, s3Key, s3Bucket)
	verifyTimesheetEntry(t, parser, time.Saturday, []TimesheetEntry{{currentDate, "11:10", "12:10"}, {currentDate, "13:10", "14:10"}})
	verifyS3Call(t, *mockS3Repo, s3Key, s3Bucket)
	// Sunday is not written in correct format
	verifyTimesheetEntry(t, parser, time.Sunday, []TimesheetEntry{})
	verifyS3Call(t, *mockS3Repo, s3Key, s3Bucket)
	verifyTimesheetEntry(t, parser, time.Tuesday, []TimesheetEntry{})
	verifyS3Call(t, *mockS3Repo, s3Key, s3Bucket)
}
