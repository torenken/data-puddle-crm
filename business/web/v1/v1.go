package v1

type RequestError struct {
	Err    error
	Status int
}

func NewRequestError(err error, status int) error {
	return &RequestError{err, status}
}

func (re *RequestError) Error() string {
	return re.Err.Error()
}
