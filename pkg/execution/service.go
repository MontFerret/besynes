package execution

import (
	"context"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	"github.com/MontFerret/ferret/pkg/drivers/http"
	"github.com/MontFerret/ferret/pkg/runtime"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Service struct {
	logger   zerolog.Logger
	compiler *compiler.Compiler
}

func NewService(
	logger zerolog.Logger,
	compiler *compiler.Compiler,
) (*Service, error) {
	if compiler == nil {
		return nil, errors.New("missed compiler")
	}

	s := new(Service)
	s.logger = logger
	s.compiler = compiler

	return s, nil
}

func (svc *Service) Execute(ctx context.Context, query Query) ([]byte, error) {
	ctx = drivers.WithContext(ctx, http.NewDriver(), drivers.AsDefault())
	ctx = drivers.WithContext(ctx, cdp.NewDriver(cdp.WithAddress(query.CDPAddress)))

	program, err := svc.compiler.Compile(query.Text)

	if err != nil {
		return nil, errors.Wrap(err, "compile query")
	}

	params := make(map[string]interface{}, len(query.Params))

	for k, v := range query.Params {
		params[k] = v
	}

	out, err := program.Run(
		ctx,
		runtime.WithLog(svc.logger),
		runtime.WithParams(params),
	)

	if err != nil {
		return nil, err
	}

	return out, nil
}
