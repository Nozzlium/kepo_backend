package exception

import "nozzlium/kepo_backend/constants"

type NotFoundError struct {
}

func (err NotFoundError) Error() string {
	return constants.NOT_FOUND
}
