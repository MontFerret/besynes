package messaging

import (
	"context"
	"encoding/json"

	"github.com/MontFerret/besynes/worker/app/execution"
	zmq "github.com/pebbe/zmq4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Subscriber struct {
	logger zerolog.Logger
	addr   string
}

func NewSubscriber(logger zerolog.Logger, addr string) *Subscriber {
	return &Subscriber{
		logger: logger,
		addr:   addr,
	}
}

// run the subscriber
func (sub *Subscriber) Produce(ctx context.Context) (<-chan execution.Query, <-chan error) {
	onData := make(chan execution.Query, 50)
	onError := make(chan error, 50)

	go func() {
		stop := func() {
			close(onData)
			close(onError)
		}

		zctx, err := zmq.NewContext()

		if err != nil {
			onError <- errors.Wrap(err, "create new sub context")

			stop()

			return
		}

		socket, err := zctx.NewSocket(zmq.SUB)

		if err != nil {
			onError <- errors.Wrap(err, "create new sub socket")

			stop()

			return
		}

		if err := socket.Connect(sub.addr); err != nil {
			onError <- errors.Wrap(err, "connect socket")

			stop()

			return
		}

		for {
			select {
			case <-ctx.Done():
				if err := socket.Close(); err != nil {
					onError <- errors.Wrap(err, "socket close")
				}

				stop()

				return

			default:
				msg, err := socket.RecvMessageBytes(0)

				if err != nil {
					sub.logger.Err(err).Msg("failed to read sub message")

					onError <- errors.Wrap(err, "read from socket")

					continue
				}

				if len(msg) == 0 {
					sub.logger.Warn().Msg("received an empty message")

					continue
				}

				var query execution.Query

				if err := json.Unmarshal(msg[0], &query); err != nil {
					sub.logger.Err(err).Msg("failed to unmarshal a message")

					onError <- errors.Wrap(err, "unmarshal a message")

					continue
				}

				onData <- query
			}
		}
	}()

	return onData, onError
}
