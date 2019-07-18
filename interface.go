package goetl

import "context"

type ETLRunner interface {
	Extract(ctx context.Context) (interface{}, error)
	Transform(ctx context.Context, target interface{}) (interface{}, error)
	Load(ctx context.Context, target interface{}) error
}
