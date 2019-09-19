package gogridengine

import (
	"encoding/xml"
)

//JobInfo is the top level object for the SGE Qstat output
type JobInfo struct {
	XMLName   xml.Name  `xml:"job_info" json:"job_info"`
	QueueInfo QueueInfo `xml:"queue_info" json:"queue_info"`
}
