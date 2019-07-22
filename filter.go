package goetl

import "context"

type BeforeFilterFunc func(ctx context.Context, target interface{}) (interface{}, error)
type AfterFilterFunc func(ctx context.Context, target interface{}) error
type BeforeFilter func(next BeforeFilterFunc) BeforeFilterFunc

func beforeFilter(f BeforeFilterFunc) BeforeFilter {
	return func(next BeforeFilterFunc) BeforeFilterFunc {
		return func(ctx context.Context, target interface{}) (interface{}, error) {
			v, err := f(ctx, target)
			if err != nil {
				return nil, err
			}
			return next(ctx, v)
		}
	}
}
