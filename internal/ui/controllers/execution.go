package controllers

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/internal/ui/models"
	"github.com/MontFerret/besynes/pkg/execution"
)

type Execution struct {
	logger  zerolog.Logger
	service *execution.Service
}

func NewExecution(bridge *bridges.Execution, logger zerolog.Logger, service *execution.Service) *Execution {
	ctl := &Execution{
		logger:  logger,
		service: service,
	}

	bridge.ConnectExecute(ctl.execute)

	return ctl
}

func (ctl *Execution) execute(query string, addr string) *models.Result {
	data, err := ctl.service.Execute(context.Background(), execution.Query{
		Text:       query,
		Params:     nil,
		CDPAddress: "127.0.0.1:9222",
	})

	result := models.NewResult(nil)

	if err != nil {
		result.SetError(err.Error())
	} else {
		result.SetData(string(data))
	}

	return result
}
