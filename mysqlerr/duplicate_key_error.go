package mysqlerr

type DuplicateKeyError struct {
}

func (err DuplicateKeyError) Error() string {
	return "duplicate key"
}
