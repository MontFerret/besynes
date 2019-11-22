package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/cloudfoundry/bytefmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"

	"github.com/MontFerret/besynes/pkg/execution"
	"github.com/MontFerret/besynes/pkg/settings"
)

var ErrInvalidParams = errors.New("Invalid parameter values. Use valid JSON object.")

type (
	ExecutionStatistics struct {
		Size        string `json:"size"`
		Runtime     string `json:"runtime"`
		Compilation string `json:"compilation"`
	}

	ExecutionResult struct {
		Data  string              `json:"data"`
		Error string              `json:"error"`
		Stats ExecutionStatistics `json:"stats"`
	}

	Execution struct {
		logger   zerolog.Logger
		settings *settings.Service
		executor *execution.Executor
	}
)

func NewExecution(
	logger zerolog.Logger,
	settingsSvc *settings.Service,
	executor *execution.Executor,
) *Execution {
	return &Execution{
		logger:   logger,
		settings: settingsSvc,
		executor: executor,
	}
}

func (ctl *Execution) Execute(query *core.QJsonObject) (ExecutionResult, error) {
	q, err := ctl.parseQuery(query)

	if err != nil {
		return ExecutionResult{}, err
	}

	out := ctl.executor.Execute(context.Background(), q)

	result := ExecutionResult{
		Data:  ctl.formatJSON(out.Data),
		Error: ctl.formatError(out.Error),
		Stats: ExecutionStatistics{},
	}

	result.Stats.Compilation = ctl.formatDuration(out.Stats.Compilation)
	result.Stats.Runtime = ctl.formatDuration(out.Stats.Runtime)
	result.Stats.Size = ctl.formatSize(len(out.Data))

	return result, nil
}

func (ctl *Execution) parseQuery(query *core.QJsonObject) (execution.Query, error) {
	var text string
	var params map[string]interface{}

	if query.Contains("text") {
		text = query.Value("text").ToString()
	}

	if query.Contains("params") {
		paramsBytes := []byte(query.Value("params").ToString())

		// Make sure that params string is not empty
		if !ctl.isParamsEmpty(paramsBytes) {
			// Check if it's an object
			if !ctl.isParamsObjectValid(paramsBytes) {
				return execution.Query{}, ErrInvalidParams
			}

			params = make(map[string]interface{})

			err := json.Unmarshal(paramsBytes, &params)

			if err != nil {
				return execution.Query{}, ErrInvalidParams
			}
		}
	}

	opts := ctl.settings.Get()

	return execution.Query{
		Text:       text,
		Params:     params,
		CDPAddress: opts.CDPAddress,
	}, nil
}

func (ctl *Execution) formatJSON(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	var b bytes.Buffer

	err := json.Indent(&b, data, "", "  ")

	if err != nil {
		ctl.logger.Error().Err(err).Msg("failed to format output")

		return string(data)
	}

	return b.String()
}

func (ctl *Execution) formatError(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}

func (ctl *Execution) isParamsEmpty(text []byte) bool {
	return len(text) <= 2
}

func (ctl *Execution) isParamsObjectValid(text []byte) bool {
	if len(text) < 2 {
		return false
	}

	return string(text[0:1]) == "{" && string(text[len(text)-1:]) == "}"
}

func (ctl *Execution) formatDuration(d time.Duration) string {
	return d.String()
}

func (ctl *Execution) formatSize(sizeInBytes int) string {
	return bytefmt.ByteSize(uint64(sizeInBytes))
}
