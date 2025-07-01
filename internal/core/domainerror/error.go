package domainerror

import "errors"

var ErrTooManyRequests = errors.New("too many requests, try again later")
