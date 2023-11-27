package notification

import (
	"context"
	"nozzlium/kepo_backend/data/param"
)

type NotificationService interface {
	Find(ctx context.Context, param param.NotificationParam)
}
