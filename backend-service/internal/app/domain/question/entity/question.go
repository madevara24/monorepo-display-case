package entity

import (
	"net/http"

	commonError "github.com/madevara24/go-common/errors"
)

type Question struct {
	Question string `json:"question"`
}

const (
	ERR_CODE_QUESTION_EMPTY = "QUESTION_001"
)

var (
	ERR_QUESTION_EMPTY = commonError.NewErr(http.StatusBadRequest, ERR_CODE_QUESTION_EMPTY, "question is empty")
)

func (q *Question) Validate() error {
	if q.Question == "" {
		return ERR_QUESTION_EMPTY
	}
	return nil
}
