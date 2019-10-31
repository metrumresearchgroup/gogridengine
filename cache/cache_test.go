package cache

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheInitialize(t *testing.T) {
	//Enable generated content
	os.Setenv("TEST", "true")

	c := Initialize(5 * time.Second)

	output := c.Get()

	assert.NotEmpty(t, output)

	o2 := c.Get()

	//Verify an immediate re-request returns the same value.
	assert.Equal(t, output, o2)

	time.Sleep(1 * time.Second)

	o3 := c.Get()

	//Verify near-term window value is the same
	assert.Equal(t, output, o3)

	time.Sleep(5 * time.Second)

	o4 := c.Get()

	assert.NotEqual(t, output, o4)

}

func TestCacheCancel(t *testing.T) {

	//Enable generated content
	os.Setenv("TEST", "true")

	c := Initialize(1 * time.Second)

	c.Stop()

	time.Sleep(10 * time.Millisecond)

	assert.False(t, c.Active())
}
