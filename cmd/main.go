package main

import (
	"github.com/metrumresearchgroup/gogridengine/cmd/jobs"
	log "github.com/sirupsen/logrus"
)

func main() {
	extractor := jobs.JobsCmd()
	if err := extractor.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
