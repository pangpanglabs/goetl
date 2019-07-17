package etl

import "context"

type Extractor interface {
	Extract(ctx context.Context) (interface{}, error)
}

type Transformer interface {
	Transform(ctx context.Context, target interface{}) (interface{}, error)
}

type Loader interface {
	Load(ctx context.Context, target interface{}) error
}
