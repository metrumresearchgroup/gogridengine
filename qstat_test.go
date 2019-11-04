package gogridengine

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCLIModeFailureGetQstatOutput(t *testing.T) {
	//Force to run output and fail
	os.Setenv(environmentPrefix+"TEST", "false")

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "failure operation",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetQstatOutput(make(map[string]string))
			if (err != nil) != tt.wantErr {
				t.Errorf("getQstatOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getQstatOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Just testing to make sure that it doesn't generate unexpected errors.
func TestGeneratedOutputGenerateQState(t *testing.T) {
	//Force to run output and fail
	os.Setenv(environmentPrefix+"TEST", "true")

	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{
			name:    "generated output operation",
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetQstatOutput(make(map[string]string))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQstatOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDeleteQueuedJobByID(t *testing.T) {
	fakeBinary("qdel")
	updatePathWithCurrentDir()
	originalValue := os.Getenv(environmentPrefix + "TEST")

	type args struct {
		jobs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Execution",
			args: args{
				jobs: []string{
					"1",
					"2",
				},
			},
			wantErr: false,
		},
		{
			name: "Test Mode. Should pass",
			args: args{
				jobs: []string{
					"1",
					"2",
				},
			},
			wantErr: false,
		},
	}
	for k, tt := range tests {
		if (k+1)%2 == 0 {
			os.Setenv(environmentPrefix+"TEST", "true")
		} else {
			os.Unsetenv(environmentPrefix + "TEST")
		}
		t.Run(tt.name, func(t *testing.T) {
			if _, err := DeleteQueuedJobByID(tt.args.jobs); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQueuedJobByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	os.Setenv(environmentPrefix+"TEST", originalValue)
	purgeBinary("qdel")
}

//Create an executable qstat file that will exit ok. Will also print the raw input just so we can verify it
func fakeBinary(name string) {
	contents := `#!/bin/bash
	echo $0 $@
	exit 0`

	err := ioutil.WriteFile(name, []byte(contents), 0755)
	if err != nil {
		log.Error("Unable to create the file", err)
	}
}

func purgeBinary(name string) {
	err := os.Remove(name)

	if err != nil {
		log.Error("Unable to create the file", err)
	}
}

func TestDeleteQueuedJobByUsernames(t *testing.T) {
	fakeBinary("qdel")
	type args struct {
		usernames []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful activation",
			args: args{
				usernames: []string{
					"darrellb",
					"dbreeden",
				},
			},
			wantErr: false,
		},
		{
			name: "Test Mode",
			args: args{
				usernames: []string{
					"darrellb",
					"dbreeden",
				},
			},
			wantErr: false,
		},
	}
	for k, tt := range tests {
		if (k+1)%2 == 0 {
			os.Setenv(environmentPrefix+"TEST", "true")
		} else {
			os.Unsetenv(environmentPrefix + "TEST")
		}
		t.Run(tt.name, func(t *testing.T) {
			if _, err := DeleteQueuedJobByUsernames(tt.args.usernames); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQueuedJobByUsernames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	purgeBinary("qdel")
}

func getCurrentPath() string {
	return os.Getenv("PATH")
}

func updatePathWithCurrentDir() {
	os.Setenv("PATH", getCurrentPath()+":.")
}

func TestQSTATWithFakeBinary(t *testing.T) {
	fakeBinary("qstat")
	updatePathWithCurrentDir()

	os.Unsetenv("TEST")
	output, err := GetQstatOutput(make(map[string]string))

	assert.NotEmpty(t, output)
	assert.Nil(t, err)

	purgeBinary("qstat")
	os.Setenv("TEST", "true")
}
