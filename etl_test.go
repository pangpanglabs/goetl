package etl

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestETL(t *testing.T) {
	etl := buildETL()
	if err := etl.Run(context.Background()); err != nil {
		t.Fail()
	}
}

func buildETL() *ETL {
	etl := New(Mock{})
	etl.BeforeTransform(toUpper)
	etl.BeforeTransform(trim)

	etl.AfterTransform(after1)
	etl.AfterTransform(after2)
	return etl
}

type Mock struct {
}

func (m Mock) Extract(ctx context.Context) (interface{}, error) {
	fmt.Println("Extract")
	return []string{"  hello  ", "  world  "}, nil
}

func (m Mock) Transform(ctx context.Context, target interface{}) (interface{}, error) {
	fmt.Println("Transform")

	fmt.Println(target)
	ss, ok := target.([]string)
	if !ok {
		return nil, errors.New("Invalid type")
	}
	v := map[string]int{}
	for _, s := range ss {
		v[s] = len(s)
	}

	return v, nil
}
func (m Mock) Load(ctx context.Context, target interface{}) error {
	fmt.Println("Load")

	v, ok := target.(map[string]int)
	if !ok {
		return errors.New("Invalid type")
	}
	fmt.Println(v)
	return nil
}

func toUpper(ctx context.Context, target interface{}) (interface{}, error) {
	fmt.Println("BeforeFilter1")
	ss, ok := target.([]string)
	if !ok {
		return nil, errors.New("Invalid type")
	}

	var result []string
	for _, s := range ss {
		result = append(result, strings.ToUpper(s))
	}
	return result, nil
}
func trim(ctx context.Context, target interface{}) (interface{}, error) {
	fmt.Println("BeforeFilter2")
	ss, ok := target.([]string)
	if !ok {
		return nil, errors.New("Invalid type")
	}

	var result []string
	for _, s := range ss {
		result = append(result, strings.TrimSpace(s))
	}
	return result, nil
}

func after1(ctx context.Context, target interface{}) (interface{}, error) {
	fmt.Println("AfterFilter1")
	return target, nil
}
func after2(ctx context.Context, target interface{}) (interface{}, error) {
	fmt.Println("AfterFilter2")
	return target, nil
}
