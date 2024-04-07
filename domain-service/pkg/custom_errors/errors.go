package customerrors

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrUnknownType
	ErrUnknownType = errors.New("unkown method type, add it to interface")

	ErrNotFound = errors.New("not found")

	// ErrUserNotFound - ошибка появляется, когда не был найден пользователь
	ErrUserNotFound = errors.New("no users have found")
	// ErrUserAlreadyExists - ошибка, говорящая, что этот пользователь уже существует
	ErrUserAlreadyExists = errors.New("user already exists")
)

type ErrCodes struct {
	Err  error `json:"err"`
	Code int   `json:"code"`
}

func (e ErrCodes) Error() string {
	return e.Err.Error()
}
func (e ErrCodes) StatusCode() int {
	return e.Code
}
func (e ErrCodes) MarshalJSON() ([]byte, error) {
	var result struct {
		Err  string `json:"err"`
		Code int    `json:"code"`
	}

	result.Code = e.Code
	result.Err = e.Error()

	return json.Marshal(result)
}

func CodesNotFound(err error) error {
	return ErrCodes{
		Err:  err,
		Code: http.StatusNotFound,
	}
}

func CodesBadRequest(err error) error {
	return ErrCodes{
		Err:  err,
		Code: http.StatusBadRequest,
	}
}

func CodesNotImplemented(err error) error {
	return ErrCodes{
		Err:  err,
		Code: http.StatusNotImplemented,
	}
}
