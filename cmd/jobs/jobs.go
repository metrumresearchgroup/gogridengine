package jobs

import (
	"github.com/metrumresearchgroup/gogridengine/extractor"
	"github.com/metrumresearchgroup/gogridengine/repository/sge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func flags(c *cobra.Command) {
	c.Flags().String("path", "qstat", "The path from which we will access qstat")
	c.Flags().String("host", "", "filter jobs for an existing queue host")
	c.Flags().Int("job-id", 0, "An id of a job to filter for")
	c.Flags().String("state", "", "A specific state for which to query jobs on")
	c.Flags().String("user", "", "A specific user for which to query jobs)")
	_ = viper.BindPFlags(c.Flags())
}

func JobsCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "jobs",
		Short: "Extract jobs from SGE",
		RunE: func(cmd *cobra.Command, args []string) error {
			var input struct {
				Path  string `mapstructure:"path"`
				JobId int    `mapstructure:"job-id"`
				User  string `mapstructure:"user"`
				Host  string `mapstructure:"host"`
				State string `mapstructure:"state"`
			}

			var err error
			if err = viper.Unmarshal(&input); err != nil {
				return err
			}

			sgeExtractor, err := extractor.New(&input.Path)
			if err != nil {
				return err
			}

			getRequest := &sge.GetJobsRequest{}

			if input.Host != "" {
				getRequest.Host = &input.Host
			}

			if input.User != "" {
				getRequest.User = &input.User
			}

			if input.JobId != 0 {
				getRequest.ID = &input.JobId
			}

			if input.State != "" {
				var sc sge.StateCode
				sc = sge.StateCode(input.State)
				getRequest.State = &sc
			}

			sgeRepo := sge.New(sgeExtractor)
			output, err := sgeRepo.Get(getRequest)

			if err != nil {
				return err
			}

			for _, v := range output {
				println(v.JBJobNumber)
			}

			return nil
		},
	}
	flags(c)

	return c
}
