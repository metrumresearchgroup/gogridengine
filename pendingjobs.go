package gogridengine

import (
	"encoding/xml"
)

// PendingJobs is a sub tag of job_info (also labeled job_info) which details jobs not yet executing.
type PendingJobs struct {
	XMLName xml.Name `xml:"job_info" json:"-"`
	JobList []Job    `xml:"job_list" json:"job_list"`
}
