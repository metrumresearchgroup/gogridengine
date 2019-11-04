package filters

import (
	"strings"

	"github.com/metrumresearchgroup/gogridengine"
)

const (
	//ISO8601FMT is a constant format used for parsing ISO 8601 compliant datetimes
	ISO8601FMT string = "2006-01-02T15:04:05"
)

//NewUsernameFilter returns a filter function for specifying an owner to filter a JobList Down
func NewUsernameFilter(username string) func(job gogridengine.Job) bool {
	return func(job gogridengine.Job) bool {
		return job.JobOwner == username
	}
}

//NewLooseStateFilter is returns a filter function for specifying a loose match on a state code. Any state code containing the code provided will be returned
func NewLooseStateFilter(state string) func(job gogridengine.Job) bool {
	return func(job gogridengine.Job) bool {
		return strings.Contains(job.State, state)
	}
}

//NewStrictStateFilter returns only jobs that match the state code you provide exactly
func NewStrictStateFilter(state string) func(job gogridengine.Job) bool {
	return func(job gogridengine.Job) bool {
		return job.State == state
	}
}
