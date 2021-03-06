package reader

import (
	"context"

	"github.com/criticalstack/swoll/pkg/kernel"
	"github.com/criticalstack/swoll/pkg/topology"
	"github.com/go-redis/redis"
)

type EventReader interface {
	Read() chan interface{}
	Run(ctx context.Context) error
}

type EmptyReader struct{}

func (e *EmptyReader) Read() chan interface{}        { return nil }
func (e *EmptyReader) Run(ctx context.Context) error { return nil }

// NewEventReader attempts to create an EventReader of a known type
// Currently only knows: *kernel.Probe, *redis.PubSub
func NewEventReader(src interface{}) EventReader {
	switch src := src.(type) {
	case *kernel.Probe:
		return NewKernelReader(src)
	case *redis.PubSub:
		return NewRedisReader(src)
	case *topology.Topology:
		return NewTopologyReader(src)
	}

	return &EmptyReader{}
}
