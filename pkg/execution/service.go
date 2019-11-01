package execution

import (
	"context"
	"time"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	"github.com/MontFerret/ferret/pkg/drivers/http"
	"github.com/MontFerret/ferret/pkg/runtime"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Executor struct {
	logger   zerolog.Logger
	compiler *compiler.Compiler
}

func NewExecutor(
	logger zerolog.Logger,
	compiler *compiler.Compiler,
) (*Executor, error) {
	if compiler == nil {
		return nil, errors.New("missed compiler")
	}

	s := new(Executor)
	s.logger = logger
	s.compiler = compiler

	return s, nil
}

func (svc *Executor) Execute(ctx context.Context, query Query) Result {
	ctx = drivers.WithContext(ctx, http.NewDriver(), drivers.AsDefault())
	ctx = drivers.WithContext(ctx, cdp.NewDriver(cdp.WithAddress(query.CDPAddress)))

	compiletimeStart := time.Now()

	program, err := svc.compiler.Compile(query.Text)

	compiletimeStop := time.Since(compiletimeStart)

	result := Result{
		Data: nil,
		Stats: Statistics{
			Compilation: compiletimeStop,
			Runtime:     time.Duration(0),
		},
	}

	if err != nil {
		result.Error = err

		return result
	}

	params := make(map[string]interface{}, len(query.Params))

	for k, v := range query.Params {
		params[k] = v
	}

	runtimeStart := time.Now()

	out, err := program.Run(
		ctx,
		runtime.WithLog(svc.logger),
		runtime.WithParams(params),
	)

	runtimeStop := time.Since(runtimeStart)

	result.Stats.Runtime = runtimeStop

	if err != nil {
		result.Error = err

		return result
	}

	result.Data = out

	return result
}
