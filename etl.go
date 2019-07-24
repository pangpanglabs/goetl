package goetl

import (
	"context"
)

type ETL struct {
	runner        ETLRunner
	beforeFilters []BeforeFilter
	afterFilters  []AfterFilterFunc
}

func NewEtl(runner ETLRunner) *ETL {
	return &ETL{
		runner: runner,
	}
}

func (e *ETL) Before(f BeforeFilterFunc) {
	e.beforeFilters = append(e.beforeFilters, beforeFilter(f))
}

func (e *ETL) After(f AfterFilterFunc) {
	e.afterFilters = append(e.afterFilters, f)
}

func (e *ETL) Run(ctx context.Context) error {
	source, err := e.runner.Extract(ctx)
	if err != nil {
		return err
	}

	transform := e.runner.Transform
	for i := range e.beforeFilters {
		transform = e.beforeFilters[len(e.beforeFilters)-1-i](transform)
	}

	target, err := transform(ctx, source)
	if err != nil {
		return err
	}

	for i := range e.afterFilters {
		if err := e.afterFilters[i](ctx, target); err != nil {
			return err
		}
	}

	if err := e.runner.Load(ctx, target); err != nil {
		return err
	}

	return nil
}
