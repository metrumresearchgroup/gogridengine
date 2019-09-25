package gogridengine

import (
	"os"
	"testing"
)

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
