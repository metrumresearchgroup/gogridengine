package filters

import (
	"sort"
	"testing"

	"github.com/metrumresearchgroup/gogridengine"
	"github.com/stretchr/testify/assert"
)

func TestNewJobOwnerFilter(t *testing.T) {
	jl := newJobList()

	jl, err := gogridengine.FilterJobs(jl, NewJobOwnerFilter("janed", "jilld"))

	assert.Nil(t, err)

	sort.Slice(jl, func(i, j int) bool {
		return jl[i].JBJobNumber < jl[j].JBJobNumber
	})

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 3)

	assert.Equal(t, int64(13), jl[0].JBJobNumber)
	assert.Equal(t, int64(14), jl[1].JBJobNumber)
	assert.Equal(t, int64(16), jl[2].JBJobNumber)

	jl = newJobList()

	jl, err = gogridengine.FilterJobs(jl, NewJobOwnerFilter(""))
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(jl, NewJobOwnerFilter())
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func newJobList() gogridengine.JobList {
	return gogridengine.JobList{
		{
			JBJobNumber:   11,
			JobOwner:      "johnd",
			State:         "r",
			StartTime:     "2019-01-13T11:21:14",
			SubmittedTime: "2018-01-13T11:21:15",
		},
		{
			JBJobNumber:   12,
			JobOwner:      "johnd",
			State:         "r",
			StartTime:     "2019-01-14T11:21:15",
			SubmittedTime: "2018-01-14T11:21:15",
		},
		{
			JBJobNumber:   13,
			JobOwner:      "jilld",
			State:         "eh",
			StartTime:     "2019-01-14T11:21:17",
			SubmittedTime: "2018-01-14T11:21:17",
		},
		{
			JBJobNumber:   14,
			JobOwner:      "janed",
			State:         "h",
			StartTime:     "2019-01-14T11:34:15",
			SubmittedTime: "2018-01-14T11:34:15",
		},
		{
			JBJobNumber:   15,
			JobOwner:      "joed",
			State:         "qw",
			StartTime:     "2019-01-15T08:34:15",
			SubmittedTime: "2018-01-15T08:34:15",
		},
		{
			JBJobNumber:   16,
			JobOwner:      "janed",
			State:         "h",
			StartTime:     "2019-01-15T23:34:15",
			SubmittedTime: "2018-01-15T23:34:15",
		},
	}
}

//10,13,14,16
func TestNewJobStateFilter(t *testing.T) {
	jl := newJobList()

	jl, err := gogridengine.FilterJobs(jl, NewJobStateFilter("e", "h"))

	assert.Nil(t, err)

	sort.Slice(jl, func(i, j int) bool {
		return jl[i].JBJobNumber < jl[j].JBJobNumber
	})

	assert.NotEmpty(t, jl)

	assert.Len(t, jl, 3)

	assert.Equal(t, int64(13), jl[0].JBJobNumber)
	assert.Equal(t, int64(14), jl[1].JBJobNumber)
	assert.Equal(t, int64(16), jl[2].JBJobNumber)

	jl = newJobList()
	jl, err = gogridengine.FilterJobs(jl, NewJobStateFilter(""))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(jl, NewJobStateFilter())

	assert.NotNil(t, err)
	assert.Error(t, err)

	assert.Equal(t, newJobList(), jl)
}

func TestNewStartingJobNumberFilter(t *testing.T) {
	jl := newJobList()

	jl, err := gogridengine.FilterJobs(jl, NewStartingJobNumberFilter("13"))

	assert.Nil(t, err)

	sort.Slice(jl, func(i, j int) bool {
		return jl[i].JBJobNumber < jl[j].JBJobNumber
	})

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 4)

	for i := 0; i < len(jl); i++ {
		assert.Equal(t, int64(13+i), jl[i].JBJobNumber)
	}

	jl, err = gogridengine.FilterJobs(jl, NewStartingJobNumberFilter("cat"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewStartingJobNumberFilter())

	//No input means every job should validate false.
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(jl, NewStartingJobNumberFilter(""))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(jl, NewStartingJobNumberFilter("13", "15"))

	assert.NotNil(t, err)
	assert.Error(t, err)
}

