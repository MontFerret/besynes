package app

import (
	"context"
	"fmt"
	"github.com/MontFerret/besynes/worker/app/execution"
	"github.com/MontFerret/besynes/worker/app/messaging"
	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/pkg/errors"
	"os"

	"github.com/rs/zerolog"
)

type Application struct {
	settings   Settings
	logger     zerolog.Logger
	executor   *execution.Service
	subscriber *messaging.Subscriber
	publisher  *messaging.Publisher
}

func New(settings Settings) (*Application, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	exec, err := execution.NewService(
		execution.NewDefaultSettings(),
		logger,
		compiler.New(),
	)

	if err != nil {
		return nil, errors.Wrap(err, "create execution service")
	}

	app := new(Application)
	app.settings = settings
	app.logger = logger
	app.executor = exec
	app.subscriber = messaging.NewSubscriber(logger, fmt.Sprintf("tcp://localhost:%d", settings.SubPort))
	app.publisher = messaging.NewPublisher(logger, fmt.Sprintf("tcp://*:%d", settings.PubPort))

	return app, nil
}

func (app *Application) Run(ctx context.Context) {
	childCtx, fn := context.WithCancel(ctx)
	onQuery, onSubErr := app.subscriber.Produce(childCtx)
	onPubErr := app.publisher.Consume(childCtx, app.executor.Consume(childCtx, onQuery, nil))
	stop := func() {
		fn()
	}

	for {
		select {
		case <-ctx.Done():
			return
		case err, closed := <-onSubErr:
			app.logger.Err(err).Msg("sub error")

			if closed {
				stop()
				return
			}
		case err, closed := <-onPubErr:
			app.logger.Err(err).Msg("pub error")

			if closed {
				stop()
				return
			}
		}
	}
}
