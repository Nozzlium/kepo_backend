package exception

type BadRequestError struct {
}

func (err BadRequestError) Error() string {
	return "bad request"
}
