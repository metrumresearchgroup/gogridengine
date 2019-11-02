package filters

import (
	"testing"

	"github.com/metrumresearchgroup/gogridengine"
	"github.com/stretchr/testify/assert"
)

func TestNewUsernameFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			JobOwner: "Bob",
		},
		{
			JobOwner: "Cindy",
		},
	}

	jl = jl.Filter(NewUsernameFilter("Cindy"))
	assert.Equal(t, 1, len(jl))

	jl = jl.Filter(NewUsernameFilter("bobby"))

	assert.Empty(t, jl)
}

func TestNewLooseStateFilter(t *testing.T) {
	jl := gogridengine.JobList{
		{
			State: "r",
		},
		{
			State: "e",
		},
		{
			State: "ew",
		},
		{
			State: "qw",
		},
		{
			State: "ce",
		},
	}

	//Test the loose filter first
	r1 := jl.Filter(NewLooseStateFilter("r"))

	assert.NotEmpty(t, r1)
	assert.Len(t, r1, 1)

	r2 := jl.Filter(NewLooseStateFilter("w"))

	assert.NotEmpty(t, r2)
	assert.Len(t, r2, 2)

	r3 := jl.Filter(NewLooseStateFilter("e"))

	assert.NotEmpty(t, r3)
	assert.Len(t, r3, 3)

	//Test for Chained Loose Filter
	r4 := jl.
		Filter(NewLooseStateFilter("e")).
		Filter(NewLooseStateFilter("w"))

	assert.NotEmpty(t, r4)
	assert.Len(t, r4, 1)
	assert.Equal(t, r4[0].State, "ew")
}

func TestNewStrictStateFilter(t *testing.T) {

	jl := gogridengine.JobList{
		{
			State: "r",
		},
		{
			State: "e",
		},
		{
			State: "ew",
		},
		{
			State: "qw",
		},
		{
			State: "ce",
		},
		{
			State: "r",
		},
	}

	r1 := jl.Filter(NewStrictStateFilter("r"))

	assert.NotEmpty(t, r1)
	assert.Len(t, r1, 2)

	r2 := jl.
		Filter(NewLooseStateFilter("e")).
		Filter(NewStrictStateFilter("ew"))

	assert.NotEmpty(t, r2)
	assert.Len(t, r2, 1)

}
