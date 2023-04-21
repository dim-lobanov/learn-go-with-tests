package syncwaitgroup

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// You can write just sync.Mutex in struct -> with this you can lock on like c.Lock() but it makes mutex public

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

// Use channels when passing ownership of data
// Use mutexes for managing state
