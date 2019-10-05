package execution

import (
	"context"
	"runtime"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Service struct {
	logger   zerolog.Logger
	compiler *compiler.Compiler
	pool     *WorkerPool
}

func NewService(
	settings Settings,
	logger zerolog.Logger,
	compiler *compiler.Compiler,
) (*Service, error) {
	if compiler == nil {
		return nil, errors.New("missed compiler")
	}

	s := new(Service)
	s.logger = logger
	s.compiler = compiler

	size := runtime.NumCPU() * settings.PoolSize
	pool, err := NewWorkerPool(size, logger, func(job Job) Worker {
		return NewFQLWorker(
			compiler,
			logger.With().
				Str("job_id", job.ID).
				Str("query_id", job.Query.ID).
				Logger(),
			job,
		)
	})

	if err != nil {
		return nil, err
	}

	s.pool = pool

	return s, nil
}

func (service *Service) Consume(ctx context.Context, process <-chan Job, interrupt <-chan string) <-chan Result {
	onJob := make(chan Job, service.pool.Size())

	go func() {
		stop := func() {
			close(onJob)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case job, closed := <-process:
				if closed {
					stop()

					return
				}

				onJob <- job
			case jobID := <-interrupt:
				service.pool.Cancel(ctx, jobID)
			}
		}
	}()

	return service.pool.Consume(ctx, onJob)
}
