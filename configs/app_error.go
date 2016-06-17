package configs

import (
	"errors"
	"net/http"
)

type AppError struct {
	Error  error
	Status int
	Code   int
}

var (
	ErrAlreadyPay        = AppError{errors.New("order already Paid"), http.StatusBadRequest, 2801}
	ErrCanNotFindOrder   = AppError{errors.New("can not find order"), http.StatusBadRequest, 2802}
	ErrCanNotCreateOrder = AppError{errors.New("can not create order"), http.StatusBadRequest, 2803}
	ErrCanNotUpdateOrder = AppError{errors.New("can not update order"), http.StatusBadRequest, 2804}

	ErrCanNotActiveLicense  = AppError{errors.New("can not active licence"), http.StatusBadRequest, 2901}
	ErrAlreadyActiveLicense = AppError{errors.New("device license already active"), http.StatusBadRequest, 2902}
	ErrCanNotCreateLicense  = AppError{errors.New("can not create license"), http.StatusBadRequest, 2903}
	ErrCanNotFindLicense    = AppError{errors.New("could not find relate item"), http.StatusBadRequest, 2904}

	ErrCanNotReachThirdPartyService   = AppError{errors.New("can not reach thirdparty service"), http.StatusBadRequest, 3001}
	ErrInvalidReceive                 = AppError{errors.New("invalid receiveId"), http.StatusBadRequest, 3002}
	ErrCanNotDecodeThirdPartyResponse = AppError{errors.New("can not decode thirdparty response"), http.StatusBadRequest, 3003}
)

func SetAppError(appErr AppError, err error) AppError {
	appErr.Error = err
	return appErr
}
