package pub

import (
	"context"

	"github.com/mchmarny/kpush/pkg/msg"
)

// Publisher represents generic message publisher
type Publisher interface {
	Publish(ctx context.Context, key []byte, content *msg.SimpleMessage) error
}
