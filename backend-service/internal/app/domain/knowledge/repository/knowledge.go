package repository

import (
	"backend-service/internal/app/domain/knowledge/entity"
	"backend-service/internal/pkg/errors"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/madevara24/go-common/logger"
	"github.com/madevara24/go-common/txmanager"
)

type KnowledgeRepository struct {
	db *sqlx.DB
}

func NewKnowledgeRepository(db *sqlx.DB) *KnowledgeRepository {
	return &KnowledgeRepository{
		db: db,
	}
}

func (r *KnowledgeRepository) StoreKnowledge(ctx context.Context, knowledge entity.Knowledge) error {
	var (
		query string = "INSERT INTO knowledge (uuid, category, granularity, content, embedding) VALUES ($1, $2, $3, $4, $5::float4[]::vector)"
		tx    *sqlx.Tx
		stmt  *sql.Stmt
		err   error
	)

	tx, _ = txmanager.ExtractTx(ctx)

	if tx != nil {
		stmt, err = tx.PrepareContext(ctx, tx.Rebind(query))
	} else {
		stmt, err = r.db.PrepareContext(ctx, r.db.Rebind(query))
	}

	if err != nil {
		logger.Log.Error(ctx, errors.ERR_DB_DEFAULT(err).Error(), err)
		return err
	}

	_, err = stmt.ExecContext(ctx, knowledge.UUID, knowledge.Category, knowledge.Granularity, knowledge.Content, pq.Array(knowledge.Embedding))
	if err != nil {
		logger.Log.Error(ctx, errors.ERR_DB_DEFAULT(err).Error(), err)
		return err
	}

	return nil
}

func (r *KnowledgeRepository) FindSimilar(ctx context.Context, embedding []float32, limit int) ([]entity.Knowledge, error) {
	var knowledge []entity.Knowledge
	var args []interface{}

	var query = `
		WITH similar_embeddings AS (
			SELECT uuid, category, granularity, content, embedding::float4[], 
				   embedding <=> $1::float4[]::vector as score
			FROM knowledge
		)
		SELECT * FROM similar_embeddings
		ORDER BY score`

	args = append(args, pq.Array(embedding))

	if limit > 0 {
		query += " LIMIT $2"
		args = append(args, limit)
	}

	rows, err := r.db.QueryContext(ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var k entity.Knowledge
		if err := rows.Scan(
			&k.UUID,
			&k.Category,
			&k.Granularity,
			&k.Content,
			pq.Array(&k.Embedding),
			&k.Score,
		); err != nil {
			return nil, err
		}
		knowledge = append(knowledge, k)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return knowledge, nil
}
