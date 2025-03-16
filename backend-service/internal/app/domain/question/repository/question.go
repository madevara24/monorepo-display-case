package repository

import (
	"backend-service/internal/app/domain/question/entity"
	"backend-service/internal/pkg/errors"
	"context"
	"database/sql"
	"fmt"

	"github.com/madevara24/go-common/logger"
	"github.com/madevara24/go-common/txmanager"

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

func (r *QuestionRepository) StoreQuestion(ctx context.Context, question entity.Question) error {
	var (
		query string = "INSERT INTO questions (uuid, question, embedding) VALUES ($1, $2, $3::float4[]::vector)"
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

	fmt.Println(query, question.UUID, question.Content, pq.Array(question.Embedding))

	_, err = stmt.ExecContext(ctx, question.UUID, question.Content, pq.Array(question.Embedding))
	if err != nil {
		logger.Log.Error(ctx, errors.ERR_DB_DEFAULT(err).Error(), err)
		return err
	}

	return nil
}

func (r *QuestionRepository) StoreQuestionAnswers(ctx context.Context, questionAnswers []entity.QuestionAnswer) error {
	var (
		query  string = "INSERT INTO question_answers (uuid, question_uuid, knowledge_uuid, score) VALUES "
		values []interface{}
		tx     *sqlx.Tx
		stmt   *sql.Stmt
		err    error
	)

	for i, qa := range questionAnswers {
		num := i * 4
		query += fmt.Sprintf("($%d, $%d, $%d, $%d),", num+1, num+2, num+3, num+4)
		values = append(values, qa.UUID, qa.QuestionUUID, qa.KnowledgeUUID, qa.Score)
	}

	query = query[:len(query)-1] // Remove the trailing comma

	tx, _ = txmanager.ExtractTx(ctx)

	if tx != nil {
		stmt, err = tx.PrepareContext(ctx, tx.Rebind(query))
	} else {
		stmt, err = r.db.PrepareContext(ctx, r.db.Rebind(query))
	}

	fmt.Println(query, values)

	if err != nil {
		logger.Log.Error(ctx, errors.ERR_DB_DEFAULT(err).Error(), err)
		return err
	}

	_, err = stmt.ExecContext(ctx, values...)
	if err != nil {
		logger.Log.Error(ctx, errors.ERR_DB_DEFAULT(err).Error(), err)
		return err
	}

	return nil
}
