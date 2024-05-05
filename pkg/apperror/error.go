package apperror

type CErr struct {
	CError   error
	RawError error
}

func (e *CErr) Error() (string, string) {
	return e.CError.Error(), e.RawError.Error()
}

func NewCErr(cError error, rawError error) *CErr {
	return &CErr{
		CError:   cError,
		RawError: rawError,
	}
}

// TODO: Implemented the error handler for DB and auth
