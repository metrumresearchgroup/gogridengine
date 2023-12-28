package gogridengine

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type XMLDataSource struct {
	location string
}

func (d *XMLDataSource) Get() (string, error) {

	xmlresponse, err := http.Get(d.location)

	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(xmlresponse.Body)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

type XmlContentReader struct {
	resource XmlResourceGetter
}

func (x *XmlContentReader) Read() (string, error) {
	return x.resource.Get()
}

type XmlResourceGetter interface {
	Get() (string, error)
}

type XmlResourceReader interface {
	Read() (string, error)
}

// GetQstatOutput is used to pull in XML content from either the QSTAT command or generated data for testing purpoes
func GetQstatOutput(filters map[string]string) (string, error) {

	if os.Getenv(environmentPrefix+"TEST") != "true" {
		return qStatFromExec(filters)
	}

	//Fallthrough to generation by object randomly
	return generatedQstatOputput()
}

// DeleteQueuedJobByID is used to delete (1 or many) jobs by concatenating their IDs together and passing them to qdel
func DeleteQueuedJobByID(targets []string) (string, error) {

	//If this is in test mode, just return empty error and exit quickly
	if os.Getenv(environmentPrefix+"TEST") == "true" {
		outputs := []string{}
		for _, v := range targets {
			JobID, _ := strconv.ParseInt(v, 10, 64)
			outputs = append(outputs, fmt.Sprintf("username has deleted job %d", JobID))
		}

		return strings.Join(outputs, "\n"), nil
	}

	s := strings.Join(targets, ",")
	s = strings.TrimSpace(s)

	//Locate the binary in existing path
	binary, err := exec.LookPath("qdel")

	if err != nil {
		log.Error("Couldn't locate binary", err)
		return "", errors.New("Couldn't locate the binary")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//Cowardly cancel on any other exit mode
	defer cancel()

	log.Info("Requesting qdel with a list of IDs: ", s)
	command := exec.CommandContext(ctx, binary, s)
	command.Env = os.Environ()
	output := &bytes.Buffer{}
	command.Stdout = output
	err = command.Run()
	if err != nil {
		log.Error(output.String())
		return output.String(), err
	}

	return output.String(), nil
}

// DeleteQueuedJobByUsernames is used to delete (1 or many) jobs by concatenating usernames together and feeding them to qdel
func DeleteQueuedJobByUsernames(targets []string) (string, error) {

	//If this is in test mode, just return empty error and exit quickly
	if os.Getenv(environmentPrefix+"TEST") == "true" {
		responses := rand.Intn(1000)
		outputs := []string{}

		for i := 0; i < responses; i++ {
			jobID := rand.Intn(responses)
			outputs = append(outputs, fmt.Sprintf("username has deleted job %d", jobID))
		}

		return strings.Join(outputs, "\n"), nil
	}

	s := strings.Join(targets, ",")
	s = strings.TrimSpace(s)

	//Locate the binary in existing path
	binary, err := exec.LookPath("qdel")

	if err != nil {
		log.Error("Couldn't locate binary", err)
		return "", errors.New("Couldn't locate the binary")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//Cowardly cancel on any other exit mode
	defer cancel()

	log.Info("Running qdel with the following user input ", s)
	command := exec.CommandContext(ctx, binary, "-u", s)
	command.Env = os.Environ()
	output := &bytes.Buffer{}
	command.Stdout = output
	err = command.Run()
	if err != nil {
		log.Error(output.String())
		log.Error(err)
		return output.String(), err
	}

	return output.String(), nil
}

// Filters are meant to be in the form of [key] being being a switch and the value to be the anything passed to the option
func qStatFromExec(filters map[string]string) (string, error) {

	//Locate the binary in existing path
	binary, err := exec.LookPath("qstat")

	if err != nil {
		log.Error("Couldn't locate binary", err)
		return "", errors.New("Couldn't locate the binary")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	//Cowardly cancel on any other exit mode
	defer cancel()

	arguments := buildQstatArgumentList(filters)

	command := exec.CommandContext(ctx, binary, arguments...)
	command.Env = os.Environ()
	outputBytes, err := command.CombinedOutput()

	if err != nil {
		log.Errorf("An error occurred during execution of the the binary %s. Execution details are %s ", binary, string(outputBytes))
		return "", fmt.Errorf("an error occurred during execution of the the binary %s. Execution details are %s. Error: %w", binary, string(outputBytes), err)
	}

	return string(outputBytes), nil
}

func buildQstatArgumentList(filters map[string]string) []string {
	var arguments []string
	userFiltered := false

	//Let's iterate over all the provided kvps
	for k, v := range filters {
		//If a user has been provided, let's specify those users
		if k == "-u" {
			userFiltered = true
		}

		//If the value is empty, we'll only pass the key
		if len(v) == 0 {
			arguments = append(arguments, k)
		}

		//Otherwise let's append both
		if len(v) > 0 {
			arguments = append(arguments, k, v)
		}

	}

	//No user provided, let's make sure to provide the details for listing all users
	if !userFiltered {
		arguments = append(arguments, "-u", "*")
	}

	//Always add the strictest requirements last (IE the Full output and XML Compoenent)
	arguments = append(arguments, "-F", "-xml")

	return arguments
}

func generatedQstatOputput() (string, error) {

	return `<?xml version='1.0'?>
<job_info  xmlns:xsd="https://github.com/gridengine/gridengine/raw/master/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>broad@devuger-c001.broadinstitute.org</name>
      <qtype>BP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>24</slots_total>
      <arch>lx-amd64</arch>
      <state>u</state>
      <resource name="genetorrent" type="gc">50</resource>
      <resource name="matlab" type="gc">0</resource>
      <resource name="arch" type="hl">lx-amd64</resource>
      <resource name="num_proc" type="hl">24</resource>
      <resource name="mem_total" type="hl">251.689G</resource>
      <resource name="swap_total" type="hl">4.000G</resource>
      <resource name="virtual_total" type="hl">255.689G</resource>
      <resource name="m_topology" type="hl">SCCCCCCCCCCCCSCCCCCCCCCCCC</resource>
      <resource name="m_topology_inuse" type="hl">SCCCCCCCCCCCCSCCCCCCCCCCCC</resource>
      <resource name="m_socket" type="hl">2</resource>
      <resource name="m_core" type="hl">24</resource>
      <resource name="m_thread" type="hl">24</resource>
      <resource name="h_rt" type="hf">00:44:30</resource>
      <resource name="m_mem_free" type="hf">251.688G</resource>
      <resource name="slots" type="hc">24</resource>
      <resource name="processor" type="hf">Intel(R) Xeon(R) CPU E5-2695 v2 @ 2.40GHz</resource>
      <resource name="operating_system" type="hf">RedHat7</resource>
      <resource name="h_vmem" type="hc">239.688G</resource>
      <resource name="qname" type="qf">broad</resource>
      <resource name="hostname" type="qf">devuger-c001.broadinstitute.org</resource>
      <resource name="seq_no" type="qf">10</resource>
      <resource name="rerun" type="qf">1</resource>
      <resource name="tmpdir" type="qf">/local/scratch</resource>
      <resource name="calendar" type="qf">NONE</resource>
      <resource name="s_rt" type="qf">infinity</resource>
      <resource name="d_rt" type="qf">infinity</resource>
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
      <resource name="min_cpu_interval" type="qf">00:05:00</resource>
      <resource name="concurjob" type="qc">999999</resource>
    </Queue-List>
    <Queue-List>
      <name>broad@devuger-c002.broadinstitute.org</name>
      <qtype>BP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>20</slots_total>
      <np_load_avg>0.00019</np_load_avg>
      <arch>lx-amd64</arch>
      <resource name="genetorrent" type="gc">50</resource>
      <resource name="matlab" type="gc">0</resource>
      <resource name="arch" type="hl">lx-amd64</resource>
      <resource name="num_proc" type="hl">52</resource>
      <resource name="mem_total" type="hl">503.357G</resource>
      <resource name="swap_total" type="hl">4.000G</resource>
      <resource name="virtual_total" type="hl">507.357G</resource>
      <resource name="m_topology" type="hl">SCCCCCCCCCCCCCCCCCCCCCCCCCCSCCCCCCCCCCCCCCCCCCCCCCCCCC</resource>
      <resource name="m_topology_inuse" type="hl">SCCCCCCCCCCCCCCCCCCCCCCCCCCSCCCCCCCCCCCCCCCCCCCCCCCCCC</resource>
      <resource name="m_socket" type="hl">2</resource>
      <resource name="m_core" type="hl">52</resource>
      <resource name="m_thread" type="hl">52</resource>
      <resource name="load_avg" type="hl">0.010000</resource>
      <resource name="load_short" type="hl">0.000000</resource>
      <resource name="load_medium" type="hl">0.010000</resource>
      <resource name="load_long" type="hl">0.050000</resource>
      <resource name="mem_free" type="hl">497.185G</resource>
      <resource name="swap_free" type="hl">4.000G</resource>
      <resource name="virtual_free" type="hl">501.185G</resource>
      <resource name="mem_used" type="hl">6.172G</resource>
      <resource name="swap_used" type="hl">0.000</resource>
      <resource name="virtual_used" type="hl">6.172G</resource>
      <resource name="cpu" type="hl">0.000000</resource>
      <resource name="m_cache_l1" type="hl">32.000K</resource>
      <resource name="m_cache_l2" type="hl">1.000M</resource>
      <resource name="m_cache_l3" type="hl">35.750M</resource>
      <resource name="m_mem_total" type="hl">503.356G</resource>
      <resource name="m_mem_used" type="hl">13.188G</resource>
      <resource name="m_mem_free" type="hl">490.168G</resource>
      <resource name="m_numa_nodes" type="hl">1</resource>
      <resource name="m_topology_numa" type="hl">[SCCCCCCCCCCCCCCCCCCCCCCCCCCSCCCCCCCCCCCCCCCCCCCCCCCCCC]</resource>
      <resource name="docker" type="hl">0</resource>
      <resource name="np_load_avg" type="hl">0.000192</resource>
      <resource name="np_load_short" type="hl">0.000000</resource>
      <resource name="np_load_medium" type="hl">0.000192</resource>
      <resource name="np_load_long" type="hl">0.000962</resource>
      <resource name="h_rt" type="hf">30:00:00:00</resource>
      <resource name="operating_system" type="hf">RedHat7</resource>
      <resource name="slots" type="qc">20</resource>
      <resource name="h_vmem" type="hc">491.356G</resource>
      <resource name="processor" type="hf">Intel(R) Xeon(R) Gold 6230R CPU @ 2.10GHz</resource>
      <resource name="qname" type="qf">broad</resource>
      <resource name="hostname" type="qf">devuger-c002.broadinstitute.org</resource>
      <resource name="seq_no" type="qf">10</resource>
      <resource name="rerun" type="qf">1</resource>
      <resource name="tmpdir" type="qf">/local/scratch</resource>
      <resource name="calendar" type="qf">NONE</resource>
      <resource name="s_rt" type="qf">infinity</resource>
      <resource name="d_rt" type="qf">infinity</resource>
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
      <resource name="min_cpu_interval" type="qf">00:05:00</resource>
      <resource name="concurjob" type="qc">999999</resource>
    </Queue-List>
    <Queue-List>
      <name>interactive@devuger-c001.broadinstitute.org</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>24</slots_total>
      <arch>lx-amd64</arch>
      <state>u</state>
      <resource name="genetorrent" type="gc">50</resource>
      <resource name="matlab" type="gc">0</resource>
      <resource name="arch" type="hl">lx-amd64</resource>
      <resource name="num_proc" type="hl">24</resource>
      <resource name="mem_total" type="hl">251.689G</resource>
      <resource name="swap_total" type="hl">4.000G</resource>
      <resource name="virtual_total" type="hl">255.689G</resource>
      <resource name="m_topology" type="hl">SCCCCCCCCCCCCSCCCCCCCCCCCC</resource>
      <resource name="m_topology_inuse" type="hl">SCCCCCCCCCCCCSCCCCCCCCCCCC</resource>
      <resource name="m_socket" type="hl">2</resource>
      <resource name="m_core" type="hl">24</resource>
      <resource name="m_thread" type="hl">24</resource>
      <resource name="h_rt" type="hf">00:44:30</resource>
      <resource name="m_mem_free" type="hf">251.688G</resource>
      <resource name="slots" type="hc">24</resource>
      <resource name="processor" type="hf">Intel(R) Xeon(R) CPU E5-2695 v2 @ 2.40GHz</resource>
      <resource name="operating_system" type="hf">RedHat7</resource>
      <resource name="h_vmem" type="hc">239.688G</resource>
      <resource name="qname" type="qf">interactive</resource>
      <resource name="hostname" type="qf">devuger-c001.broadinstitute.org</resource>
      <resource name="seq_no" type="qf">20</resource>
      <resource name="rerun" type="qf">1</resource>
      <resource name="tmpdir" type="qf">/local/scratch</resource>
      <resource name="calendar" type="qf">NONE</resource>
      <resource name="s_rt" type="qf">infinity</resource>
      <resource name="d_rt" type="qf">infinity</resource>
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
      <resource name="min_cpu_interval" type="qf">00:05:00</resource>
      <resource name="concurjob" type="qc">999999</resource>
    </Queue-List>
    <Queue-List>
      <name>interactive@devuger-c002.broadinstitute.org</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>20</slots_total>
      <np_load_avg>0.00019</np_load_avg>
      <arch>lx-amd64</arch>
      <resource name="genetorrent" type="gc">50</resource>
      <resource name="matlab" type="gc">0</resource>
      <resource name="arch" type="hl">lx-amd64</resource>
      <resource name="num_proc" type="hl">52</resource>
      <resource name="mem_total" type="hl">503.357G</resource>
      <resource name="swap_total" type="hl">4.000G</resource>
      <resource name="virtual_total" type="hl">507.357G</resource>
      <resource name="m_topology" type="hl">SCCCCCCCCCCCCCCCCCCCCCCCCCCSCCCCCCCCCCCCCCCCCCCCCCCCCC</resource>
      <resource name="m_topology_inuse" type="hl">SCCCCCCCCCCCCCCCCCCCCCCCCCCSCCCCCCCCCCCCCCCCCCCCCCCCCC</resource>
      <resource name="m_socket" type="hl">2</resource>
      <resource name="m_core" type="hl">52</resource>
      <resource name="m_thread" type="hl">52</resource>
      <resource name="load_avg" type="hl">0.010000</resource>
      <resource name="load_short" type="hl">0.000000</resource>
      <resource name="load_medium" type="hl">0.010000</resource>
      <resource name="load_long" type="hl">0.050000</resource>
      <resource name="mem_free" type="hl">497.185G</resource>
      <resource name="swap_free" type="hl">4.000G</resource>
      <resource name="virtual_free" type="hl">501.185G</resource>
      <resource name="mem_used" type="hl">6.172G</resource>
      <resource name="swap_used" type="hl">0.000</resource>
      <resource name="virtual_used" type="hl">6.172G</resource>
      <resource name="cpu" type="hl">0.000000</resource>
      <resource name="m_cache_l1" type="hl">32.000K</resource>
      <resource name="m_cache_l2" type="hl">1.000M</resource>
      <resource name="m_cache_l3" type="hl">35.750M</resource>
      <resource name="m_mem_total" type="hl">503.356G</resource>
      <resource name="m_mem_used" type="hl">13.188G</resource>
      <resource name="m_mem_free" type="hl">490.168G</resource>
      <resource name="m_numa_nodes" type="hl">1</resource>
      <resource name="m_topology_numa" type="hl">[SCCCCCCCCCCCCCCCCCCCCCCCCCCSCCCCCCCCCCCCCCCCCCCCCCCCCC]</resource>
      <resource name="docker" type="hl">0</resource>
      <resource name="np_load_avg" type="hl">0.000192</resource>
      <resource name="np_load_short" type="hl">0.000000</resource>
      <resource name="np_load_medium" type="hl">0.000192</resource>
      <resource name="np_load_long" type="hl">0.000962</resource>
      <resource name="h_rt" type="qf">1:12:00:00</resource>
      <resource name="operating_system" type="hf">RedHat7</resource>
      <resource name="slots" type="qc">20</resource>
      <resource name="h_vmem" type="hc">491.356G</resource>
      <resource name="processor" type="hf">Intel(R) Xeon(R) Gold 6230R CPU @ 2.10GHz</resource>
      <resource name="qname" type="qf">interactive</resource>
      <resource name="hostname" type="qf">devuger-c002.broadinstitute.org</resource>
      <resource name="seq_no" type="qf">20</resource>
      <resource name="rerun" type="qf">1</resource>
      <resource name="tmpdir" type="qf">/local/scratch</resource>
      <resource name="calendar" type="qf">NONE</resource>
      <resource name="s_rt" type="qf">infinity</resource>
      <resource name="d_rt" type="qf">infinity</resource>
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
      <resource name="min_cpu_interval" type="qf">00:05:00</resource>
      <resource name="concurjob" type="qc">999999</resource>
    </Queue-List>
  </queue_info>
  <job_info>
  </job_info>
</job_info>`, nil

	xmlLocation := XMLDataSource{location: "https://raw.githubusercontent.com/metrumresearchgroup/gogridengine/master/test_data/medium.xml"}

	if os.Getenv("GOGRIDENGINE_TEST_SOURCE") != "" {
		xmlLocation.location = os.Getenv("GOGRIDENGINE_TEST_SOURCE")
	}

	xmlResp := &XmlContentReader{resource: &xmlLocation}

	return xmlResp.Read()
}
