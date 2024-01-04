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
	xmlLocation := XMLDataSource{location: "https://raw.githubusercontent.com/metrumresearchgroup/gogridengine/master/test_data/medium.xml"}

	if os.Getenv("GOGRIDENGINE_TEST_SOURCE") != "" {
		xmlLocation.location = os.Getenv("GOGRIDENGINE_TEST_SOURCE")
	}

	xmlResp := &XmlContentReader{resource: &xmlLocation}

	return xmlResp.Read()
}
