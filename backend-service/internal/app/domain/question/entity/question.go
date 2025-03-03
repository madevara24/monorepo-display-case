package entity

import (
	"net/http"
	"time"

	commonError "github.com/madevara24/go-common/errors"
)

type Question struct {
	Content   string    `json:"content"`
	Embedding []float32 `json:"embedding"`
}

type Embedding struct {
	UUID        string    `json:"uuid"`
	Category    string    `json:"category"`
	Granularity string    `json:"granularity"`
	Content     string    `json:"content"`
	Embedding   []float32 `json:"embedding"`
	Score       float64   `json:"score"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	EMBEDDING_TABLE_NAME = "embeddings"

	// Question
	ERR_CODE_QUESTION_EMPTY = "QUESTION_001"

	// Embedding
	ERR_CODE_EMBEDDING_EMPTY = "EMBEDDING_001"
)

var (
	// Question
	ERR_QUESTION_EMPTY = commonError.NewErr(http.StatusBadRequest, ERR_CODE_QUESTION_EMPTY, "question is empty")

	// Embedding
	ERR_EMBEDDING_EMPTY = commonError.NewErr(http.StatusBadRequest, ERR_CODE_EMBEDDING_EMPTY, "embedding is empty")
)

func (q *Question) Validate() error {
	if q.Content == "" {
		return ERR_QUESTION_EMPTY
	}
	return nil
}

func (e *Embedding) Validate() error {
	if e.Content == "" {
		return ERR_EMBEDDING_EMPTY
	}
	if e.Category == "" {
		return ERR_EMBEDDING_EMPTY
	}
	if e.Granularity == "" {
		return ERR_EMBEDDING_EMPTY
	}
	return nil
}
