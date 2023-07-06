package exception

type UnauthorizedError struct {
}

func (err UnauthorizedError) Error() string {
	return "unauthorized"
}
