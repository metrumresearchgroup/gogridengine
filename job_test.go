package gogridengine

import (
	"encoding/xml"
	"os"
	"testing"
	"time"

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

	assert.Equal(t, 772, len(jobs))

	for _, v := range jobs {
		if v.State == "r" {
			runningCount++
			continue
		}
		pendingCount++
	}

	assert.Equal(t, 761, runningCount)
	assert.Equal(t, 11, pendingCount)
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

	assert.Len(t, jobs, 761)

	os.Unsetenv(environmentPrefix + "TEST")

	//Test Negative Path
	jobs, err := GetJobsWithFilter(func(j Job) bool {
		return j.State == "r"
	})

	assert.NotNil(t, err)
	assert.Empty(t, jobs)
}

func TestJobList_Sort(t *testing.T) {
	jl := JobList{
		{
			JBJobNumber: 1,
		},
		{
			JBJobNumber: 3,
		},
		{
			JBJobNumber: 2,
		},
	}

	jl = jl.Sort(func(i, j int) bool {
		return jl[i].JBJobNumber < jl[j].JBJobNumber
	})

	assert.Len(t, jl, 3)
	assert.Equal(t, Job{JBJobNumber: 2}, jl[1])
}

func TestFilterJobs(t *testing.T) {

	jl := JobList{
		{
			JobName: "Meow",
		},
		{
			JobName: "Woof",
		},
		{
			JobName: "Moo",
		},
	}

	r1 := FilterJobs(jl, func(j Job) bool {
		//Cow / Dog Filter
		if j.JobName == "Meow" || j.JobName == "Woof" {
			return true
		}

		return false
	})

	assert.NotEmpty(t, r1)
	assert.Len(t, r1, 2)

}

func TestExtrapolateTasksToJobs(t *testing.T) {
	j := Job{
		StateAttribute: "pending",
		State:          "qw",
		JBJobNumber:    545,
		SubmittedTime:  time.Now().String(),
		Tasks: Task{
			Source: "40-50:1",
		},
	}

	//Should have 11 jobs 40-50 inclusive (40-50 incremented by 1)
	jl, err := ExtrapolateTasksToJobs(j)

	assert.Nil(t, err)
	assert.NotEmpty(t, jl)

	assert.Len(t, jl, 11)

	for i := 0; i < 10; i++ {
		assert.Equal(t, int64(40+i), jl[i].Tasks.TaskID)
	}

	assert.Equal(t, int64(40), jl[0].Tasks.TaskID)
	assert.Equal(t, int64(50), jl[len(jl)-1].Tasks.TaskID)

	j = Job{
		StateAttribute: "pending",
		State:          "qw",
		JBJobNumber:    545,
		SubmittedTime:  time.Now().String(),
		Tasks: Task{
			Source: "40-50:5",
		},
	}

	jl, err = ExtrapolateTasksToJobs(j)

	assert.Nil(t, err)

	assert.Len(t, jl, 3)

	assert.Equal(t, int64(40), jl[0].Tasks.TaskID)
	assert.Equal(t, int64(50), jl[len(jl)-1].Tasks.TaskID)

	//Let's test some invalid job content

	j.Tasks.Source = "meow-cat:jupiter"

	jl, err = ExtrapolateTasksToJobs(j)

	assert.Error(t, err)
	assert.Equal(t, err, ErrInvalidTaskRangeIdentifier)

}

func TestTaskDeSerialization(t *testing.T) {
	input := `<?xml version='1.0'?>
	<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
	  <queue_info>
		<Queue-List>
		  <name>all.q@ip-10-0-1-80.ec2.internal</name>
		  <qtype>BIP</qtype>
		  <slots_used>8</slots_used>
		  <slots_resv>0</slots_resv>
		  <slots_total>8</slots_total>
		  <load_avg>0.06000</load_avg>
		  <arch>lx-amd64</arch>
		  <resource name="load_avg" type="hl">0.060000</resource>
		  <resource name="load_short" type="hl">0.090000</resource>
		  <resource name="load_medium" type="hl">0.060000</resource>
		  <resource name="load_long" type="hl">0.010000</resource>
		  <resource name="arch" type="hl">lx-amd64</resource>
		  <resource name="num_proc" type="hl">8</resource>
		  <resource name="mem_free" type="hl">13.251G</resource>
		  <resource name="swap_free" type="hl">0.000</resource>
		  <resource name="virtual_free" type="hl">13.251G</resource>
		  <resource name="mem_total" type="hl">14.686G</resource>
		  <resource name="swap_total" type="hl">0.000</resource>
		  <resource name="virtual_total" type="hl">14.686G</resource>
		  <resource name="mem_used" type="hl">1.435G</resource>
		  <resource name="swap_used" type="hl">0.000</resource>
		  <resource name="virtual_used" type="hl">1.435G</resource>
		  <resource name="cpu" type="hl">0.800000</resource>
		  <resource name="m_topology" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_topology_inuse" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_socket" type="hl">1</resource>
		  <resource name="m_core" type="hl">4</resource>
		  <resource name="m_thread" type="hl">8</resource>
		  <resource name="np_load_avg" type="hl">0.007500</resource>
		  <resource name="np_load_short" type="hl">0.011250</resource>
		  <resource name="np_load_medium" type="hl">0.007500</resource>
		  <resource name="np_load_long" type="hl">0.001250</resource>
		  <resource name="qname" type="qf">all.q</resource>
		  <resource name="hostname" type="qf">ip-10-0-1-80.ec2.internal</resource>
		  <resource name="slots" type="qc">0</resource>
		  <resource name="tmpdir" type="qf">/tmp</resource>
		  <resource name="seq_no" type="qf">0</resource>
		  <resource name="rerun" type="qf">0.000000</resource>
		  <resource name="calendar" type="qf">NONE</resource>
		  <resource name="s_rt" type="qf">infinity</resource>
		  <resource name="h_rt" type="qf">infinity</resource>
		  <resource name="s_cpu" type="qf">infinity</resource>
		  <resource name="h_cpu" type="qf">infinity</resource>
		  <resource name="s_fsize" type="qf">infinity</resource>
		  <resource name="h_fsize" type="qf">infinity</resource>
		  <resource name="s_data" type="qf">infinity</resource>
		  <resource name="h_data" type="qf">infinity</resource>
		  <resource name="s_stack" type="qf">infinity</resource>
		  <resource name="h_stack" type="qf">infinity</resource>
		  <resource name="s_core" type="qf">infinity</resource>
		  <resource name="h_core" type="qf">infinity</resource>
		  <resource name="s_rss" type="qf">infinity</resource>
		  <resource name="h_rss" type="qf">infinity</resource>
		  <resource name="s_vmem" type="qf">infinity</resource>
		  <resource name="h_vmem" type="qf">infinity</resource>
		  <resource name="min_cpu_interval" type="qf">00:05:00</resource>
		  <job_list state="running">
			<JB_job_number>1005</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>cat</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>10</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>15</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>20</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>25</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>30</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>35</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>40</tasks>
		  </job_list>
		</Queue-List>
	  </queue_info>
	  <job_info>
		<job_list state="pending">
		  <JB_job_number>1006</JB_job_number>
		  <JAT_prio>0.55500</JAT_prio>
		  <JB_name>task_array.sh</JB_name>
		  <JB_owner>darrellb</JB_owner>
		  <state>qw</state>
		  <JB_submission_time>2019-11-15T11:31:34</JB_submission_time>
		  <slots>1</slots>
		  <tasks>41-150:1</tasks>
		</job_list>
	  </job_info>
	</job_info>`

	_, err := NewJobInfo(input)

	//Verify that unparseable Task content will generate errors.
	assert.Error(t, err)
}

