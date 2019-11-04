package gogridengine

import (
	"encoding/xml"
)

//JobList is a slice of Jobs that is filterable and otherwise actionable via receiver.
type JobList []Job

//Job is the Sun Grid Engine XML Definition for a job running on a specific host, its details and current status
type Job struct {
	//Because this is a node, we still need the XMLName identifier
	XMLName        xml.Name `xml:"job_list" json:"-"`
	StateAttribute string   `xml:"state,attr" json:"state_attribute_text"`
	State          string   `xml:"state" json:"state"`
	JBJobNumber    int64    `xml:"JB_job_number" json:"jb_job_number"`
	JATPriority    float64  `xml:"JAT_prio" json:"jat_prio"`
	JobName        string   `xml:"JB_name" json:"jb_name"`
	JobOwner       string   `xml:"JB_owner" json:"jb_owner"`
	StartTime      string   `xml:"JAT_start_time,omitempty" json:"start_time"`
	SubmittedTime  string   `xml:"JB_submission_time,omitempty" json:"submitted_time"`
	Slots          int32    `xml:"slots" json:"slots"`
}

//IsJobRunning returns a int (1 - running) (0 - not)
func IsJobRunning(job Job) int {

	if job.State == "r" {
		return 1
	}

	return 0
}

//GetJobs returns a slice of only jobs from both scheduled and unscheduled queues
func GetJobs() (JobList, error) {
	var jobs []Job

	ji, err := GetJobInfo()

	if err != nil {
		return []Job{}, err
	}

	//Add running jobs to the slice first
	for _, q := range ji.QueueInfo.Queues {
		jobs = append(jobs, q.JobList...)
	}

	//Add pending jobs
	jobs = append(jobs, ji.PendingJobs.JobList...)

	return jobs, nil
}

//GetJobsWithFilter allows you to specify a filter at the time of retrieving the JobList
func GetJobsWithFilter(filterfunc func(j Job) bool) (JobList, error) {
	jobs, err := GetJobs()
	if err != nil {
		return JobList{}, err
	}

	return jobs.Filter(filterfunc), nil
}

//Filter allows for the passage of any function taking a JobList and Filtering its contents down.
//Should be usable in fluent fashion as long as JobList is being returned
func (jl JobList) Filter(filter func(j Job) bool) JobList {
	var jobs JobList

	for _, v := range jl {
		if filter(v) {
			jobs = append(jobs, v)
		}
	}

	return jobs
}
