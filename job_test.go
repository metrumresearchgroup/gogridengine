package gogridengine

import (
	"encoding/xml"
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

	var t2 JobList
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
