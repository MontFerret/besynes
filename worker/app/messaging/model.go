package messaging

import "github.com/MontFerret/besynes/worker/app/execution"

type QueryStream struct {
	Error error
	Data  execution.Query
}
