package etl

import "context"

type TransformFunc func(ctx context.Context, target interface{}) (interface{}, error)
type TransformFilter func(next TransformFunc) TransformFunc

func newTransformFilter(f TransformFunc) TransformFilter {
	return func(next TransformFunc) TransformFunc {
		return func(ctx context.Context, target interface{}) (interface{}, error) {
			v, err := f(ctx, target)
			if err != nil {
				return nil, err
			}
			return next(ctx, v)
		}
	}
}
