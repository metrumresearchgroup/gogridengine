package gogridengine

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeSerializeXml(t *testing.T) {

	source := `<job_list state="running">
	<JB_job_number>4291</JB_job_number>
	<JAT_prio>0.50500</JAT_prio>
	<JB_name>Run487</JB_name>
	<JB_owner>ahmede</JB_owner>
	<state>r</state>
	<JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
	<slots>1</slots>
</job_list>`

	var t2 Job
	xml.Unmarshal([]byte(source), &t2)

	if t2.JBJobNumber != 4291 {
		t.Errorf("Invalid Job Number marshalled OR no job number marshalled")
	}

	if t2.JATPriority != 0.50500 {
		t.Errorf("Invalid Priority marshalled OR no priority marshalled")
	}

	if t2.JobName != "Run487" {
		t.Errorf("Invalid Job Name marshalled OR no job name marshalled")
	}

	if t2.JobOwner != "ahmede" {
		t.Errorf("Invalid Job Owner marshalled OR no job owner marshalled")
	}

	if t2.State != "r" {
		t.Errorf("Invalid State marshalled OR no state marshalled at all")
	}

	if t2.StartTime != "2019-09-15T15:26:36" {
		t.Errorf("Invalid start time recorded OR no start time recorded")
	}

	if t2.Slots != 1 {
		t.Errorf("Invalid slots value marshalled OR no slots value marshalled at all")
	}
}

func TestIsJobRunning(t *testing.T) {

	pending := `<job_list state="pending">
	<JB_job_number>3517</JB_job_number>
	<JAT_prio>0.55500</JAT_prio>
	<JB_name>Run1417</JB_name>
	<JB_owner>devinp</JB_owner>
	<state>qw</state>
	<JB_submission_time>2019-09-26T10:17:37</JB_submission_time>
	<slots>1</slots>
  </job_list>`

	running := `<job_list state="running">
  <JB_job_number>3517</JB_job_number>
  <JAT_prio>0.55500</JAT_prio>
  <JB_name>Run1417</JB_name>
  <JB_owner>devinp</JB_owner>
  <state>r</state>
  <JB_submission_time>2019-09-26T10:17:37</JB_submission_time>
  <slots>1</slots>
</job_list>`

	var pl Job
	var rl Job
	err := xml.Unmarshal([]byte(pending), &pl)
	if err != nil {
		t.Errorf("Unable to unmarshall xml")
	}

	err = xml.Unmarshal([]byte(running), &rl)
	if err != nil {
		t.Errorf("Unable to unmarshall xml")
	}

	type args struct {
		job Job
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Pending Job",
			args: args{
				job: pl,
			},
			want: 0,
		},
		{
			name: "Running Job",
			args: args{
				job: rl,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJobRunning(tt.args.job); got != tt.want {
				t.Errorf("IsJobRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJobs(t *testing.T) {
	os.Setenv(environmentPrefix+"TEST", "true")

	jobs, err := GetJobs()
	runningCount := 0
	pendingCount := 0

	//Definitely should not fail in test evaluation mode
	if err != nil {
		t.Error(err)
	}

	//Our defined structue is 3 running jobs and 1 pending jobs

	assert.Equal(t, 4, len(jobs))

	for _, v := range jobs {
		if v.State == "r" {
			runningCount++
			continue
		}
		pendingCount++
	}

	assert.Equal(t, 3, runningCount)
	assert.Equal(t, 1, pendingCount)
}

func TestJobFilters(t *testing.T) {
	os.Setenv(environmentPrefix+"TEST", "true")

	//Let's first Verify that passing parameters gets to the argument list correctly

	//Verify running empty processes fine.
	_, err := GetQstatOutput(make(map[string]string))

	//Exec component should still process with the fake binary
	assert.Nil(t, err)

	filters := make(map[string]string)

	filters["-u"] = "darrellb"
	filters["-s"] = "r"

	//Maps are unordered
	arguments := buildQstatArgumentList(filters)

	assert.Contains(t, arguments, "-u")
	assert.Contains(t, arguments, "darrellb")
	assert.Contains(t, arguments, "-s")
	assert.Contains(t, arguments, "r")

	//Get Key of User Switch
	var userIndex int

	for key, value := range arguments {
		if value == "-u" {
			userIndex = key
		}
	}

	assert.Equal(t, "darrellb", arguments[userIndex+1])

	assert.True(t, len(arguments) == (2*len(filters))+2)

	//Get State Index
	var stateIndex int

	for key, value := range arguments {
		if value == "-s" {
			stateIndex = key
		}
	}

	assert.Equal(t, "r", arguments[stateIndex+1])

	assert.True(t, len(arguments) == (2*len(filters))+2)

	//Now, let's verify the argument list for an unspecified filter
	filters = make(map[string]string)

	expectedArgs := []string{
		"-u",
		"*",
		"-F",
		"-xml",
	}

	generatedArgs := buildQstatArgumentList(filters)

	assert.Equal(t, expectedArgs, generatedArgs)
}

func TestGetJobsWithFilter(t *testing.T) {
	os.Setenv(environmentPrefix+"TEST", "true")
	jobs, _ := GetJobsWithFilter(func(j Job) bool {
		return j.State == "r"
	})

	assert.Len(t, jobs, 3)

	os.Unsetenv(environmentPrefix + "TEST")

	//Test Negative Path
	jobs, err := GetJobsWithFilter(func(j Job) bool {
		return j.State == "r"
	})

	assert.NotNil(t, err)
	assert.Empty(t, jobs)
}
