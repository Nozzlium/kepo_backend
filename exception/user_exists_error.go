package exception

type UserExistsError struct {
}

func (err UserExistsError) Error() string {
	return "user already exists"
}
