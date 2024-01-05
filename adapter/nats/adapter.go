package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/ormushq/ormus/pkg/richerror"
	"runtime"
	"time"
)

// Config is the configuration for NATS adapter.
type Config struct {
	MaxPending int           `koanf:"max_pending"`
	Timeout    time.Duration `koanf:"timeout"`
	Port       int           `koanf:"port"`
	Password   string        `koanf:"password"`
	User       string        `koanf:"user"`
	Host       string        `koanf:"host"`
}

// Adapter represents the NATS adapter.
type Adapter struct {
	conn   *nats.Conn
	config Config
}

// New creates a new NATS adapter.
func New(config Config) (Adapter, error) {
	const op = "nats.New"
	url := fmt.Sprintf("nats://%s:%s@%s:%d", config.User, config.Password, config.Host, config.Port)
	conn, err := nats.Connect(url)
	if err != nil {
		rErr := richerror.New(op).
			WhitMessage("failed to connect to NATS server:").
			WhitKind(richerror.KindUnexpected).
			WhitWarpError(err)

		return Adapter{}, rErr
	}

	return Adapter{conn: conn, config: config}, nil
}

func (a Adapter) Conn() *nats.Conn {
	return a.conn
}

func (a Adapter) Close() {
	if a.conn != nil {
		a.conn.Close()
	}
}

// Subscribe subscribes to the specified NATS topic and processes incoming events.
func (a Adapter) Subscribe(topic string, action func(msg *nats.Msg)) error {
	const op = "nats.Subscribe"

	sub, err := a.conn.Subscribe(topic, action)
	if err != nil {
		rErr := richerror.New(op).
			WhitMessage("Error subscribing to NATS subject").
			WhitKind(richerror.KindUnexpected).
			WhitWarpError(err)

		return rErr
	}

	// TODO - here we should use ormus logger
	fmt.Println("Subscribed to NATS subject")

	defer func() {
		if err := sub.Unsubscribe(); err != nil {
			// TODO - Handle the error when unsubscribing and add a logger
		}
	}()

	runtime.Goexit()
	return nil
}
