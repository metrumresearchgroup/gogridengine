package extractor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type QstatExtractor struct {
	Path   string
	Params []string
}

func (e *QstatExtractor) Extract() ([]byte, error) {
	cmd := exec.Command(e.Path, e.Params...)
	cmd.Env = os.Environ()

	buff := new(bytes.Buffer)
	cmd.Stdout = buff

	var err error
	if err = cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to execute underlying command during job get request: %w", err)
	}
	return buff.Bytes(), nil
}

func (e *QstatExtractor) AddArguments(args ...string) {
	e.Params = append(e.Params, args...)
}

func New(qstatPath *string) (*QstatExtractor, error) {
	var err error
	e := &QstatExtractor{
		// Default is to assume qstat is just called from the path
		Path: "qstat",
		Params: []string{
			"-xml",
		},
	}
	if qstatPath != nil {
		if !strings.HasPrefix(*qstatPath, "/") {
			if e.Path, err = exec.LookPath(*qstatPath); err != nil {
				return nil, fmt.Errorf("we couldn't find %s in your path. please verify and try again: %w", *qstatPath, err)
			}
		}
	}

	return e, nil
}
