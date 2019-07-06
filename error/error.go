package error

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string  {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

