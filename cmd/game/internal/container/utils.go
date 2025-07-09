package container

import (
	"reflect"
)

type (
	resouce interface {
		Free()
	}
	resourceQueueCloser interface {
		EnqueueFree(fn func())
	}
	resourcesDict map[uintptr]any
)

var knownResources = resourcesDict{}

func resolver[T any](c resourceQueueCloser, factory func() *T) *T {
	ptr := reflect.ValueOf(factory).Pointer()

	if target, exist := knownResources[ptr]; exist {
		return target.(*T)
	}

	knownResources[ptr] = factory()

	c.EnqueueFree(func() {
		if res, ok := knownResources[ptr].(resouce); ok {
			res.Free()
		}
		delete(knownResources, ptr)
	})

	return knownResources[ptr].(*T)
}

func static[T any](c *container, factory func() *T) *T {
	return resolver(c.closer, factory)
}
