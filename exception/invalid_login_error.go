package exception

import "nozzlium/kepo_backend/constants"

type InvalidLoginError struct {
}

func (err InvalidLoginError) Error() string {
	return constants.INVALID_CREDENTIAL
}