//11
func TestNewBeforeSubmissionTimeFilter(t *testing.T) {
	jl := newJobList()

	jl, err := gogridengine.FilterJobs(jl, NewBeforeSubmissionTimeFilter("2018-01-14T11:21:15"))

	assert.Nil(t, err)

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 1)
	assert.Equal(t, int64(11), jl[0].JBJobNumber)

	jl, err = gogridengine.FilterJobs(newJobList(), NewBeforeSubmissionTimeFilter("notavalidtime"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl = newJobList()
	jl[0].SubmittedTime = "meow"

	jl, err = gogridengine.FilterJobs(jl, NewBeforeSubmissionTimeFilter("2018-01-14T11:21:15"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	//Testing for basically empty submission time
	jl = newJobList()
	jl[0].SubmittedTime = ""

	jl, err = gogridengine.FilterJobs(jl, NewBeforeSubmissionTimeFilter("2018-01-14T11:21:15"))

	assert.Nil(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewBeforeSubmissionTimeFilter())
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewBeforeSubmissionTimeFilter(""))
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewBeforeSubmissionTimeFilter("2018-01-14T11:21:15", "2018-01-14T11:21:17"))
	assert.NotNil(t, err)
	assert.Error(t, err)
}

//13,14,15,16
func TestNewAfterSubmissionTimeFilter(t *testing.T) {
	jl := newJobList()

	jl, err := gogridengine.FilterJobs(jl, NewAfterSubmissionTimeFilter("2018-01-14T11:21:15"))

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.Nil(t, err)
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 4)

	for i := 0; i < len(jl); i++ {
		assert.Equal(t, int64(i+13), jl[i].JBJobNumber)
	}

	//Item without a time
	jl = newJobList()
	jl[0].SubmittedTime = ""
	jl, err = gogridengine.FilterJobs(jl, NewAfterSubmissionTimeFilter("2018-01-14T11:21:15"))

	assert.Nil(t, err)

	//Item with invalid time
	jl = newJobList()
	jl[0].SubmittedTime = "cat"
	jl, err = gogridengine.FilterJobs(jl, NewAfterSubmissionTimeFilter("2018-01-14T11:21:15"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewAfterSubmissionTimeFilter("notavalidtime"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewAfterSubmissionTimeFilter())
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewAfterSubmissionTimeFilter(""))
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(jl, NewAfterSubmissionTimeFilter("2018-01-14T11:21:15", "2018-01-14T11:21:17"))
	assert.NotNil(t, err)
	assert.Error(t, err)
}

//11
func TestNewBeforeStartTimeFilter(t *testing.T) {
	jl := newJobList()

	jl, err := gogridengine.FilterJobs(jl, NewBeforeStartTimeFilter("2019-01-14T11:21:15"))

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.Nil(t, err)
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 1)
	assert.Equal(t, int64(11), jl[0].JBJobNumber)

	jl = newJobList()
	jl[1].StartTime = ""
	jl, err = gogridengine.FilterJobs(jl, NewBeforeStartTimeFilter("2019-01-14T11:21:15"))

	assert.Nil(t, err)
	assert.NotEmpty(t, jl)

	jl = newJobList()
	jl[1].StartTime = "cat"
	jl, err = gogridengine.FilterJobs(jl, NewBeforeStartTimeFilter("2019-01-14T11:21:15"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewBeforeStartTimeFilter("notavalidtime"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewBeforeStartTimeFilter())
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewBeforeStartTimeFilter(""))
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(jl, NewBeforeStartTimeFilter("2019-01-14T11:21:15", "2019-01-14T11:21:15"))
	assert.NotNil(t, err)
	assert.Error(t, err)
}

//13,14,15,16
func TestNewAfterStartTimeFilter(t *testing.T) {
	jl := newJobList()

	jl, err := gogridengine.FilterJobs(jl, NewAfterStartTimeFilter("2019-01-14T11:21:15"))

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.Nil(t, err)
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 4)

	for i := 0; i < len(jl); i++ {
		assert.Equal(t, int64(i+13), jl[i].JBJobNumber)
	}

	jl = newJobList()
	jl[1].StartTime = ""

	jl, err = gogridengine.FilterJobs(jl, NewAfterStartTimeFilter("2019-01-14T11:21:15"))

	assert.Nil(t, err)
	assert.NotEmpty(t, jl)

	jl = newJobList()
	jl[1].StartTime = "cat"

	jl, err = gogridengine.FilterJobs(jl, NewAfterStartTimeFilter("2019-01-14T11:21:15"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewAfterStartTimeFilter("notavalidtime"))

	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewAfterStartTimeFilter())
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(newJobList(), NewAfterStartTimeFilter(""))
	assert.NotNil(t, err)
	assert.Error(t, err)

	jl, err = gogridengine.FilterJobs(jl, NewAfterStartTimeFilter("2019-01-14T11:21:15", "2019-01-14T11:21:17"))
	assert.NotNil(t, err)
	assert.Error(t, err)

}
