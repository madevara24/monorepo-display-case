package openai

import (
	"context"
	"fmt"

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
