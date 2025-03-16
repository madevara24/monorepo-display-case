package ask

import (
	"backend-service/config"
	knowledgeEntity "backend-service/internal/app/domain/knowledge/entity"
	knowledgeRepository "backend-service/internal/app/domain/knowledge/repository"
	questionEntity "backend-service/internal/app/domain/question/entity"
	questionRepository "backend-service/internal/app/domain/question/repository"
	"backend-service/internal/pkg/integration/openai"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/madevara24/go-common/txmanager"
	"golang.org/x/sync/errgroup"
)

type AskUsecase struct {
	db            *sqlx.DB
	openAIClient  *openai.Client
	questionRepo  *questionRepository.QuestionRepository
	knowledgeRepo *knowledgeRepository.KnowledgeRepository
}

func NewAskUsecase(db *sqlx.DB, openAIClient *openai.Client, questionRepo *questionRepository.QuestionRepository, knowledgeRepo *knowledgeRepository.KnowledgeRepository) *AskUsecase {
	return &AskUsecase{
		db:            db,
		openAIClient:  openAIClient,
		questionRepo:  questionRepo,
		knowledgeRepo: knowledgeRepo,
	}
}

func (i *AskUsecase) Execute(ctx context.Context, req Request) (Response, error) {

	var similar []knowledgeEntity.Knowledge
	var questionAnswers []questionEntity.QuestionAnswer
	var contents []string
	var answer string

	question, err := req.MapIntoQuestion()
	if err != nil {
		return Response{}, err
	}

	// Create vector for the question
	question.Embedding, err = i.openAIClient.CreateEmbedding(ctx, req.Question)
	if err != nil {
		return Response{}, err
	}

	questionAnswersChan := make(chan []questionEntity.QuestionAnswer, 1)

	err = txmanager.DBTransactionWrapperWithContext(ctx, i.db, func(txCtx context.Context) error {
		errgr, ctx := errgroup.WithContext(ctx)

		errgr.Go(func() error {
			return i.questionRepo.StoreQuestion(txCtx, question)
		})

		errgr.Go(func() error {
			var err error
			similar, err = i.knowledgeRepo.FindSimilar(ctx, question.Embedding, req.Limit)
			if err != nil {
				return err
			}

			// Extract contents for context
			for _, s := range similar {
				contents = append(contents, s.Content)
				questionAnswers = append(questionAnswers, questionEntity.QuestionAnswer{
					UUID:          uuid.New().String(),
					QuestionUUID:  question.UUID,
					KnowledgeUUID: s.UUID,
					Score:         s.Score,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				})
			}

			questionAnswersChan <- questionAnswers

			// Generate response using similar content
			answer, err = i.openAIClient.GenerateResponse(ctx, req.Question, contents)

			return err
		})

		errgr.Go(func() error {
			qa := <-questionAnswersChan
			if len(qa) > 0 {
				return i.questionRepo.StoreQuestionAnswers(txCtx, qa)
			}
			return nil
		})

		return errgr.Wait()
	})

	if err != nil {
		return Response{}, err
	}

	res := Response{
		Answer: answer,
	}

	if config.Get().ENV == "development" {
		res.MapIntoSet(similar)
	}

	return res, nil
}

type Request struct {
	Question string `json:"question" validate:"required"`
	Limit    int    `json:"limit"`
}

type Response struct {
	Answer string `json:"answer"`
	Set    []Set  `json:"set"`
}

type Set struct {
	Content string  `json:"content"`
	Score   float64 `json:"score"`
}

func (i *Response) MapIntoSet(similar []knowledgeEntity.Knowledge) {
	set := []Set{}
	for _, s := range similar {
		set = append(set, Set{Content: s.Content, Score: s.Score})
	}

	i.Set = set
}

func (i *Request) MapIntoQuestion() (questionEntity.Question, error) {
	var question questionEntity.Question = questionEntity.Question{
		UUID:      uuid.New().String(),
		Content:   i.Question,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := question.Validate(); err != nil {
		return questionEntity.Question{}, err
	}
	return question, nil
}
