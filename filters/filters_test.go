package filters

import (
	"testing"
	"time"

	"github.com/metrumresearchgroup/gogridengine"
	"github.com/stretchr/testify/assert"
)

func TestNewUsernameFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			JobOwner: "Bob",
		},
		{
			JobOwner: "Cindy",
		},
	}

	jl = jl.Filter(NewUsernameFilter("Cindy"))
	assert.Equal(t, 1, len(jl))

	jl = jl.Filter(NewUsernameFilter("bobby"))

	assert.Empty(t, jl)
}

func TestNewLooseStateFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			State: "r",
		},
		{
			State: "e",
		},
		{
			State: "ew",
		},
		{
			State: "qw",
		},
		{
			State: "ce",
		},
	}

	//Test the loose filter first
	r1 := jl.Filter(NewLooseStateFilter("r"))

	assert.NotEmpty(t, r1)
	assert.Len(t, r1, 1)

	r2 := jl.Filter(NewLooseStateFilter("w"))

	assert.NotEmpty(t, r2)
	assert.Len(t, r2, 2)

	r3 := jl.Filter(NewLooseStateFilter("e"))

	assert.NotEmpty(t, r3)
	assert.Len(t, r3, 3)

	//Test for Chained Loose Filter
	r4 := jl.
		Filter(NewLooseStateFilter("e")).
		Filter(NewLooseStateFilter("w"))

	assert.NotEmpty(t, r4)
	assert.Len(t, r4, 1)
	assert.Equal(t, r4[0].State, "ew")
}

func TestNewStrictStateFilter(t *testing.T) {

	jl := gogridengine.JobList{
		{
			State: "r",
		},
		{
			State: "e",
		},
		{
			State: "ew",
		},
		{
			State: "qw",
		},
		{
			State: "ce",
		},
		{
			State: "r",
		},
	}

	r1 := jl.Filter(NewStrictStateFilter("r"))

	assert.NotEmpty(t, r1)
	assert.Len(t, r1, 2)

	r2 := jl.
		Filter(NewLooseStateFilter("e")).
		Filter(NewStrictStateFilter("ew"))

	assert.NotEmpty(t, r2)
	assert.Len(t, r2, 1)

}

func TestNewBeforeSubmitTimeFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			JobName:       "TheRightOne",
			SubmittedTime: "2019-09-15T15:26:36",
		},
		{
			JobName:       "TheWrongOne",
			SubmittedTime: "2019-09-21T15:26:36",
		},
		{
			JobName:       "Invalid",
			SubmittedTime: "NotEvenAValidTime",
		},
	}

	//Two days in the future
	target := "2019-09-17T15:26:36"
	targetTime, _ := time.Parse(ISO8601FMT, target)

	//Show me jobs with a submit time earlier than the targetTime.
	jl = jl.Filter(NewBeforeSubmitTimeFilter(targetTime))

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 1)
	assert.Equal(t, "TheRightOne", jl[0].JobName)
}

func TestNewAfterSubmitTimeFilter(t *testing.T) {

	jl := gogridengine.JobList{
		{
			JobName:       "TheWrongOne",
			SubmittedTime: "2019-09-15T15:26:36",
		},
		{
			JobName:       "TheRightOne",
			SubmittedTime: "2019-09-21T15:26:36",
		},
		{
			JobName:       "TheRightOne",
			SubmittedTime: "2019-09-21T15:26:36",
		},
		{
			JobName:       "TheRightOne",
			SubmittedTime: "2019-09-21T15:26:37",
		},
		{
			JobName:       "Invalid",
			SubmittedTime: "ImNotEvenAValidTime",
		},
	}

	//Two days in the future
	target := "2019-09-17T15:26:36"
	targetTime, _ := time.Parse(ISO8601FMT, target)

	//Show me jobs with a submit time earlier than the targetTime.
	jl = jl.Filter(NewAfterSubmitTimeFilter(targetTime))

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 3)
	assert.Equal(t, "TheRightOne", jl[0].JobName)
}

func TestNewSubmitTimeBetweenFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			JobName:       "TheWrongOne",
			SubmittedTime: "2019-09-15T15:26:36",
		},
		{
			JobName:       "TheRightOne",
			SubmittedTime: "2019-09-21T15:26:36",
		},
		{
			JobName:       "SecondRightOne",
			SubmittedTime: "2019-09-21T15:26:36",
		},
		{
			JobName:       "OutofSpec",
			SubmittedTime: "2019-09-21T15:26:37",
		},
		{
			JobName:       "Invalid",
			SubmittedTime: "ImNotEvenAValidTime",
		},
	}

	//Two days in the future
	start := "2019-09-21T15:26:35"
	startTime, _ := time.Parse(ISO8601FMT, start)

	end := "2019-09-21T15:26:37"
	endTime, _ := time.Parse(ISO8601FMT, end)

	//Show me jobs with a submit time earlier than the targetTime.
	jl = jl.Filter(NewBetweenSubmitTimeFilter(startTime, endTime))

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 2)
	assert.Equal(t, "TheRightOne", jl[0].JobName)

}

func TestNewBeforeStartTimeFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			JobName:   "TheRightOne",
			StartTime: "2019-09-15T15:26:36",
		},
		{
			JobName:   "TheWrongOne",
			StartTime: "2019-09-21T15:26:36",
		},
		{
			JobName:   "Invalid",
			StartTime: "NotEvenAValidTime",
		},
	}

	//Two days in the future
	target := "2019-09-17T15:26:36"
	targetTime, _ := time.Parse(ISO8601FMT, target)

	//Show me jobs with a submit time earlier than the targetTime.
	jl = jl.Filter(NewBeforeStartTimeFilter(targetTime))

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 1)
	assert.Equal(t, "TheRightOne", jl[0].JobName)
}

func TestNewAfterStartTimeFilter(t *testing.T) {

	jl := gogridengine.JobList{
		{
			JobName:   "TheWrongOne",
			StartTime: "2019-09-15T15:26:36",
		},
		{
			JobName:   "TheRightOne",
			StartTime: "2019-09-21T15:26:36",
		},
		{
			JobName:   "TheRightOne",
			StartTime: "2019-09-21T15:26:36",
		},
		{
			JobName:   "TheRightOne",
			StartTime: "2019-09-21T15:26:37",
		},
		{
			JobName:   "Invalid",
			StartTime: "ImNotEvenAValidTime",
		},
	}

	//Two days in the future
	target := "2019-09-17T15:26:36"
	targetTime, _ := time.Parse(ISO8601FMT, target)

	//Show me jobs with a submit time earlier than the targetTime.
	jl = jl.Filter(NewAfterStartTimeFilter(targetTime))

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 3)
	assert.Equal(t, "TheRightOne", jl[0].JobName)
}

func TestNewStartTimeBetweenFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			JobName:   "TheWrongOne",
			StartTime: "2019-09-15T15:26:36",
		},
		{
			JobName:   "TheRightOne",
			StartTime: "2019-09-21T15:26:36",
		},
		{
			JobName:   "SecondRightOne",
			StartTime: "2019-09-21T15:26:36",
		},
		{
			JobName:   "OutofSpec",
			StartTime: "2019-09-21T15:26:37",
		},
		{
			JobName:   "Invalid",
			StartTime: "ImNotEvenAValidTime",
		},
	}

	//Two days in the future
	start := "2019-09-21T15:26:35"
	startTime, _ := time.Parse(ISO8601FMT, start)

	end := "2019-09-21T15:26:37"
	endTime, _ := time.Parse(ISO8601FMT, end)

	//Show me jobs with a submit time earlier than the targetTime.
	jl = jl.Filter(NewBetweenStartTimeFilter(startTime, endTime))

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 2)
	assert.Equal(t, "TheRightOne", jl[0].JobName)

}
