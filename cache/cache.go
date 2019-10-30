package cache

import (
	"context"
	"sync"
	"time"

	"github.com/metrumresearchgroup/gogridengine"
	log "github.com/sirupsen/logrus"
)

type empty struct{}

//Manager is an interface that should be used for interfacing with the cache itself
type Manager interface {
	Update() gogridengine.JobInfo
	Get() gogridengine.JobInfo
	Initialize(ttl time.Duration)
	Stop()
}

//Read contains the communication channels and the context for interacting with the managing goroutine
type Read struct {
	Request  chan empty
	Response chan gogridengine.JobInfo
	Context  context.Context
	Cancel   context.CancelFunc
}

//Write contains the communication channel and the context for interacting with the managing goroutine
type Write struct {
	Request chan gogridengine.JobInfo
	Context context.Context
	Cancel  context.CancelFunc
}

//Update contains the communication channels and the context for interacting with the managing goroutine
type Update struct {
	Request  chan empty
	Response chan gogridengine.JobInfo
	Context  context.Context
	Cancel   context.CancelFunc
}

//Poll is a struct used for interacting with the cache poller
type Poll struct {
	Context context.Context
	Cancel  context.CancelFunc
}

//Cache is the literal cache contents, its mutex, and its channels for manipulation.
type Cache struct {
	contents gogridengine.JobInfo //Not externally accessible
	Mutex    *sync.Mutex
	Read     Read
	update   Update
	write    Write
	poll     Poll
	Context  context.Context
}

//Initialize prepares the cache, setups up the managing goroutines and builds the channels
func (c Cache) Initialize(ttl time.Duration) {

	//First Build the Communication structs
	c.write = Write{
		Request: make(chan gogridengine.JobInfo),
		Context: context.Background(),
	}

	c.write.Context, c.write.Cancel = context.WithCancel(c.write.Context)

	c.Read = Read{
		Request:  make(chan empty),
		Response: make(chan gogridengine.JobInfo),
		Context:  context.Background(),
	}

	c.Read.Context, c.Read.Cancel = context.WithCancel(c.Read.Context)

	c.update = Update{
		Request:  make(chan empty),
		Response: make(chan gogridengine.JobInfo),
		Context:  context.Background(),
	}

	c.update.Context, c.update.Cancel = context.WithCancel(c.update.Context)

	c.poll = Poll{
		Context: context.Background(),
	}

	c.poll.Context, c.poll.Cancel = context.WithCancel(c.poll.Context)

	//Spawn the managing goroutines. All will take teardown

	//Read
	go func(read Read) {
		for {
			select {
			case <-read.Request:
				read.Response <- c.contents
			case <-read.Context.Done():
				close(read.Request)
				close(read.Request)
				return
			}
		}
	}(c.Read)

	//Write
	go func(write Write) {
		for {
			select {
			case info := <-write.Request:
				c.Mutex.Lock()
				c.contents = info
				c.Mutex.Unlock()
			case <-write.Context.Done():
				close(write.Request)
				return
			}
		}
	}(c.write)

	//Update
	go func(update Update) {
		for {
			select {
			case <-update.Request:
				//We need to get the JobInfo contents from
				ji, err := gogridengine.GetJobInfo()

				if err != nil {
					log.Error("We experienced issues retrieving or deserializing the qstat XML: ", err)
					continue //Don't break or exit. Just keep looping.
				}

				update.Response <- ji
			case <-update.Context.Done():
				close(update.Request)
				close(update.Response)
				return
			}
		}
	}(c.update)

	//Build Goroutine to handle sleeping and repopulating cache
	go func(write Write, poll Poll) {
		for {
			//Always running, waiting for teardown
			ji, err := gogridengine.GetJobInfo()

			if err != nil {
				log.Error("We experienced issues retrieving or deserializing the qstat XML: ", err)
				continue //Don't break or exit. Just keep looping.
			}

			write.Request <- ji

			//Look for termination
			select {
			case <-poll.Context.Done():
				return
			default:
				//Do nothing
			}

			time.Sleep(ttl)
		}
	}(c.write, c.poll)

}

//Get is the mechanism by which external parties access the cache
func (c Cache) Get() gogridengine.JobInfo {
	return c.Request(c.Read.Request, c.Read.Response)
}

//Update is a way of requesting requerying of the source data, population of the cache and accessing that data.
func (c Cache) Update() gogridengine.JobInfo {
	return c.Request(c.update.Request, c.update.Response)
}

//Request is a uniform method for performing similar workloads dealing with request / response channels
func (c Cache) Request(request chan empty, response chan gogridengine.JobInfo) gogridengine.JobInfo {
	e := empty{}

	request <- e

	ctx := context.Background()
	ctx, err := context.WithTimeout(ctx, 3*time.Second)

	if err != nil {
		log.Error("Unable to create deadlined context for some reason")
		//Fallback to normal background context
		ctx = context.Background()
	}

	for {
		select {
		case ji := <-response:
			return ji
		case <-ctx.Done():
			log.Error("Request to cache has failed")
			return gogridengine.JobInfo{}
		}
	}
}

//Stop issues cancellation requests to the communication contexts attached to all structs used for interfacing with the cache.
func (c Cache) Stop() {
	c.Read.Cancel()
	c.write.Cancel()
	c.update.Cancel()
	c.poll.Cancel()
}