func TestTaskRangeSerialization(t *testing.T) {

	//Normal
	ji := JobInfo{
		QueueInfo: QueueInfo{
			Queues: []Host{
				{
					Name: "Oh hai",
					JobList: []Job{
						{
							JBJobNumber: 100,
							Tasks: Task{
								Source: "10",
								TaskID: 10,
							},
						},
					},
				},
			},
		},
	}

	x, err := xml.Marshal(ji)

	assert.Nil(t, err)
	assert.NotEmpty(t, x)
	assert.Contains(t, string(x), "<tasks>10</tasks>")

	//0 Task --> Make sure the XML Representation doesn't have the empty, serialized Task content.
	ji.QueueInfo.Queues[0].JobList[0].Tasks.Source = "0"
	ji.QueueInfo.Queues[0].JobList[0].Tasks.TaskID = 0

	x, err = xml.Marshal(ji)

	assert.Nil(t, err)

	assert.NotContains(t, string(x), "<tasks>0</tasks>")

	//Pending Job
	ji.QueueInfo.Queues[0].JobList[0].Tasks.Source = "50-125:5"
	ji.QueueInfo.Queues[0].JobList[0].Tasks.TaskID = 0

	x, err = xml.Marshal(ji)

	assert.Nil(t, err)

	assert.Contains(t, string(x), "<tasks>50-125:5</tasks>")

	//Unparseable Source should still return content of the Task ID
	ji.QueueInfo.Queues[0].JobList[0].Tasks.Source = "cat"
	ji.QueueInfo.Queues[0].JobList[0].Tasks.TaskID = 123

	x, err = xml.Marshal(ji)

	assert.Nil(t, err)
	assert.Contains(t, string(x), "<tasks>123</tasks>")
}

func TestTaskGroupExtrapolation(t *testing.T) {

	//Normal
	ji := JobInfo{
		QueueInfo: QueueInfo{
			Queues: []Host{
				{
					Name: "Oh hai",
					JobList: []Job{
						{
							JBJobNumber: 100,
							Tasks: Task{
								Source: "10,11,13",
							},
						},
					},
				},
			},
		},
	}

	nji, err := ExtrapolateTasksToJobs(ji.QueueInfo.Queues[0].JobList[0])

	assert.Nil(t, err)
	assert.NotEmpty(t, nji)
	assert.Len(t, nji, 3)

	ji = JobInfo{
		QueueInfo: QueueInfo{
			Queues: []Host{
				{
					Name: "Oh hai",
					JobList: []Job{
						{
							JBJobNumber: 100,
							Tasks: Task{
								Source: "10-15:1",
							},
						},
					},
				},
			},
		},
	}

	nji, err = ExtrapolateTasksToJobs(ji.QueueInfo.Queues[0].JobList[0])

	assert.Nil(t, err)
	assert.Len(t, nji, 6)

	//Test for failure components. IE a split value in a group that can't be converted to a number
	//Normal
	ji = JobInfo{
		QueueInfo: QueueInfo{
			Queues: []Host{
				{
					Name: "Oh hai",
					JobList: []Job{
						{
							JBJobNumber: 100,
							Tasks: Task{
								Source: "10,dog,()",
							},
						},
					},
				},
			},
		},
	}

	nji, err = ExtrapolateTasksToJobs(ji.QueueInfo.Queues[0].JobList[0])

	assert.NotNil(t, err)
	assert.Error(t, err)

}
