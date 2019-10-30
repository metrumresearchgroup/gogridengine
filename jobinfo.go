package gogridengine

import (
	"encoding/xml"
)

const (
	header = "<?xml version='1.0'?>"
)

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

//GetJobInfo provides a way of requesting the marshalled content directly without having to intervene
func GetJobInfo() (JobInfo, error) {
	content, err := GetQstatOutput()

	if err != nil {
		return JobInfo{}, err
	}

	var ji JobInfo

	err = xml.Unmarshal([]byte(content), &ji)

	if err != nil {
		return JobInfo{}, err
	}

	return ji, nil
}
