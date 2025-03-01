package entity

import (
	"net/http"
	"time"

	commonError "github.com/madevara24/go-common/errors"
)

type Question struct {
	Content   string    `json:"content"`
	Embedding []float64 `json:"embedding"`
}

type Embedding struct {
	UUID      string    `json:"uuid"`
	Content   string    `json:"content"`
	Embedding []float64 `json:"embedding"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	EMBEDDING_TABLE_NAME = "embeddings"

	ERR_CODE_QUESTION_EMPTY = "QUESTION_001"
)

var (
	ERR_QUESTION_EMPTY = commonError.NewErr(http.StatusBadRequest, ERR_CODE_QUESTION_EMPTY, "question is empty")
)

func (q *Question) Validate() error {
	if q.Content == "" {
		return ERR_QUESTION_EMPTY
	}
	return nil
}
