package customerror

import "nozzlium/kepo_backend/constants"

type InvalidLoginError struct {
	Err string
}

func (err InvalidLoginError) Error() string {
	return constants.INVALID_CREDENTIAL
}
