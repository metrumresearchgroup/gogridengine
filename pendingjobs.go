package gogridengine

import (
	"encoding/xml"
)

//PendingJob is a sub tag of job_info (also labeled job_info) which details jobs not yet executing.
type PendingJob struct {
	XMLName xml.Name  `xml:"job_info" json:"job_info"`
	JobList []JobList `xml:"job_list" json:"job_list"`
}
