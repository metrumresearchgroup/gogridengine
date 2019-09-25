package gogridengine

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//GetQstatOutput is used to pull in XML content from either the QSTAT command or generated data for testing purpoes
func GetQstatOutput() (string, error) {

	var inTestMode bool = false

	if os.Getenv("TEST") == "true" {
		inTestMode = true
	}

	if !inTestMode {
		return qStatFromExec()
	}

	//Fallthrough to generation by object randomly
	return generatedQstatOputput()
}

func qStatFromExec() (string, error) {

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

	command := exec.CommandContext(ctx, binary, "-F", "-xml")
	command.Env = os.Environ()
	log.Debug(command.Env)
	output := &bytes.Buffer{}
	command.Stdout = output
	err = command.Run()
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func generatedQstatOputput() (string, error) {

	entropy := rand.NewSource(time.Now().UnixNano())
	random := rand.New(entropy)

	ji := JobInfo{
		XMLName: xml.Name{
			Local: "job_info",
		},
		QueueInfo: QueueInfo{
			XMLName: xml.Name{
				Local: "queue_info",
			},
			Queues: []QueueList{
				{
					XMLName: xml.Name{
						Local: "Queue-List",
					},
					Name:          "all.q@testing.local", //Always needs the @ symbol
					SlotsTotal:    int32(random.Int()),
					SlotsUsed:     int32(random.Int()),
					SlotsReserved: int32(random.Int()),
					LoadAverage:   float64(random.Float64()),
					Resources: ResourceList{
						{
							Name:  "load_average",
							Type:  "hl",
							Value: "1.04",
						},
						{
							Name:  "num_proc",
							Type:  "ag",
							Value: "3",
						},
						{
							Name:  "mem_free",
							Type:  "af",
							Value: "2.04G",
						},
						{
							Name:  "swap_free",
							Type:  "ae",
							Value: "500M",
						},
						{
							Name:  "virtual_free",
							Type:  "ad",
							Value: "4G",
						},
						{
							Name:  "mem_used",
							Type:  "ac",
							Value: "3G",
						},
						{
							Name:  "mem_total",
							Type:  "ab",
							Value: "6G",
						},
						{
							Name:  "cpu",
							Type:  "aa",
							Value: fmt.Sprintf("%f", random.Float64()),
						},
					},
					JobList: []JobList{
						{
							XMLName: xml.Name{
								Local: "job_list",
							},
							State:       "running",
							JBJobNumber: int64(random.Int()),
							JATPriority: random.Float64(),
							JobName:     "Job-" + strconv.Itoa(random.Int()),
							JobOwner:    "Owner-" + strconv.Itoa(random.Int()),
							Slots:       3,
						},
					},
				},
				{
					XMLName: xml.Name{
						Local: "Queue-List",
					},
					Name: "all.q@testing.second", //Always needs the @ symbol

					Resources: ResourceList{},
					JobList: []JobList{
						{
							XMLName: xml.Name{
								Local: "job_list",
							},
							State:       "running",
							JBJobNumber: 1,
							JATPriority: 1,
							JobName:     "Second-Host-Job",
							JobOwner:    "Owner",
							Slots:       14,
						},
					},
				},
			},
		},
	}

	return ji.GetXML()
}
