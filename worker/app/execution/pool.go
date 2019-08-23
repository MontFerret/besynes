package execution

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type (
	Worker interface {
		Job() Job
		IsRunning() bool
		Process() ([]byte, error)
		Interrupt()
	}

	WorkerFactory func(job Job) Worker

	WorkerPool struct {
		logger  zerolog.Logger
		factory WorkerFactory
		size    int
		workers map[string]Worker
		pool    chan bool
		results chan Result
	}
)

func NewWorkerPool(
	size int,
	logger zerolog.Logger,
	factory WorkerFactory,
) (*WorkerPool, error) {
	if factory == nil {
		return nil, errors.New("missed worker factory")
	}

	wp := new(WorkerPool)
	wp.logger = logger
	wp.factory = factory
	wp.size = size
	wp.workers = make(map[string]Worker)
	wp.pool = make(chan bool, size)
	wp.results = make(chan Result, size)

	return wp, nil
}

func (wp *WorkerPool) Size() int {
	return wp.size
}

func (wp *WorkerPool) Consume(ctx context.Context, queue <-chan Job) <-chan Result {
	out := make(chan Result, wp.size)

	go func() {
		stop := func() {
			// stop all running workers
			for _, w := range wp.workers {
				w.Interrupt()
			}

			// drop all pending results
			for range wp.results {
			}

			// and reset capacity
			for range wp.pool {

			}

			close(out)
		}

		for {
			select {
			case <-ctx.Done():
				stop()

				return
			case job, closed := <-queue:
				if closed {
					// do not terminate the workers, let them finish
					return
				}

				// this is a sync strategy
				// acquiring a gorouting lock by sending a message to the channel
				// its capacity is the amount of available goroutings
				// if the amount is 0 the operation will block until the capacity increases
				wp.pool <- true

				worker := wp.factory(job)
				wp.workers[job.ID] = worker

				// start a new gorouting with the worker
				go func() {
					// run the worker
					out, err := worker.Process()

					jr := Result{
						State: State{
							JobID:     job.ID,
							QueryID:   job.Query.ID,
							Timestamp: time.Now(),
							Status:    StatusCompleted,
							Error:     err,
						},
						Data: out,
					}

					if err != nil {
						jr.Status = StatusErrored
					}

					// sending data results
					wp.results <- jr

					// releasing the gorouting
					<-wp.pool
				}()
			case result := <-wp.results:
				worker, found := wp.workers[result.JobID]

				if !found {
					break
				}

				if worker.IsRunning() {
					worker.Interrupt()

					if result.QueryID == "" {
						result.QueryID = worker.Job().Query.ID
					}
				}

				delete(wp.workers, result.JobID)

				// in case of operation was terminated and the channel is closed
				select {
				case <-ctx.Done():
					stop()

					return
				default:
					out <- result
					break
				}
			}
		}
	}()

	return out
}

func (wp *WorkerPool) Cancel(_ context.Context, jobID string) {
	wp.results <- Result{
		State: State{
			JobID:  jobID,
			Status: StatusCancelled,
		},
	}
}
