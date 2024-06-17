package gptErrors

type GtpError struct {
	error

	fatal bool
}

func New(err error, fatal bool) GtpError {
	return GtpError{
		error: err,
		fatal: fatal,
	}
}

func (err *GtpError) IsFatal() bool {
	return err.fatal
}
