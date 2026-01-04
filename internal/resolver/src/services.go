package resolver

import "io"

type Resolver interface {
	Resolve(identifier string) (io.Reader, error)
}

type idResolver struct{}

func NewIDResolver() Resolver {
	return &idResolver{}
}

func (r idResolver) Resolve(identifier string) (io.Reader, error) {
	panic("not implemented")
}
