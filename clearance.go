package goetl

import (
	"context"
)

type Clearance struct {
	runner ClearanceRunner
}

func NewClearance(runner ClearanceRunner) *Clearance {
	return &Clearance{
		runner: runner,
	}
}

func (c *Clearance) Run(ctx context.Context) error {
	source, err := c.runner.Read(ctx)
	if err != nil {
		return err
	}

	target, err := c.runner.CompareWithSourceAndTarget(ctx, source)
	if err != nil {
		return err
	}

	if err := c.runner.Save(ctx, target); err != nil {
		return err
	}

	return nil
}
