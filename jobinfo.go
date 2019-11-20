package gogridengine

import (
	"encoding/xml"
	"regexp"
	"sort"

	log "github.com/sirupsen/logrus"
)

const (
	header = "<?xml version='1.0'?>"
)

const (
	environmentPrefix string = "GOGRIDENGINE_"
)

//TaskRangeRegex is the compiled regular expression used for identifying Tasks that define a pending / unscheduled range
var TaskRangeRegex *regexp.Regexp = regexp.MustCompile(TASKRANGEIDENTIFIERREGEX)

//JobInfo is the top level object for the SGE Qstat output
type JobInfo struct {
	XMLName     xml.Name   `xml:"job_info" json:"-"`
	QueueInfo   QueueInfo  `xml:"queue_info" json:"queue_info"`
	PendingJobs PendingJob `xml:"job_info,omitempty" json:"pending_jobs"`
}

//GetXML renders down the XML with UTF-8 opening tags to ensure feasability for testing of output
func (q JobInfo) GetXML() (string, error) {
	output, err := xml.Marshal(q)

	if err != nil {
		return "", err
	}

	formatted := header + string(output)

	return formatted, nil
}

//NewJobInfo returns the go struct of the qstat output
func NewJobInfo(input string) (JobInfo, error) {
	var ji JobInfo
	err := xml.Unmarshal([]byte(input), &ji)

	if err != nil {
		return JobInfo{}, err
	}

	deleteTargets := make(map[int]Job)

	//Handle extrapolation of pending tasks

	for k, p := range ji.PendingJobs.JobList {
		if DoesJobContainTaskRange(p) {
			//Mark for deletion and substitution
			deleteTargets[k] = p
		}
	}

	//If anything comes up as an extrapolatable task list, we need to extrapolate to multiple job entries and remove the original listing.
	if len(deleteTargets) > 0 {
		//Reiterate over collected jobs to cleanup and reconstruct
		for k, p := range deleteTargets {

			jobs, err := ExtrapolateTasksToJobs(p)

			if err != nil {
				//We can't do anything with this entry. Just continue along
				log.Error("An error occurred trying to extrapolate Task range into JobList", err)
				continue
			}

			//Remove the target Job
			ji.PendingJobs.JobList = append(ji.PendingJobs.JobList[:k], ji.PendingJobs.JobList[k+1:]...)

			//Append Extrapolated Jobs
			ji.PendingJobs.JobList = append(ji.PendingJobs.JobList, jobs...)
		}

		//Sort the slice after all the shuffling By Job Number and Task IDß
		sort.Slice(ji.PendingJobs.JobList, func(i, j int) bool {
			return ji.PendingJobs.JobList[i].JBJobNumber < ji.PendingJobs.JobList[j].JBJobNumber && ji.PendingJobs.JobList[i].Tasks.TaskID < ji.PendingJobs.JobList[j].Tasks.TaskID
		})
	}

	return ji, nil
}
