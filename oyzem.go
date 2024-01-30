package oyzem

import (
	"fmt"
	"reflect"
	"sync"
)

type Memoizer struct {
	cache map[string]reflect.Value
	mu    sync.Mutex
}

func New() *Memoizer {
	return &Memoizer{
		cache: make(map[string]reflect.Value),
	}
}

func (m *Memoizer) key(args ...interface{}) string {
	var key string
	for _, arg := range args {
		key += "|" + fmt.Sprintf("%v", arg)
	}
	return key
}

func (m *Memoizer) Memoize(fn interface{}) (func(...interface{}) (interface{}, error), error) {
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return nil, fmt.Errorf("not a function")
	}

	return func(args ...interface{}) (interface{}, error) {
		fnValue := reflect.ValueOf(fn)
		if fnType.NumIn() != len(args) {
			return nil, fmt.Errorf("incorrect number of arguments")
		}

		var inArgs []reflect.Value
		for i, arg := range args {
			argType := fnType.In(i)
			if reflect.TypeOf(arg) != argType {
				return nil, fmt.Errorf("incorrect argument type for parameter %d", i+1)
			}

			inArgs = append(inArgs, reflect.ValueOf(arg))
		}

		key := m.key(args...)
		m.mu.Lock()
		defer m.mu.Unlock()

		if result, ok := m.cache[key]; ok {
			return result.Interface(), nil
		}

		resultValues := fnValue.Call(inArgs)
		result := resultValues[0]
		m.cache[key] = result
		return result.Interface(), nil
	}, nil
}

func (m *Memoizer) Run(memoizedFn func(...interface{}) (interface{}, error), args ...interface{}) (interface{}, error) {
	return memoizedFn(args...)
}

func (m *Memoizer) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache = make(map[string]reflect.Value)
}
