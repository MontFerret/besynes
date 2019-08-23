package messaging

import (
	"context"
	"encoding/json"
	"github.com/MontFerret/besynes/worker/app/execution"
	zmq "github.com/pebbe/zmq4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Publisher struct {
	logger zerolog.Logger
	addr   string
}

func NewPublisher(logger zerolog.Logger, addr string) *Publisher {
	return &Publisher{
		logger: logger,
		addr:   addr,
	}
}

func (pub *Publisher) Consume(ctx context.Context, input <-chan execution.Result) <-chan error {
	onError := make(chan error, 50)

	go func() {
		stop := func() {
			close(onError)
		}

		zctx, err := zmq.NewContext()

		if err != nil {
			onError <- errors.Wrap(err, "create new pub context")

			stop()

			return
		}

		socket, err := zctx.NewSocket(zmq.PUB)

		if err != nil {
			onError <- errors.Wrap(err, "create new pub socket")

			stop()

			return
		}

		if err := socket.Bind(pub.addr); err != nil {
			onError <- errors.Wrap(err, "bind pub socket")

			stop()

			return
		}

		select {
		case <-ctx.Done():
			if err := socket.Close(); err != nil {
				onError <- errors.Wrap(err, "close pub socket")
			}

			stop()
		case data := <-input:
			b, err := json.Marshal(data)

			if err != nil {
				onError <- errors.Wrap(err, "marshal data")
			}

			_, err = socket.SendBytes(b, zmq.DONTWAIT)

			if err != nil {
				onError <- errors.Wrap(err, "send data")
			}
		}
	}()

	return onError
}
