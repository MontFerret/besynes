package driver

import (
	"context"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/mafredri/cdp/rpcc"
	"github.com/mafredri/cdp"
)

type Driver struct {
	
}

func New(debuggingURL string) drivers.Driver  {
	// Initiate a new RPC connection to the Chrome DevTools Protocol target.
	conn, err := rpcc.DialContext(ctx, debuggingURL)

	if err != nil {
		return err
	}

	defer conn.Close() // Leaving connections open will leak memory.

	c := cdp.NewClient(conn)
}

func (d *Driver) Close() error {
	panic("implement me")
}

func (d *Driver) Name() string {
	panic("implement me")
}

func (d *Driver) Open(ctx context.Context, params drivers.Params) (drivers.HTMLPage, error) {
	panic("implement me")
}