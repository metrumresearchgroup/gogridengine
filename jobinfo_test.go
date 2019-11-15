package gogridengine

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeSerializeSGEQStat(t *testing.T) {
	source := `<?xml version='1.0'?>
 <job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
   <queue_info>
	 <Queue-List>
	   <name>all.q@magicalhostname</name>
	   <qtype>BIP</qtype>
	   <slots_used>32</slots_used>
	   <slots_resv>0</slots_resv>
	   <slots_total>36</slots_total>
	   <load_avg>31.63000</load_avg>
	   <arch>lx-amd64</arch>
	   <resource name="load_avg" type="hl">31.630000</resource>
	   <resource name="load_short" type="hl">31.700000</resource>
	   <resource name="load_medium" type="hl">31.630000</resource>
	   <resource name="load_long" type="hl">31.680000</resource>
	   <resource name="arch" type="hl">lx-amd64</resource>
	   <resource name="num_proc" type="hl">36</resource>
	   <resource name="mem_free" type="hl">57.353G</resource>
	   <resource name="swap_free" type="hl">0.000</resource>
	   <resource name="virtual_free" type="hl">57.353G</resource>
	   <resource name="mem_total" type="hl">58.973G</resource>
	   <resource name="swap_total" type="hl">0.000</resource>
	   <resource name="virtual_total" type="hl">58.973G</resource>
	   <resource name="mem_used" type="hl">1.619G</resource>
	   <resource name="swap_used" type="hl">0.000</resource>
	   <resource name="virtual_used" type="hl">1.619G</resource>
	   <resource name="cpu" type="hl">89.100000</resource>
	   <resource name="m_topology" type="hl">SCTTCTTCTTCTTCTTCTTCTTCTTCTTSCTTCTTCTTCTTCTTCTTCTTCTTCTT</resource>
	   <resource name="m_topology_inuse" type="hl">SCTTCTTCTTCTTCTTCTTCTTCTTCTTSCTTCTTCTTCTTCTTCTTCTTCTTCTT</resource>
	   <resource name="m_socket" type="hl">2</resource>
	   <resource name="m_core" type="hl">18</resource>
	   <resource name="m_thread" type="hl">36</resource>
	   <resource name="np_load_avg" type="hl">0.878611</resource>
	   <resource name="np_load_short" type="hl">0.880556</resource>
	   <resource name="np_load_medium" type="hl">0.878611</resource>
	   <resource name="np_load_long" type="hl">0.880000</resource>
	   <resource name="qname" type="qf">all.q</resource>
	   <resource name="hostname" type="qf">ip-172-16-2-102.us-west-2.compute.internal</resource>
	   <resource name="slots" type="qc">4</resource>
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
		 <JB_job_number>4282</JB_job_number>
		 <JAT_prio>0.50500</JAT_prio>
		 <JB_name>Run478</JB_name>
		 <JB_owner>user</JB_owner>
		 <state>r</state>
		 <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
		 <slots>1</slots>
	   </job_list>
	 </Queue-List>
   </queue_info>
   <job_info>
   </job_info>
 </job_info>`

	var info JobInfo
	xml.Unmarshal([]byte(source), &info)

	if info.QueueInfo.Queues[0].Name != "all.q@magicalhostname" {
		t.Errorf("Failed to parse and extract in order")
	}

	if len(info.QueueInfo.Queues[0].JobList) == 0 {
		t.Errorf("Failed to parse job details")
	}

	if info.QueueInfo.Queues[0].JobList[0].JobOwner != "user" {
		t.Errorf("Looks like we failed to serialize all the way down")
	}
}

