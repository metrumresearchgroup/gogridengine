package filters

import (
	"strconv"
	"strings"
	"time"

	"github.com/metrumresearchgroup/gogridengine"
)

const (
	//ISO8601FMT is a constant format used for parsing ISO 8601 compliant datetimes
	ISO8601FMT string = "2006-01-02T15:04:05"
)

//NewJobOwnerFilter creates a filter function based on the provided inputs that only returns jobs with the specified owner / owners.
func NewJobOwnerFilter(inputs ...string) func(j gogridengine.Job) bool {
	return func(j gogridengine.Job) bool {

		for _, v := range inputs {
			if j.JobOwner == v {
				return true
			}
		}

		return false
	}
}

//NewJobStateFilter creates a filter function to be used based on the inputs provided. If more than one is provided, any match from that list will return a true for the job it's evaluating.
func NewJobStateFilter(inputs ...string) func(j gogridengine.Job) bool {
	return func(j gogridengine.Job) bool {

		for _, v := range inputs {
			if strings.Contains(j.State, v) {
				return true
			}
		}

		return false
	}
}

//NewStartingJobNumberFilter creates a filter function based on the inputs provided. Since we're filtering only on a starting point, even if more than one input is provided, only the first is evaluated.
func NewStartingJobNumberFilter(inputs ...string) func(j gogridengine.Job) bool {
	return func(j gogridengine.Job) bool {

		for _, v := range inputs {
			jobNumber, err := strconv.ParseInt(v, 10, 64)

			if err != nil {
				return false
			}

			//We'll only ever evaluate the first one.
			return j.JBJobNumber >= jobNumber
		}

		return false
	}
}

//NewBeforeSubmissionTimeFilter creates a filter function based on the provided inputs to determine if a job's submission time occurs before the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewBeforeSubmissionTimeFilter(inputs ...string) func(j gogridengine.Job) bool {
	return func(j gogridengine.Job) bool {

		for _, v := range inputs {
			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false
			}

			jobTime, err := time.Parse(ISO8601FMT, j.SubmittedTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false
			}

			return jobTime.Before(providedtime)

		}

		return false
	}
}

//NewAfterSubmissionTimeFilter creates a filter function based on the provided inputs to determine if a job's submission time occurs after the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewAfterSubmissionTimeFilter(inputs ...string) func(j gogridengine.Job) bool {
	return func(j gogridengine.Job) bool {

		for _, v := range inputs {
			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false
			}

			jobTime, err := time.Parse(ISO8601FMT, j.SubmittedTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false
			}

			return jobTime.After(providedtime)

		}

		return false
	}
}

//NewBeforeStartTimeFilter creates a filter function based on the provided inputs to determine if a job's start time occurs before the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewBeforeStartTimeFilter(inputs ...string) func(j gogridengine.Job) bool {
	return func(j gogridengine.Job) bool {

		for _, v := range inputs {
			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false
			}

			jobTime, err := time.Parse(ISO8601FMT, j.StartTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false
			}

			return jobTime.Before(providedtime)

		}

		return false
	}
}

//NewAfterStartTimeFilter creates a filter function based on the provided inputs to determine if a job's start time occurs after the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewAfterStartTimeFilter(inputs ...string) func(j gogridengine.Job) bool {
	return func(j gogridengine.Job) bool {

		for _, v := range inputs {
			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false
			}

			jobTime, err := time.Parse(ISO8601FMT, j.StartTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false
			}

			return jobTime.After(providedtime)

		}

		return false
	}
}
