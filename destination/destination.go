package destination

import (
	"github.com/ormushq/ormus/event"
)

// Repo TODO: do we need method for connect to layer 11 here?
type Repo interface {
	GetEvent()
	EventHandler() (event.CoreEvent, error)

	UpdateEventRepo()
}

type Broker interface{}

type Config struct {
	CacheTTL uint
}

type Service struct {
	config Config
	repo   Repo
	broker Broker
}

func (s Service) GetEvent() {
	// TODO: lessen for event then send it to event handler
}

func (s Service) EventHandler(event event.CoreEvent) (event.CoreEvent, error) {
	// TODO: send event.WriteKey to manager and get third party destination configuration if WriteKey dose not exists in our cache db
	// TODO: check our cache if event.MessageID already exists then
	// TODO: send event to the worker pub/sub channel
	return event, nil
	// TODO: This poor method do a lot, let's break it and give its tasks to other methods
}

// UpdateEventRepo TODO: what this method going to do?
func (s Service) UpdateEventRepo() {}
