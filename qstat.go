package gogridengine

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

//GetQstatOutput is used to pull in XML content from either the QSTAT command or generated data for testing purpoes
func GetQstatOutput(filters map[string]string) (string, error) {

	if os.Getenv(environmentPrefix+"TEST") != "true" {
		return qStatFromExec(filters)
	}

	//Fallthrough to generation by object randomly
	return generatedQstatOputput()
}

//DeleteQueuedJobByID is used to delete (1 or many) jobs by concatenating their IDs together and passing them to qdel
func DeleteQueuedJobByID(targets []string) (string, error) {

	//If this is in test mode, just return empty error and exit quickly
	if os.Getenv(environmentPrefix+"TEST") == "true" {
		return "test", nil
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

//DeleteQueuedJobByUsernames is used to delete (1 or many) jobs by concatenating usernames together and feeding them to qdel
func DeleteQueuedJobByUsernames(targets []string) (string, error) {

	//If this is in test mode, just return empty error and exit quickly
	if os.Getenv(environmentPrefix+"TEST") == "true" {
		return "test", nil
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

//Filters are meant to be in the form of [key] being being a switch and the value to be the anything passed to the option
func qStatFromExec(filters map[string]string) (string, error) {

	var arguments []string
	userFiltered := false

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

	command := exec.CommandContext(ctx, binary, arguments...)
	command.Env = os.Environ()
	outputBytes, err := command.Output()

	if err != nil {
		log.Error("An error occurred during execution of the requested binary: ", err)
		return "", err
	}

	if err != nil {
		log.Error("There was an error while attempting to run the ", binary, ": ", err)
		return "", err
	}

	return string(outputBytes), nil
}

func generatedQstatOputput() (string, error) {

	entropy := rand.NewSource(time.Now().UnixNano())
	random := rand.New(entropy)

	ji := JobInfo{
		XMLName: xml.Name{
			Local: "job_info",
		},
		PendingJobs: PendingJob{
			JobList: []Job{
				{
					XMLName: xml.Name{
						Local: "job_list",
					},
					State:          "pw",
					StateAttribute: "pending",
					JBJobNumber:    int64(random.Int()),
					JATPriority:    random.Float64(),
					JobName:        "Job-" + strconv.Itoa(random.Int()),
					JobOwner:       "Owner-" + strconv.Itoa(random.Int()),
					Slots:          3,
				},
			},
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
					JobList: []Job{
						{
							XMLName: xml.Name{
								Local: "job_list",
							},
							State:          "r",
							JBJobNumber:    int64(random.Int()),
							JATPriority:    random.Float64(),
							StateAttribute: "running",
							JobName:        "Job-" + strconv.Itoa(random.Int()),
							JobOwner:       "Owner-" + strconv.Itoa(random.Int()),
							Slots:          3,
						},
						{
							XMLName: xml.Name{
								Local: "job_list",
							},
							State:          "r",
							JBJobNumber:    44,
							JATPriority:    random.Float64(),
							StateAttribute: "running",
							JobName:        "validation",
							JobOwner:       "Owner-" + strconv.Itoa(random.Int()),
							Slots:          3,
						},
					},
				},
				{
					XMLName: xml.Name{
						Local: "Queue-List",
					},
					Name: "all.q@testing.second", //Always needs the @ symbol

					Resources: ResourceList{},
					JobList: []Job{
						{
							XMLName: xml.Name{
								Local: "job_list",
							},
							State:          "r",
							StateAttribute: "running",
							JBJobNumber:    1,
							JATPriority:    1,
							JobName:        "Second-Host-Job",
							JobOwner:       "Owner",
							Slots:          14,
						},
					},
				},
			},
		},
	}

	return ji.GetXML()
}
