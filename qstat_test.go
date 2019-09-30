package gogridengine

import (
	"io/ioutil"
	"os"
	"testing"
)

const qstatPath string = "/usr/local/bin/"

func TestCLIModeFailureGetQstatOutput(t *testing.T) {
	//Force to run output and fail
	os.Setenv("TEST", "false")

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
			got, err := GetQstatOutput()
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
	os.Setenv("TEST", "true")

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
			_, err := GetQstatOutput()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQstatOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDeleteQueuedJobByID(t *testing.T) {
	fakeBinary("qdel")
	originalValue := os.Getenv("TEST")

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
			os.Setenv("TEST", "true")
		} else {
			os.Unsetenv("TEST")
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteQueuedJobByID(tt.args.jobs); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQueuedJobByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	os.Setenv("TEST", originalValue)
	purgeBinary("qdel")
}

//Create an executable qstat file that will exit ok
func fakeBinary(name string) {
	contents := `#!/bin/bash
	exit 0`

	ioutil.WriteFile(qstatPath+name, []byte(contents), 0755)
}

func purgeBinary(name string) {
	os.Remove(qstatPath + name)
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
			os.Setenv("TEST", "true")
		} else {
			os.Unsetenv("TEST")
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteQueuedJobByUsernames(tt.args.usernames); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQueuedJobByUsernames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	purgeBinary("qdel")
}