func TestDeSerializePendingQStat(t *testing.T) {
	source := `<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
	<queue_info>
	  <Queue-List>
		<name>all.q@ip-10-0-1-87.ec2.internal</name>
		<qtype>BIP</qtype>
		<slots_used>0</slots_used>
		<slots_resv>0</slots_resv>
		<slots_total>8</slots_total>
		<load_avg>0.82000</load_avg>
		<arch>lx-amd64</arch>
		<resource name="load_avg" type="hl">0.820000</resource>
		<resource name="load_short" type="hl">0.510000</resource>
		<resource name="load_medium" type="hl">0.820000</resource>
		<resource name="load_long" type="hl">0.500000</resource>
		<resource name="arch" type="hl">lx-amd64</resource>
		<resource name="num_proc" type="hl">8</resource>
		<resource name="mem_free" type="hl">14.086G</resource>
		<resource name="swap_free" type="hl">0.000</resource>
		<resource name="virtual_free" type="hl">14.086G</resource>
		<resource name="mem_total" type="hl">14.686G</resource>
		<resource name="swap_total" type="hl">0.000</resource>
		<resource name="virtual_total" type="hl">14.686G</resource>
		<resource name="mem_used" type="hl">614.492M</resource>
		<resource name="swap_used" type="hl">0.000</resource>
		<resource name="virtual_used" type="hl">614.492M</resource>
		<resource name="cpu" type="hl">0.800000</resource>
		<resource name="m_topology" type="hl">SCTTCTTCTTCTT</resource>
		<resource name="m_topology_inuse" type="hl">SCTTCTTCTTCTT</resource>
		<resource name="m_socket" type="hl">1</resource>
		<resource name="m_core" type="hl">4</resource>
		<resource name="m_thread" type="hl">8</resource>
		<resource name="np_load_avg" type="hl">0.102500</resource>
		<resource name="np_load_short" type="hl">0.063750</resource>
		<resource name="np_load_medium" type="hl">0.102500</resource>
		<resource name="np_load_long" type="hl">0.062500</resource>
		<resource name="qname" type="qf">all.q</resource>
		<resource name="hostname" type="qf">ip-10-0-1-87.ec2.internal</resource>
		<resource name="slots" type="qc">8</resource>
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
			<JB_job_number>4282</JB_job_number>
			<JAT_prio>0.50500</JAT_prio>
			<JB_name>Run478</JB_name>
			<JB_owner>ahmede</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
			<slots>1</slots>
      	</job_list>
		<job_list state="running">
			<JB_job_number>4291</JB_job_number>
			<JAT_prio>0.50500</JAT_prio>
			<JB_name>Run487</JB_name>
			<JB_owner>ahmede</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
			<slots>1</slots>
		</job_list>
	  </Queue-List>
	</queue_info>
	<job_info>
	  <job_list state="pending">
		<JB_job_number>2</JB_job_number>
		<JAT_prio>0.00000</JAT_prio>
		<JB_name>test.sh</JB_name>
		<JB_owner>darrellb</JB_owner>
		<state>qw</state>
		<JB_submission_time>2019-09-26T15:42:29</JB_submission_time>
		<slots>1</slots>
	  </job_list>
	</job_info>
  </job_info>`

	var ji JobInfo
	err := xml.Unmarshal([]byte(source), &ji)
	if err != nil {
		t.Errorf(err.Error())
	}

	//Verify that we have at least one running job
	if len(ji.QueueInfo.Queues[0].JobList) <= 0 {
		t.Errorf("There are no running jobs in the first queue list serialized")
	}

	//Verify that we have pending jobs
	if len(ji.PendingJobs.JobList) <= 0 {
		t.Errorf("No pending jobs were serialized")
	}
}

