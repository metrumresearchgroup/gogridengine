package gogridengine

import (
	"encoding/xml"
)

//QueueInfo is the child object for qstat job output
type QueueInfo struct {
	XMLName xml.Name    `xml:"queue_info" json:"queue_info"`
	Queues  []QueueList `xml:"Queue-List" json:"queue_list"`
}
