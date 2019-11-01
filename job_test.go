package gogridengine

import (
	"encoding/xml"
	"os"
	"testing"
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
}
