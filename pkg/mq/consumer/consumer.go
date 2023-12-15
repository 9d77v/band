package consumer

import (
	"context"

	"github.com/9d77v/band/pkg/mq"
)

type Consumer interface {
	Shutdown() error
	Subscribe(topic string, callback func(ctx context.Context, msgs ...*mq.MessageExt) (int, error)) error
}
