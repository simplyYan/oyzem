package oyzem

import (
	"fmt"
	"sync"
)

type Memoizer struct {
	cache map[string]interface{}
	mu    sync.Mutex
}

func New() *Memoizer {
	return &Memoizer{
		cache: make(map[string]interface{}),
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
	// Verifica se a função passada é válida
	fnType := fmt.Sprintf("%T", fn)
	if fnType != "func(...interface{}) interface {}" {
		return nil, fmt.Errorf("função com assinatura inválida: %s", fnType)
	}

	return func(args ...interface{}) (interface{}, error) {
		key := m.key(args...)
		m.mu.Lock()
		defer m.mu.Unlock()

		if result, ok := m.cache[key]; ok {
			return result, nil
		}

		result := fn.(func(...interface{}) interface{})(args...)
		m.cache[key] = result
		return result, nil
	}, nil
}

func (m *Memoizer) Run(memoizedFn func(...interface{}) (interface{}, error), args ...interface{}) (interface{}, error) {
	return memoizedFn(args...)
}

func (m *Memoizer) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache = make(map[string]interface{})
}
