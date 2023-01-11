package common

import "fmt"

type ErrCode int

const (
	ErrCodeHeaderEmpty ErrCode = 1
)

var code2msg = map[ErrCode]string{
	ErrCodeHeaderEmpty: "header_is_empty",
}

var (
	ErrHeaderEmpty = BlockChainErr{Code: ErrCodeHeaderEmpty, Message: "header_is_empty" }
)

var _ error = BlockChainErr{}

type BlockChainErr struct {
	Code    ErrCode
	Message string
	Cause   error
}

func (e BlockChainErr) Error() string {
	return fmt.Sprintf("[code]=%+v,[message]=%+v,[cause]=%+v", e.Code, e.Message, e.Cause)
}

func (e BlockChainErr) Upwrap() error {
	return e.Cause
}

func (e BlockChainErr) WithCause(err error) BlockChainErr {
	e.Cause = err
	return e
}


