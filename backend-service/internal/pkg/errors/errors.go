package errors

import (
	"fmt"
	"net/http"

	commonError "github.com/madevara24/go-common/errors"
)

const (
	// QUESTION
	ERR_CODE_UNAUTHORIZED      = "UNAUTHORIZED_001"
	ERR_CODE_RATE_LIMIT_EXCEED = "RATE_LIMIT_EXCEED_001"

	// DATABASE
	ERR_CODE_DB_DEFAULT = "DB_001"
)

var (
	// QUESTION
	ERR_UNAUTHORIZED      = commonError.NewErr(http.StatusUnauthorized, ERR_CODE_UNAUTHORIZED, "invalid API key")
	ERR_RATE_LIMIT_EXCEED = commonError.NewErr(http.StatusTooManyRequests, ERR_CODE_RATE_LIMIT_EXCEED, "rate limit exceeded")

	// DATABASE
	ERR_DB_DEFAULT = func(err error) *commonError.Err {
		return commonError.NewErr(http.StatusInternalServerError, ERR_CODE_DB_DEFAULT, fmt.Errorf("database error: %w", err).Error())
	}
)
