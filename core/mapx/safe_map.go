package mapx

import (
	"context"
	"sync"
	"time"
)

const (
	copyThresholdDefault   = 1000
	maxDeletionDefault     = 10000
	defaultCleanupDuration = time.Minute
)

// Entry represents a key-value pair with expiration time.
type (
	Entry struct {
		Value      interface{}
		Expiration time.Time
	}

	SafeMapOption func(opt *SafeMapWithExpirationOption)
)

func (e Entry) expire() bool {
	if !e.Expiration.IsZero() && e.Expiration.Before(time.Now()) {
		return true
	}
	return false
}

// SafeMapWithExpiration provides a map alternative to avoid memory leak
// and supports key-value pairs with expiration time.
type SafeMapWithExpiration struct {
	lock        sync.RWMutex
	deletionOld int
	deletionNew int
	dirtyOld    map[interface{}]Entry
	dirtyNew    map[interface{}]Entry
	opt         *SafeMapWithExpirationOption
}

type SafeMapWithExpirationOption struct {
	ctx             context.Context
	cleanUpDuration time.Duration
	copyThreshold   int
	maxDeletion     int
}

// NewSafeMapWithExpiration returns a SafeMapWithExpiration.
func NewSafeMapWithExpiration(opts ...SafeMapOption) *SafeMapWithExpiration {
	o := &SafeMapWithExpirationOption{}
	for _, opt := range opts {
		opt(o)
	}
	return &SafeMapWithExpiration{
		dirtyOld: make(map[interface{}]Entry),
		dirtyNew: make(map[interface{}]Entry),
		opt:      o,
	}
}
func (s *SafeMapWithExpirationOption) init() {
	if s.cleanUpDuration <= 0 {
		s.cleanUpDuration = defaultCleanupDuration
	}
	if s.copyThreshold <= 0 {
		s.copyThreshold = copyThresholdDefault
	}
	if s.maxDeletion <= 0 {
		s.maxDeletion = maxDeletionDefault
	}
}

func WithCopyThreshold(val int) SafeMapOption {
	return func(opt *SafeMapWithExpirationOption) {
		opt.copyThreshold = val
	}
}

func WithMaxDeletion(val int) SafeMapOption {
	return func(opt *SafeMapWithExpirationOption) {
		opt.maxDeletion = val
	}
}

func WithCleanUpDefaultTime(t time.Duration) SafeMapOption {
	return func(opt *SafeMapWithExpirationOption) {
		opt.cleanUpDuration = t
	}
}

// Del deletes the value with the given key from m.
func (s *SafeMapWithExpiration) Del(key interface{}) {
	s.lock.Lock()
	if _, ok := s.dirtyOld[key]; ok {
		delete(s.dirtyOld, key)
		s.deletionOld++
	} else if _, ok := s.dirtyNew[key]; ok {
		delete(s.dirtyNew, key)
		s.deletionNew++
	}
	if s.deletionOld >= s.opt.maxDeletion && len(s.dirtyOld) < s.opt.copyThreshold {
		for k, v := range s.dirtyOld {
			s.dirtyNew[k] = v
		}
		s.dirtyOld = s.dirtyNew
		s.deletionOld = s.deletionNew
		s.dirtyNew = make(map[interface{}]Entry)
		s.deletionNew = 0
	}
	if s.deletionNew >= s.opt.maxDeletion && len(s.dirtyNew) < s.opt.copyThreshold {
		for k, v := range s.dirtyNew {
			s.dirtyOld[k] = v
		}
		s.dirtyNew = make(map[interface{}]Entry)
		s.deletionNew = 0
	}
	s.lock.Unlock()
}

// Get gets the value with the given key from m.
func (s *SafeMapWithExpiration) Get(key interface{}) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if val, ok := s.dirtyOld[key]; ok {
		if val.expire() {
			delete(s.dirtyOld, key)
			s.deletionOld++
			return nil, false
		}
		return val, true
	}

	val, ok := s.dirtyNew[key]
	if ok {
		if val.expire() {
			delete(s.dirtyNew, key)
			s.deletionNew++
			return nil, false
		}
		return val, ok
	}
	return nil, false
}

func (s *SafeMapWithExpiration) isExpire(t time.Time) bool {
	if !t.IsZero() && t.Before(time.Now()) {
		return true
	}
	return false
}

// Set sets the value into m with the given key and expiration time.
func (s *SafeMapWithExpiration) Set(key, value interface{}, expiration time.Time) {
	s.lock.Lock()
	if s.deletionOld <= s.opt.maxDeletion {
		if _, ok := s.dirtyNew[key]; ok {
			delete(s.dirtyNew, key)
			s.deletionNew++
		}
		s.dirtyOld[key] = Entry{
			Value:      value,
			Expiration: expiration,
		}
	} else {
		if _, ok := s.dirtyOld[key]; ok {
			delete(s.dirtyOld, key)
			s.deletionOld++
		}
		s.dirtyNew[key] = Entry{
			Value:      value,
			Expiration: expiration,
		}
	}
	s.lock.Unlock()
}

// Size returns the size of m.
func (s *SafeMapWithExpiration) Size() int {
	s.lock.RLock()
	size := len(s.dirtyOld) + len(s.dirtyNew)
	s.lock.RUnlock()
	return size
}

// CleanUp periodically removes expired entries from the map.
func (s *SafeMapWithExpiration) CleanUp() {
	for {
		time.Sleep(s.opt.cleanUpDuration) // Adjust the interval as needed
		s.lock.Lock()
		for key, entry := range s.dirtyOld {
			if s.isExpire(entry.Expiration) {
				delete(s.dirtyOld, key)
				s.deletionOld++
			}
		}
		for key, entry := range s.dirtyNew {
			if s.isExpire(entry.Expiration) {
				delete(s.dirtyNew, key)
				s.deletionNew++
			}
		}
		s.lock.Unlock()
	}
}

// StartCleanUpRoutine starts a goroutine to perform regular clean-up.
func (s *SafeMapWithExpiration) StartCleanUpRoutine() {
	go s.CleanUp()
}
