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
			Content: "- You are an AI assistant that helps answer questions regarding Devara based on the provided context. " +
				"- Devara uses he/him pronouns. " +
				"- If multiple topics are mentioned, answer only the ones relevant to Devara's work and experience. Ignore unrelated topics." +
				"- If no useful information is found in the context, say: **I don't have information on that topic.** Do **not** attempt to infer or generate unrelated details. " +
				"- If the provided context contains **partially relevant** information, use only the parts that are useful and ignore the rest. " +
				"- Never mention \"from the context provided\" or \"from the portfolio.\" Instead, answer naturally as Devara's assistant. " +
				"- Structure responses clearly, making them easy to understand.",
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
