package errors

import (
	"net/http"

	commonError "github.com/madevara24/go-common/errors"
)

const (
	// Question
	ERR_CODE_UNAUTHORIZED      = "UNAUTHORIZED_001"
	ERR_CODE_RATE_LIMIT_EXCEED = "RATE_LIMIT_EXCEED_001"
)

var (
	// Question
	ERR_UNAUTHORIZED      = commonError.NewErr(http.StatusUnauthorized, ERR_CODE_UNAUTHORIZED, "invalid API key")
	ERR_RATE_LIMIT_EXCEED = commonError.NewErr(http.StatusTooManyRequests, ERR_CODE_RATE_LIMIT_EXCEED, "rate limit exceeded")
)
