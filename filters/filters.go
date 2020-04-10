package filters

import (
	"errors"
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
func NewJobOwnerFilter(inputs ...string) func(j gogridengine.Job) (bool, error) {
	return func(j gogridengine.Job) (bool, error) {

		if len(inputs) == 0 || inputs == nil {
			return false, errors.New("No inputs were provided upon which to operate")
		}

		for _, v := range inputs {

			if v == "" {
				return false, errors.New("One of the provided options was empty and is non-operable")
			}

			if j.JobOwner == v {
				return true, nil
			}
		}

		return false, nil
	}
}

//NewJobStateFilter creates a filter function to be used based on the inputs provided. If more than one is provided, any match from that list will return a true for the job it's evaluating.
func NewJobStateFilter(inputs ...string) func(j gogridengine.Job) (bool, error) {
	return func(j gogridengine.Job) (bool, error) {

		if len(inputs) == 0 || inputs == nil {
			return false, errors.New("No inputs were provided upon which to operate")
		}

		for _, v := range inputs {
			if v == "" {
				return false, errors.New("One of the provided options was empty and is non-operable")
			}
			if strings.Contains(j.State, v) {
				return true, nil
			}
		}

		return false, nil
	}
}

//NewStartingJobNumberFilter creates a filter function based on the inputs provided. Since we're filtering only on a starting point, if more than one are provided we generate an error.
func NewStartingJobNumberFilter(inputs ...string) func(j gogridengine.Job) (bool, error) {
	return func(j gogridengine.Job) (bool, error) {

		if len(inputs) == 0 || inputs == nil {
			return false, errors.New("No inputs were provided upon which to operate")
		}

		if len(inputs) > 1 {
			return false, errors.New("This filter only validly accepts a single input")
		}

		for _, v := range inputs {

			if v == "" {
				return false, errors.New("One of the provided options was empty and is non-operable")
			}

			jobNumber, err := strconv.ParseInt(v, 10, 64)

			if err != nil {
				return false, errors.New("The provided input could not be safely parsed into an integer")
			}

			//We'll only ever evaluate the first one.
			return j.JBJobNumber >= jobNumber, nil
		}

		return false, nil
	}
}

//NewBeforeSubmissionTimeFilter creates a filter function based on the provided inputs to determine if a job's submission time occurs before the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewBeforeSubmissionTimeFilter(inputs ...string) func(j gogridengine.Job) (bool, error) {
	return func(j gogridengine.Job) (bool, error) {

		if len(inputs) == 0 || inputs == nil {
			return false, errors.New("No inputs were provided upon which to operate")
		}

		if len(inputs) > 1 {
			return false, errors.New("Only one provided time is allowed on this method")
		}

		//Don't throw an error, but don't try matching if it's an empty string on the job
		if j.SubmittedTime == "" {
			return false, nil
		}

		for _, v := range inputs {

			if v == "" {
				return false, errors.New("One of the provided options was empty and is non-operable")
			}

			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false, errors.New("Unable to parse the provided input into RFC8601 formatted time")
			}

			jobTime, err := time.Parse(ISO8601FMT, j.SubmittedTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false, errors.New("Unable to parse the Submmission time on the job input into RFC8601 formatted time")
			}

			return jobTime.Before(providedtime), nil

		}

		return false, nil
	}
}

//NewAfterSubmissionTimeFilter creates a filter function based on the provided inputs to determine if a job's submission time occurs after the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewAfterSubmissionTimeFilter(inputs ...string) func(j gogridengine.Job) (bool, error) {
	return func(j gogridengine.Job) (bool, error) {

		if len(inputs) == 0 || inputs == nil {
			return false, errors.New("No inputs were provided upon which to operate")
		}

		if len(inputs) > 1 {
			return false, errors.New("Only one provided time is allowed on this method")
		}

		//Don't throw an error, but don't try matching if it's an empty string on the job
		if j.SubmittedTime == "" {
			return false, nil
		}

		for _, v := range inputs {

			if v == "" {
				return false, errors.New("One of the provided options was empty and is non-operable")
			}

			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false, errors.New("Unable to parse the provided input into RFC8601 formatted time")
			}

			jobTime, err := time.Parse(ISO8601FMT, j.SubmittedTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false, errors.New("Unable to parse the Submmission time on the job input into RFC8601 formatted time")
			}

			return jobTime.After(providedtime), nil

		}

		return false, nil
	}
}

//NewBeforeStartTimeFilter creates a filter function based on the provided inputs to determine if a job's start time occurs before the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewBeforeStartTimeFilter(inputs ...string) func(j gogridengine.Job) (bool, error) {
	return func(j gogridengine.Job) (bool, error) {

		if len(inputs) == 0 || inputs == nil {
			return false, errors.New("No inputs were provided upon which to operate")
		}

		if len(inputs) > 1 {
			return false, errors.New("Only one provided time is allowed on this method")
		}

		//Don't throw an error, but don't try matching if it's an empty string on the job
		if j.StartTime == "" {
			return false, nil
		}

		for _, v := range inputs {

			if v == "" {
				return false, errors.New("One of the provided options was empty and is non-operable")
			}

			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false, errors.New("Unable to parse the provided input into RFC8601 formatted time")
			}

			jobTime, err := time.Parse(ISO8601FMT, j.StartTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false, errors.New("Unable to format the start time of the job into RFC8601 formatted time")
			}

			return jobTime.Before(providedtime), nil

		}

		return false, nil
	}
}

//NewAfterStartTimeFilter creates a filter function based on the provided inputs to determine if a job's start time occurs after the provided time. Please note that even though multiple inputs are accepted, only the first is evaluated.
func NewAfterStartTimeFilter(inputs ...string) func(j gogridengine.Job) (bool, error) {
	return func(j gogridengine.Job) (bool, error) {

		if len(inputs) == 0 || inputs == nil {
			return false, errors.New("No inputs were provided upon which to operate")
		}

		if len(inputs) > 1 {
			return false, errors.New("This method only supports a single input")
		}

		//Just don't process if not present
		if j.StartTime == "" {
			return false, nil
		}

		for _, v := range inputs {

			if v == "" {
				return false, errors.New("One of the provided options was empty and is non-operable")
			}

			providedtime, err := time.Parse(ISO8601FMT, v)

			//If we can't parse it. It fails
			if err != nil {
				return false, errors.New("The provided time could not be parsed into a valid RFC8601 Time")
			}

			jobTime, err := time.Parse(ISO8601FMT, j.StartTime)

			//If we can't parse it from the job. It is teh gone.
			if err != nil {
				return false, errors.New("The start time of the job could not be parsed into a valid RFC8601 Time")
			}

			return jobTime.After(providedtime), nil

		}

		return false, nil
	}
}
