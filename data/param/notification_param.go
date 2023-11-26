package param

import "nozzlium/kepo_backend/data/entity"

type NotificationParam struct {
	PaginationParam
	Notification entity.Notification
}
