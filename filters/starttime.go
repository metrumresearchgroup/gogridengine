package filters

import (
	"time"

	"github.com/metrumresearchgroup/gogridengine"
	log "github.com/sirupsen/logrus"
)

//NewBeforeStartTimeFilter returns only jobs whose start time occurs before the provided time.
func NewBeforeStartTimeFilter(t time.Time) func(job gogridengine.Job) bool {
	return func(job gogridengine.Job) bool {
		jobTime, err := time.Parse(ISO8601FMT, job.StartTime)
		if err != nil {
			//If we can't parse the value, discard the job
			log.Error("Failed parsing the time content: ", err)
			return false
		}

		return jobTime.Before(t)
	}
}

//NewAfterStartTimeFilter returns only jobs whose start time occurs after the provided time.
func NewAfterStartTimeFilter(t time.Time) func(job gogridengine.Job) bool {
	return func(job gogridengine.Job) bool {
		jobTime, err := time.Parse(ISO8601FMT, job.StartTime)
		if err != nil {
			//If we can't parse the value, discard the job
			return false
		}

		return jobTime.After(t)
	}
}

//NewBetweenStartTimeFilter allows you to provide a start and end time to return jobs whos start time falls within that range
func NewBetweenStartTimeFilter(start time.Time, end time.Time) func(job gogridengine.Job) bool {
	return func(job gogridengine.Job) bool {
		jobTime, err := time.Parse(ISO8601FMT, job.StartTime)
		if err != nil {
			//If we can't parse the value, discard the job
			return false
		}

		return jobTime.After(start) && jobTime.Before(end)
	}
}
