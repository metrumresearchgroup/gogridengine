package gogridengine

import ()
import "encoding/xml"

//JobList is the Sun Grid Engine XML Definition for a job running on a specific host, its details and current status
type JobList struct {
	//Because this is a node, we still need the XMLName identifier
	XMLName     xml.Name `xml:"job_list" json:"-"`
	State       string   `xml:"state,attr" json:"state"`
	JBJobNumber int64    `xml:"JB_job_number" json:"jb_job_number"`
	JATPriority float64  `xml:"JAT_prio" json:"jat_prio"`
	JobName     string   `xml:"JB_name" json:"jb_name"`
	JobOwner    string   `xml:"JB_owner" json:"jb_owner"`
	StartTime   string   `xml:"JAT_start_time" json:"start_time"`
	Slots       int32    `xml:"slots" json:"slots"`
}
