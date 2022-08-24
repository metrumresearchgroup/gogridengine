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

// TaskRangeRegex is the compiled regular expression used for identifying Tasks that define a pending / unscheduled range
var TaskRangeRegex *regexp.Regexp = regexp.MustCompile(TASKRANGEIDENTIFIERREGEX)

// JobInfo is the top level object for the SGE Qstat output
type JobInfo struct {
	XMLName   xml.Name  `xml:"job_info" json:"-"`
	QueueInfo QueueInfo `xml:"queue_info" json:"queue_info"`
	Jobs      Jobs      `xml:"job_info,omitempty" json:"pending_jobs"`
}

// GetXML renders down the XML with UTF-8 opening tags to ensure feasability for testing of output
func (q JobInfo) GetXML() (string, error) {
	output, err := xml.Marshal(q)

	if err != nil {
		return "", err
	}

	formatted := header + string(output)

	return formatted, nil
}

// NewJobInfo returns the go struct of the qstat output
func NewJobInfo(input string) (JobInfo, error) {
	var ji JobInfo
	err := xml.Unmarshal([]byte(input), &ji)

	if err != nil {
		return JobInfo{}, err
	}

	deleteTargets := make(map[int]Job)

	//Handle extrapolation of pending tasks

	for k, p := range ji.Jobs.JobList {
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
			ji.Jobs.JobList = append(ji.Jobs.JobList[:k], ji.Jobs.JobList[k+1:]...)

			//Append Extrapolated Jobs
			ji.Jobs.JobList = append(ji.Jobs.JobList, jobs...)
		}

		//Sort the slice after all the shuffling By Job Number and Task IDß
		sort.Slice(ji.Jobs.JobList, func(i, j int) bool {
			return ji.Jobs.JobList[i].JBJobNumber < ji.Jobs.JobList[j].JBJobNumber && ji.Jobs.JobList[i].Tasks.TaskID < ji.Jobs.JobList[j].Tasks.TaskID
		})
	}

	return ji, nil
}
