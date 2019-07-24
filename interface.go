package goetl

import "context"

type ETLRunner interface {
	Extract(ctx context.Context) (interface{}, error)
	Transform(ctx context.Context, target interface{}) (interface{}, error)
	Load(ctx context.Context, target interface{}) error
}

type ClearanceRunner interface {
	Read(ctx context.Context) (interface{}, error)
	CompareWithSourceAndTarget(ctx context.Context, source interface{}) (interface{}, error)
	Save(ctx context.Context, target interface{}) error
}
