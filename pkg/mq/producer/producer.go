package producer

import (
	"context"

	"github.com/9d77v/band/pkg/mq"
)

type Producer interface {
	SendSync(ctx context.Context, msg ...*mq.Message) (*mq.SendResult, error)
	SendAsync(ctx context.Context, callback func(ctx context.Context, result *mq.SendResult, err error), msgs ...*mq.Message) error
}
