package sge

import (
	"encoding/xml"
	"fmt"
	"github.com/metrumresearchgroup/gogridengine"
)

type GridExtractor interface {
	Extract() ([]byte, error)
	AddArguments(args ...string)
}

type JobRepository struct {
	extractor GridExtractor
}

func New(extractor GridExtractor) *JobRepository {
	return &JobRepository{
		extractor: extractor,
	}
}

type StateCode string

const (
	StatusPending  StateCode = "p"
	StatusRunning  StateCode = "r"
	StatusStopped  StateCode = "s"
	StatusComplete StateCode = "z"
)

type GetJobsRequest struct {
	User  *string
	State *StateCode
	ID    *int    //Need to figure out how this would work. Our qstat bin doesn't seem to take a jobID. Could always filter the top level list
	Host  *string // uses the -q option
}

// Bootstrap is responsible for generating additional parameters passed down to the underlying cmd
func (request *GetJobsRequest) Bootstrap(extractor GridExtractor) {
	if request.User != nil {
		extractor.AddArguments([]string{"-u", *request.User}...)
	}

	if request.State != nil {
		extractor.AddArguments([]string{"-s", string(*request.State)}...)
	}

	if request.Host != nil {
		extractor.AddArguments([]string{"-q", *request.Host}...)
	}
}

func (r *JobRepository) Get(request *GetJobsRequest) (gogridengine.JobList, error) {
	var outputJobs gogridengine.JobList

	request.Bootstrap(r.extractor)

	outputBytes, err := r.extractor.Extract()
	if err != nil {
		return nil, err
	}

	var ji gogridengine.JobInfo
	if err = xml.Unmarshal(outputBytes, &ji); err != nil {
		return nil, fmt.Errorf("unable to process expected xml output from command: %w", err)
	}

	//Pending jobs
	outputJobs = append(outputJobs, ji.PendingJobs.JobList...)

	for _, v := range ji.QueueInfo.Queues {
		outputJobs = append(outputJobs, v.JobList...)
	}

	return outputJobs, nil
}
