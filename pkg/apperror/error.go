package apperror

type cErr struct {
	CError   error
	RawError error
}

func (e *cErr) Error() (string, string) {
	return e.CError.Error(), e.RawError.Error()
}

func NewCErr(cError error, rawError error) *cErr {
	return &cErr{
		CError:   cError,
		RawError: rawError,
	}
}

// TODO: Implemented the error handler for DB and auth
