package repository

import (
	"backend-service/internal/app/domain/question/entity"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type QuestionRepository struct {
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) StoreEmbedding(ctx context.Context, embedding entity.Embedding) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO embeddings (uuid, category, granularity, content, embedding) VALUES ($1, $2, $3, $4, $5::float4[]::vector)",
		embedding.UUID,
		embedding.Category,
		embedding.Granularity,
		embedding.Content,
		pq.Array(embedding.Embedding),
	)
	return err
}

func (r *QuestionRepository) FindSimilar(ctx context.Context, embedding []float32, limit int) ([]entity.Embedding, error) {
	var embeddings []entity.Embedding
	err := r.db.SelectContext(ctx, &embeddings,
		`SELECT uuid, category, granularity, content, embedding::float4[]::vector
         FROM embeddings 
         ORDER BY embedding <=> $1 
         LIMIT $2`,
		pq.Array(embedding), limit,
	)
	return embeddings, err
}
