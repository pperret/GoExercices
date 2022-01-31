// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

// entry is a cache item
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result   // the client wants a single result
	done     <-chan struct{} // cancel channel
}

// Memo is the cache manager based on a monitor
type Memo struct{ requests chan request }

// New returns a memorization of f.  Clients must subsequently call Close.
func New(fct Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	// Start the monitor
	go memo.server(fct)
	return memo
}

// Get fetchs a resource by sending a request to the monitor and waiting for its reply
func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	return res.value, res.err
}

// Close releases resources used by a cache manager instance
func (memo *Memo) Close() { close(memo.requests) }

// Monitor manages requests
func (memo *Memo) server(fct Func) {
	cache := make(map[string]*entry)
	pending := make(map[string]*entry)
	var mapLock sync.Mutex
	for req := range memo.requests {
		mapLock.Lock()
		e := cache[req.key]
		if e == nil {
			e = pending[req.key]
			if e == nil {
				// This is the first request for this key.
				e = &entry{ready: make(chan struct{})}
				pending[req.key] = e
				go e.call(fct, req.key, req.done, cache, pending, &mapLock) // call f(key)
			}
		}
		mapLock.Unlock()
		go e.deliver(req.response)
	}
}

// call calls the fetch function
func (e *entry) call(fct Func, key string, done <-chan struct{}, cache, pending map[string]*entry, mapLock *sync.Mutex) {
	// Evaluate the function.
	e.res.value, e.res.err = fct(key, done)
	mapLock.Lock()
	select {
	case <-done:
		// Canceled
	default:
		// cache contains only successful results
		if e.res.err == nil {
			cache[key] = e
		}
	}
	delete(pending, key)
	mapLock.Unlock()

	// Broadcast the ready condition.
	close(e.ready)
}

// deliver sends the reply to the caller
func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
