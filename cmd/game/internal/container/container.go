package container

import ()

type container struct {
	closer closer
}

func newContainer(closer closer) *container {
	return &container{
		closer: closer,
	}
}
