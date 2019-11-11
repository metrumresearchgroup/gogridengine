package filters

import (
	"sort"
	"testing"

	"github.com/metrumresearchgroup/gogridengine"
	"github.com/stretchr/testify/assert"
)

func TestNewJobOwnerFilter(t *testing.T) {
	jl := newJobList()

	jl = gogridengine.FilterJobs(jl, NewJobOwnerFilter("janed", "jilld"))

	sort.Slice(jl, func(i, j int) bool {
		return jl[i].JBJobNumber < jl[j].JBJobNumber
	})

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 4)

	assert.Equal(t, int64(10), jl[0].JBJobNumber)
	assert.Equal(t, int64(13), jl[1].JBJobNumber)
	assert.Equal(t, int64(14), jl[2].JBJobNumber)
	assert.Equal(t, int64(16), jl[3].JBJobNumber)
}

func newJobList() gogridengine.JobList {
	return gogridengine.JobList{
		{
			JBJobNumber:   10,
			JobOwner:      "janed",
			State:         "h",
			StartTime:     "notavalidtime",
			SubmittedTime: "notavalidtime",
		},
		{
			JBJobNumber:   11,
			JobOwner:      "johnd",
			State:         "r",
			StartTime:     "2019-01-13T11:21:15",
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

	jl = gogridengine.FilterJobs(jl, NewJobStateFilter("e", "h"))

	sort.Slice(jl, func(i, j int) bool {
		return jl[i].JBJobNumber < jl[j].JBJobNumber
	})

	assert.NotEmpty(t, jl)

	assert.Len(t, jl, 4)

	assert.Equal(t, int64(10), jl[0].JBJobNumber)
	assert.Equal(t, int64(13), jl[1].JBJobNumber)
	assert.Equal(t, int64(14), jl[2].JBJobNumber)
	assert.Equal(t, int64(16), jl[3].JBJobNumber)
}

func TestNewStartingJobNumberFilter(t *testing.T) {
	jl := newJobList()

	jl = gogridengine.FilterJobs(jl, NewStartingJobNumberFilter("13"))

	sort.Slice(jl, func(i, j int) bool {
		return jl[i].JBJobNumber < jl[j].JBJobNumber
	})

	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 4)

	for i := 0; i < len(jl); i++ {
		assert.Equal(t, int64(13+i), jl[i].JBJobNumber)
	}

	jl = gogridengine.FilterJobs(jl, NewStartingJobNumberFilter("cat"))

	assert.Empty(t, jl)

	jl = gogridengine.FilterJobs(newJobList(), NewStartingJobNumberFilter())

	//No input means every job should validate false.
	assert.Empty(t, jl)
}

//11
func TestNewBeforeSubmissionTimeFilter(t *testing.T) {
	jl := newJobList()

	jl = gogridengine.FilterJobs(jl, NewBeforeSubmissionTimeFilter("2018-01-14T11:21:15", "2025-01-14T11:21:15"))

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 1)
	assert.Equal(t, int64(11), jl[0].JBJobNumber)

	jl = gogridengine.FilterJobs(newJobList(), NewBeforeSubmissionTimeFilter("notavalidtime"))

	assert.Empty(t, jl)

	jl = gogridengine.FilterJobs(newJobList(), NewBeforeSubmissionTimeFilter())
	assert.Empty(t, jl)
}

//13,14,15,16
func TestNewAfterSubmissionTimeFilter(t *testing.T) {
	jl := newJobList()

	jl = gogridengine.FilterJobs(jl, NewAfterSubmissionTimeFilter("2018-01-14T11:21:15", "2025-01-14T11:21:15"))

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 4)

	for i := 0; i < len(jl); i++ {
		assert.Equal(t, int64(i+13), jl[i].JBJobNumber)
	}

	jl = gogridengine.FilterJobs(newJobList(), NewAfterSubmissionTimeFilter("notavalidtime"))

	assert.Empty(t, jl)

	jl = gogridengine.FilterJobs(newJobList(), NewAfterSubmissionTimeFilter())
	assert.Empty(t, jl)
}

//11
func TestNewBeforeStartTimeFilter(t *testing.T) {
	jl := newJobList()

	jl = gogridengine.FilterJobs(jl, NewBeforeStartTimeFilter("2019-01-14T11:21:15", "2025-01-14T11:21:15"))

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 1)
	assert.Equal(t, int64(11), jl[0].JBJobNumber)

	jl = gogridengine.FilterJobs(newJobList(), NewBeforeStartTimeFilter("notavalidtime"))

	assert.Empty(t, jl)

	jl = gogridengine.FilterJobs(newJobList(), NewBeforeStartTimeFilter())
	assert.Empty(t, jl)
}

//13,14,15,16
func TestNewAfterStartTimeFilter(t *testing.T) {
	jl := newJobList()

	jl = gogridengine.FilterJobs(jl, NewAfterStartTimeFilter("2019-01-14T11:21:15", "2025-01-14T11:21:15"))

	//Remember that Job 1 can't be processed so isn't eligible for the filter.
	assert.NotEmpty(t, jl)
	assert.Len(t, jl, 4)

	for i := 0; i < len(jl); i++ {
		assert.Equal(t, int64(i+13), jl[i].JBJobNumber)
	}

	jl = gogridengine.FilterJobs(newJobList(), NewAfterStartTimeFilter("notavalidtime"))

	assert.Empty(t, jl)

	jl = gogridengine.FilterJobs(newJobList(), NewAfterStartTimeFilter())
	assert.Empty(t, jl)
}
