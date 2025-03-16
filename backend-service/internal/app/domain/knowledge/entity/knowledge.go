package entity

import (
	"net/http"
	"time"

	commonError "github.com/madevara24/go-common/errors"
)

type Knowledge struct {
	UUID        string    `json:"uuid"`
	Category    string    `json:"category"`
	Granularity string    `json:"granularity"`
	Content     string    `json:"content"`
	Embedding   []float32 `json:"embedding"`
	Score       float64   `json:"score"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type KnowledgeRaw struct {
	Metadata Metadata `json:"metadata" validate:"required"`
	Content  string   `json:"content" validate:"required"`
}

type Metadata struct {
	Company     string `json:"company" validate:"required"`
	Role        string `json:"role" validate:"required"`
	Project     string `json:"project" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Year        string `json:"year" validate:"required"`
	Granularity string `json:"granularity" validate:"required"`
}

const (
	KNOWLEDGE_TABLE_NAME = "knowledge"

	ERR_CODE_KNOWLEDGE_CATEGORY_EMPTY    = "KNOWLEDGE_EMPTY_001"
	ERR_CODE_KNOWLEDGE_GRANULARITY_EMPTY = "KNOWLEDGE_EMPTY_002"
	ERR_CODE_KNOWLEDGE_CONTENT_EMPTY     = "KNOWLEDGE_EMPTY_003"
	ERR_CODE_KNOWLEDGE_EMBEDDING_EMPTY   = "KNOWLEDGE_EMPTY_004"
)

var (
	// Knowledge
	ERR_KNOWLEDGE_CATEGORY_EMPTY    = commonError.NewErr(http.StatusBadRequest, ERR_CODE_KNOWLEDGE_CATEGORY_EMPTY, "category is empty")
	ERR_KNOWLEDGE_GRANULARITY_EMPTY = commonError.NewErr(http.StatusBadRequest, ERR_CODE_KNOWLEDGE_GRANULARITY_EMPTY, "granularity is empty")
	ERR_KNOWLEDGE_CONTENT_EMPTY     = commonError.NewErr(http.StatusBadRequest, ERR_CODE_KNOWLEDGE_CONTENT_EMPTY, "content is empty")
	ERR_KNOWLEDGE_EMBEDDING_EMPTY   = commonError.NewErr(http.StatusBadRequest, ERR_CODE_KNOWLEDGE_EMBEDDING_EMPTY, "embedding is empty")
)

func (e *Knowledge) Validate() error {
	if e.Content == "" {
		return ERR_KNOWLEDGE_CONTENT_EMPTY
	}
	if e.Category == "" {
		return ERR_KNOWLEDGE_CATEGORY_EMPTY
	}
	if e.Granularity == "" {
		return ERR_KNOWLEDGE_GRANULARITY_EMPTY
	}
	return nil
}
