package etl

import "context"

func ETL(ctx context.Context, e Extractor, t Transformer, l Loader) error {
	source, err := e.Extract(ctx)
	if err != nil {
		return err
	}

	target, err := t.Transform(ctx, source)
	if err != nil {
		return err
	}

	if err := l.Load(ctx, target); err != nil {
		return err
	}

	return nil
}
