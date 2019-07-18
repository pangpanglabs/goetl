package etl

import (
	"context"
)

type ETL struct {
	runner          ETLRunner
	beforeTransform []TransformFilter
	afterTransform  []TransformFunc
}

func New(runner ETLRunner) *ETL {
	return &ETL{
		runner: runner,
	}
}

func (e *ETL) BeforeTransform(f TransformFunc) {
	e.beforeTransform = append(e.beforeTransform, newTransformFilter(f))
}

func (e *ETL) AfterTransform(f TransformFunc) {
	e.afterTransform = append(e.afterTransform, f)
}

func (e *ETL) Run(ctx context.Context) error {
	source, err := e.runner.Extract(ctx)
	if err != nil {
		return err
	}

	transform := e.runner.Transform
	for i := range e.beforeTransform {
		transform = e.beforeTransform[len(e.beforeTransform)-1-i](transform)
	}

	target, err := transform(ctx, source)
	if err != nil {
		return err
	}

	for i := range e.afterTransform {
		var err error
		target, err = e.afterTransform[i](ctx, target)
		if err != nil {
			return err
		}
	}

	if err := e.runner.Load(ctx, target); err != nil {
		return err
	}

	return nil
}
