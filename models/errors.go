package models

type Errors struct {
	typ int
	ok  bool
	err error
	msg string
}

// Error Type for the customized Errors
const (
	ERROR = iota
	OK
)

func NewError(err error, msg string) Errors {
	return Errors{typ: ERROR, msg: msg, err: err}
}

func NewOk(ok bool, msg string) Errors {
	return Errors{typ: OK, msg: msg, ok: ok}
}

func (e *Errors) GetOk() bool {
	return e.ok
}

func (e *Errors) GetError() error {
	return e.err
}

func (e *Errors) GetType() int {
	return e.typ
}

func (e *Errors) GetMessage() string {
	return e.msg
}
