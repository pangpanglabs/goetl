# goetl

A toolkit for ETL(Extract, Transform, Load) in Go

## Getting Started

Reference: https://github.com/pangpanglabs/goetl/blob/master/etl_test.go

```go
import (
	"github.com/pangpanglabs/goetl"
)

func TestETL(t *testing.T) {
	etl := buildETL()

	/*
	    실행 여부 확인
	*/

	if err := etl.Run(context.Background()); err != nil {
		t.Fail()
	}
}

func buildETL() *goetl.ETL {
	etlRunner := Mock{
		// 생성 로직
	}

	etl := goetl.New(etlRunner)

	etl.Before(toUpper)
	etl.Before(trim)

	etl.After(after1)
	etl.After(after2)

	return etl
}
```
