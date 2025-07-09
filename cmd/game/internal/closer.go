package internal

import "log"

type Closer struct {
	queue []func()
}

func NewCloser() *Closer {
	return &Closer{
		queue: make([]func(), 0, 32),
	}
}

func (c *Closer) EnqueueFree(fn func()) {
	c.queue = append(c.queue, fn)
}

func (c *Closer) EnqueueClose(fn func() error) {
	c.queue = append(c.queue, func() {
		if err := fn(); err != nil {
			log.Printf("Failed close resouce '#{{n}': #{err}")
		}
	})
}

func (c Closer) close() {
	for i := len(c.queue) - 1; i >= 0; i-- {
		c.queue[i]()
	}
}
