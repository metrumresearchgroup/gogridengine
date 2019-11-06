package gogridengine

import (
	"encoding/xml"
)

//Host is the top-level object (per host) that includes all subsequent data including jobs, resources etc
type Host struct {
	XMLName       xml.Name     `xml:"Queue-List" json:"-"`
	Name          string       `xml:"name" json:"name"`
	QType         string       `xml:"qtype" json:"qtype"`
	SlotsUsed     int32        `xml:"slots_used" json:"slots_used"`
	SlotsReserved int32        `xml:"slots_rsv" json:"slots_reserved"`
	SlotsTotal    int32        `xml:"slots_total" json:"slots_total"`
	LoadAverage   float64      `xml:"load_avg" json:"load_average"`
	Resources     ResourceList `xml:"resource" json:"resources"`
	JobList       []Job        `xml:"job_list" json:"job_list"`
}
