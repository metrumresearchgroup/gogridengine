package gogridengine

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestDeserializeQueueList(t *testing.T) {
	source := `<Queue-List>
 <name>all.q@ip-172-16-2-102.us-west-2.compute.internal</name>
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
 <job_list state="running">
   <JB_job_number>4294</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run490</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4297</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run493</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4300</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run496</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4303</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run499</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4306</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run502</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4309</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run505</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4312</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run508</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4315</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run511</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4318</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run514</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4321</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run517</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4324</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run520</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4327</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run523</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4330</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run526</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4333</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run529</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4336</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run532</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4339</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run535</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4342</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run538</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4345</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run541</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4348</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run544</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4351</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run547</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4354</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run550</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4357</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run553</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4360</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run556</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4363</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run559</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4366</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run562</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4372</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run568</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4375</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run571</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4378</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run574</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4381</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run577</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
 <job_list state="running">
   <JB_job_number>4384</JB_job_number>
   <JAT_prio>0.50500</JAT_prio>
   <JB_name>Run580</JB_name>
   <JB_owner>ahmede</JB_owner>
   <state>r</state>
   <JAT_start_time>2019-09-15T15:26:36</JAT_start_time>
   <slots>1</slots>
 </job_list>
</Queue-List>`

	var ql Host
	xml.Unmarshal([]byte(source), &ql)

	if ql.Name == "" {
		t.Errorf("We didn't parse a name for some reason")
	}
}

func TestSerializeToXML(t *testing.T) {
	QueueList := Host{
		XMLName: xml.Name{
			Local: "Queue-List",
		},
		JobList: []Job{
			{
				XMLName: xml.Name{
					Local: "job_list",
				},
				JATPriority: 1.04,
				JBJobNumber: 4,
				JobName:     "Meow01",
				JobOwner:    "Darrell Breeden",
				Slots:       4,
				State:       "running",
			},
		},
		LoadAverage: 4.04,
		Resources: ResourceList{
			Resource{
				Name:  "free_mem",
				Type:  "hl",
				Value: "1.04G",
			},
		},
		SlotsTotal:    4,
		SlotsReserved: 2,
		SlotsUsed:     2,
	}

	output, err := xml.Marshal(QueueList)

	if err != nil {
		t.Errorf("Could not marshall object correctly")
	}

	formatted := string(output)

	if !strings.Contains(formatted, "Meow01") {
		t.Errorf("Does not contain one of the raw components")
	}
}
