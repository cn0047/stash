package main

import (
	"sync"
	"time"
)

type ErrorsBatch struct {
	batch []error
	ok    bool
	sync.Mutex
}

func NewErrorsBatch(size int) *ErrorsBatch {
	return &ErrorsBatch{batch: make([]error, 0, size)}
}

func (e *ErrorsBatch) ReadFromChan(errs <-chan error) {
	go func() {
		defer func() {
			e.Lock()
			e.ok = true
			e.Unlock()
		}()

		for {
			select {
			case err, ok := <-errs:
				if !ok {
					return
				}
				e.batch = append(e.batch, err)
			default:
				time.Sleep(1)
			}
		}
	}()
}

func (e *ErrorsBatch) GetErrorsBatch() []error {
	for {
		e.Lock()
		v := e.ok
		e.Unlock()

		if v == true {
			break
		}
	}

	return e.batch
}
