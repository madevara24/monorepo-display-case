package openai

import (
	"fmt"
	"net/http"

	commonError "github.com/madevara24/go-common/errors"
)

const (
	ERR_CODE_OPENAI_DEFAULT         = "OPENAI_001"
	ERR_CODE_OPENAI_EMBEDDING_EMPTY = "OPENAI_002"
)

var (
	ERR_OPENAI_DEFAULT = func(err error) *commonError.Err {
		return commonError.NewErr(http.StatusInternalServerError, ERR_CODE_OPENAI_DEFAULT, fmt.Errorf("openai error: %w", err).Error())
	}

	ERR_OPENAI_EMBEDDING_EMPTY = commonError.NewErr(http.StatusInternalServerError, ERR_CODE_OPENAI_EMBEDDING_EMPTY, "no embedding data received")
)