func TestJobInfo_GetXML(t *testing.T) {

	os.Setenv(environmentPrefix+"TEST", "true")
	type fields struct {
		XMLName   xml.Name
		QueueInfo QueueInfo
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Verify has headers",
			fields: fields{
				XMLName: xml.Name{
					Local: "job_info",
				},
				QueueInfo: QueueInfo{
					XMLName: xml.Name{
						Local: "queue_info",
					},
					Queues: []Host{
						{
							XMLName: xml.Name{
								Local: "Queue-List",
							},
							Name:          "testing.local",
							SlotsTotal:    4,
							SlotsUsed:     1,
							SlotsReserved: 3,
							LoadAverage:   2.04,
							Resources: ResourceList{
								{
									Name:  "free_mem",
									Type:  "hl",
									Value: "1.4G",
								},
							},
							JobList: []Job{
								{
									XMLName: xml.Name{
										Local: "job_list",
									},
									StateAttribute: "running",
									State:          "r",
									JBJobNumber:    13,
									JATPriority:    1.04,
									JobName:        "Initial Test",
									JobOwner:       "You",
									Slots:          3,
								},
							},
						},
					},
				},
			},
			want:    `<?xml version='1.0'?><job_info><queue_info><Queue-List><name>testing.local</name><qtype></qtype><slots_used>1</slots_used><slots_rsv>3</slots_rsv><slots_total>4</slots_total><load_avg>2.04</load_avg><resource name="free_mem" type="hl">1.4G</resource><job_list state="running"><state>r</state><JB_job_number>13</JB_job_number><JAT_prio>1.04</JAT_prio><JB_name>Initial Test</JB_name><JB_owner>You</JB_owner><slots>3</slots></job_list></Queue-List></queue_info><job_info></job_info></job_info>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := JobInfo{
				XMLName:   tt.fields.XMLName,
				QueueInfo: tt.fields.QueueInfo,
			}
			got, err := q.GetXML()
			if (err != nil) != tt.wantErr {
				t.Errorf("JobInfo.GetXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				println(got)
				t.Errorf("JobInfo.GetXML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeserializingWithTasks(t *testing.T) {
	input := `<?xml version='1.0'?>
	<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
	  <queue_info>
		<Queue-List>
		  <name>all.q@ip-10-0-1-113.ec2.internal</name>
		  <qtype>BIP</qtype>
		  <slots_used>8</slots_used>
		  <slots_resv>0</slots_resv>
		  <slots_total>8</slots_total>
		  <load_avg>0.04000</load_avg>
		  <arch>lx-amd64</arch>
		  <resource name="load_avg" type="hl">0.040000</resource>
		  <resource name="load_short" type="hl">0.040000</resource>
		  <resource name="load_medium" type="hl">0.040000</resource>
		  <resource name="load_long" type="hl">0.000000</resource>
		  <resource name="arch" type="hl">lx-amd64</resource>
		  <resource name="num_proc" type="hl">8</resource>
		  <resource name="mem_free" type="hl">13.282G</resource>
		  <resource name="swap_free" type="hl">0.000</resource>
		  <resource name="virtual_free" type="hl">13.282G</resource>
		  <resource name="mem_total" type="hl">14.686G</resource>
		  <resource name="swap_total" type="hl">0.000</resource>
		  <resource name="virtual_total" type="hl">14.686G</resource>
		  <resource name="mem_used" type="hl">1.403G</resource>
		  <resource name="swap_used" type="hl">0.000</resource>
		  <resource name="virtual_used" type="hl">1.403G</resource>
		  <resource name="cpu" type="hl">0.800000</resource>
		  <resource name="m_topology" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_topology_inuse" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_socket" type="hl">1</resource>
		  <resource name="m_core" type="hl">4</resource>
		  <resource name="m_thread" type="hl">8</resource>
		  <resource name="np_load_avg" type="hl">0.005000</resource>
		  <resource name="np_load_short" type="hl">0.005000</resource>
		  <resource name="np_load_medium" type="hl">0.005000</resource>
		  <resource name="np_load_long" type="hl">0.000000</resource>
		  <resource name="qname" type="qf">all.q</resource>
		  <resource name="hostname" type="qf">ip-10-0-1-113.ec2.internal</resource>
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
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>1</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>7</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>11</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>17</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>21</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>27</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>31</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>37</tasks>
		  </job_list>
		</Queue-List>
		<Queue-List>
		  <name>all.q@ip-10-0-1-127.ec2.internal</name>
		  <qtype>BIP</qtype>
		  <slots_used>8</slots_used>
		  <slots_resv>0</slots_resv>
		  <slots_total>8</slots_total>
		  <load_avg>0.05000</load_avg>
		  <arch>lx-amd64</arch>
		  <resource name="load_avg" type="hl">0.050000</resource>
		  <resource name="load_short" type="hl">0.020000</resource>
		  <resource name="load_medium" type="hl">0.050000</resource>
		  <resource name="load_long" type="hl">0.020000</resource>
		  <resource name="arch" type="hl">lx-amd64</resource>
		  <resource name="num_proc" type="hl">8</resource>
		  <resource name="mem_free" type="hl">13.282G</resource>
		  <resource name="swap_free" type="hl">0.000</resource>
		  <resource name="virtual_free" type="hl">13.282G</resource>
		  <resource name="mem_total" type="hl">14.686G</resource>
		  <resource name="swap_total" type="hl">0.000</resource>
		  <resource name="virtual_total" type="hl">14.686G</resource>
		  <resource name="mem_used" type="hl">1.404G</resource>
		  <resource name="swap_used" type="hl">0.000</resource>
		  <resource name="virtual_used" type="hl">1.404G</resource>
		  <resource name="cpu" type="hl">0.800000</resource>
		  <resource name="m_topology" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_topology_inuse" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_socket" type="hl">1</resource>
		  <resource name="m_core" type="hl">4</resource>
		  <resource name="m_thread" type="hl">8</resource>
		  <resource name="np_load_avg" type="hl">0.006250</resource>
		  <resource name="np_load_short" type="hl">0.002500</resource>
		  <resource name="np_load_medium" type="hl">0.006250</resource>
		  <resource name="np_load_long" type="hl">0.002500</resource>
		  <resource name="qname" type="qf">all.q</resource>
		  <resource name="hostname" type="qf">ip-10-0-1-127.ec2.internal</resource>
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
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>3</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>9</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>13</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>19</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>23</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>29</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>33</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>39</tasks>
		  </job_list>
		</Queue-List>
		<Queue-List>
		  <name>all.q@ip-10-0-1-193.ec2.internal</name>
		  <qtype>BIP</qtype>
		  <slots_used>8</slots_used>
		  <slots_resv>0</slots_resv>
		  <slots_total>8</slots_total>
		  <load_avg>0.06000</load_avg>
		  <arch>lx-amd64</arch>
		  <resource name="load_avg" type="hl">0.060000</resource>
		  <resource name="load_short" type="hl">0.050000</resource>
		  <resource name="load_medium" type="hl">0.060000</resource>
		  <resource name="load_long" type="hl">0.010000</resource>
		  <resource name="arch" type="hl">lx-amd64</resource>
		  <resource name="num_proc" type="hl">8</resource>
		  <resource name="mem_free" type="hl">13.250G</resource>
		  <resource name="swap_free" type="hl">0.000</resource>
		  <resource name="virtual_free" type="hl">13.250G</resource>
		  <resource name="mem_total" type="hl">14.686G</resource>
		  <resource name="swap_total" type="hl">0.000</resource>
		  <resource name="virtual_total" type="hl">14.686G</resource>
		  <resource name="mem_used" type="hl">1.435G</resource>
		  <resource name="swap_used" type="hl">0.000</resource>
		  <resource name="virtual_used" type="hl">1.435G</resource>
		  <resource name="cpu" type="hl">1.100000</resource>
		  <resource name="m_topology" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_topology_inuse" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_socket" type="hl">1</resource>
		  <resource name="m_core" type="hl">4</resource>
		  <resource name="m_thread" type="hl">8</resource>
		  <resource name="np_load_avg" type="hl">0.007500</resource>
		  <resource name="np_load_short" type="hl">0.006250</resource>
		  <resource name="np_load_medium" type="hl">0.007500</resource>
		  <resource name="np_load_long" type="hl">0.001250</resource>
		  <resource name="qname" type="qf">all.q</resource>
		  <resource name="hostname" type="qf">ip-10-0-1-193.ec2.internal</resource>
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
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>2</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>6</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>12</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>16</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>22</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>26</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>32</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>36</tasks>
		  </job_list>
		</Queue-List>
		<Queue-List>
		  <name>all.q@ip-10-0-1-44.ec2.internal</name>
		  <qtype>BIP</qtype>
		  <slots_used>8</slots_used>
		  <slots_resv>0</slots_resv>
		  <slots_total>8</slots_total>
		  <load_avg>0.05000</load_avg>
		  <arch>lx-amd64</arch>
		  <resource name="load_avg" type="hl">0.050000</resource>
		  <resource name="load_short" type="hl">0.040000</resource>
		  <resource name="load_medium" type="hl">0.050000</resource>
		  <resource name="load_long" type="hl">0.020000</resource>
		  <resource name="arch" type="hl">lx-amd64</resource>
		  <resource name="num_proc" type="hl">8</resource>
		  <resource name="mem_free" type="hl">13.280G</resource>
		  <resource name="swap_free" type="hl">0.000</resource>
		  <resource name="virtual_free" type="hl">13.280G</resource>
		  <resource name="mem_total" type="hl">14.686G</resource>
		  <resource name="swap_total" type="hl">0.000</resource>
		  <resource name="virtual_total" type="hl">14.686G</resource>
		  <resource name="mem_used" type="hl">1.406G</resource>
		  <resource name="swap_used" type="hl">0.000</resource>
		  <resource name="virtual_used" type="hl">1.406G</resource>
		  <resource name="cpu" type="hl">0.800000</resource>
		  <resource name="m_topology" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_topology_inuse" type="hl">SCTTCTTCTTCTT</resource>
		  <resource name="m_socket" type="hl">1</resource>
		  <resource name="m_core" type="hl">4</resource>
		  <resource name="m_thread" type="hl">8</resource>
		  <resource name="np_load_avg" type="hl">0.006250</resource>
		  <resource name="np_load_short" type="hl">0.005000</resource>
		  <resource name="np_load_medium" type="hl">0.006250</resource>
		  <resource name="np_load_long" type="hl">0.002500</resource>
		  <resource name="qname" type="qf">all.q</resource>
		  <resource name="hostname" type="qf">ip-10-0-1-44.ec2.internal</resource>
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
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>4</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>8</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>14</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>18</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>24</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>28</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>34</tasks>
		  </job_list>
		  <job_list state="running">
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>38</tasks>
		  </job_list>
		</Queue-List>
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
			<JB_job_number>1006</JB_job_number>
			<JAT_prio>0.55500</JAT_prio>
			<JB_name>task_array.sh</JB_name>
			<JB_owner>darrellb</JB_owner>
			<state>r</state>
			<JAT_start_time>2019-11-15T11:31:47</JAT_start_time>
			<slots>1</slots>
			<tasks>5</tasks>
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

	var ji JobInfo

	xml.Unmarshal([]byte(input), &ji)

	assert.NotEmpty(t, ji)

	for _, q := range ji.QueueInfo.Queues {
		for _, j := range q.JobList {
			assert.True(t, j.Tasks.Source != "")
		}
	}

	for _, q := range ji.PendingJobs.JobList {
		assert.True(t, q.Tasks.Source != "")
	}
}
