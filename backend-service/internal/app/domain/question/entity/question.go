package entity

import (
	"net/http"
	"strings"
	"time"

	commonError "github.com/madevara24/go-common/errors"
)

type Question struct {
	UUID      string    `json:"uuid"`
	Content   string    `json:"content"`
	Embedding []float32 `json:"embedding"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QuestionAnswer struct {
	UUID          string    `json:"uuid"`
	QuestionUUID  string    `json:"question_uuid"`
	KnowledgeUUID string    `json:"knowledge_uuid"`
	Score         float64   `json:"score"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

const (
	EMBEDDING_TABLE_NAME = "embeddings"

	// Question
	ERR_CODE_QUESTION_EMPTY    = "QUESTION_001"
	ERR_CODE_QUESTION_TOO_LONG = "QUESTION_002"

	MAX_QUESTION_LENGTH = 1000 // Maximum length for question content

	// Embedding
	ERR_CODE_EMBEDDING_EMPTY = "EMBEDDING_001"
)

var (
	// Question
	ERR_QUESTION_EMPTY    = commonError.NewErr(http.StatusBadRequest, ERR_CODE_QUESTION_EMPTY, "question is empty")
	ERR_QUESTION_TOO_LONG = commonError.NewErr(http.StatusBadRequest, ERR_CODE_QUESTION_TOO_LONG, "question is too long")

	// Embedding
	ERR_EMBEDDING_EMPTY = commonError.NewErr(http.StatusBadRequest, ERR_CODE_EMBEDDING_EMPTY, "embedding is empty")
)

func (q *Question) Validate() error {
	if q.Content == "" {
		return ERR_QUESTION_EMPTY
	}

	// Sanitize the content
	q.Content = sanitizeContent(q.Content)

	if len(q.Content) > MAX_QUESTION_LENGTH {
		return ERR_QUESTION_TOO_LONG
	}

	return nil
}

// sanitizeContent handles special characters in the content
func sanitizeContent(s string) string {
	// Only replace control characters (except newlines, tabs, carriage returns) with spaces
	var builder strings.Builder
	for _, r := range s {
		if r < 32 && r != '\n' && r != '\t' && r != '\r' {
			builder.WriteRune(' ')
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
