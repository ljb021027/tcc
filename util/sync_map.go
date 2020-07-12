package util

import (
	"sync"
)

type SyncMap struct {
	Cache sync.Map
	mu    sync.Mutex
}

func (s *SyncMap) ComputeIfAbsent(key interface{}, f func(key interface{}) (interface{}, error)) (interface{}, error) {
	value, ok := s.Cache.Load(key)
	if !ok {
		s.mu.Lock()
		defer s.mu.Unlock()
		i, err := f(key)
		if err != nil {
			return nil, err
		}
		value, _ = s.Cache.LoadOrStore(key, i)
	}
	return value, nil
}
