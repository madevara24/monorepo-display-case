package openai

import (
	"context"
	"fmt"
	"strings"

	gopenai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client *gopenai.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		client: gopenai.NewClient(apiKey),
	}
}

func (c *Client) CreateEmbedding(ctx context.Context, text string) ([]float32, error) {
	resp, err := c.client.CreateEmbeddings(ctx, gopenai.EmbeddingRequest{
		Model: gopenai.AdaEmbeddingV2,
		Input: []string{text},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embedding data received")
	}

	return resp.Data[0].Embedding, nil
}

func (c *Client) GenerateResponse(ctx context.Context, question string, similarContents []string) (string, error) {
	// Combine similar contents into context
	context := strings.Join(similarContents, "\n\n")

	// Create the system and user messages
	messages := []gopenai.ChatCompletionMessage{
		{
			Role: gopenai.ChatMessageRoleSystem,
			Content: "You are a helpful assistant that answers questions based on the provided context. " +
				"Only use the information from the context to answer questions. " +
				"If you cannot answer the question based on the context, say so.",
		},
		{
			Role:    gopenai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Context:\n%s\n\nQuestion: %s", context, question),
		},
	}

	// Call OpenAI API
	resp, err := c.client.CreateChatCompletion(ctx, gopenai.ChatCompletionRequest{
		Model:    gopenai.GPT4, // or GPT3Dot5Turbo based on your needs
		Messages: messages,
	})

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response generated")
	}

	return resp.Choices[0].Message.Content, nil
}
