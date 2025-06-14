package groq

import (
	"context"
	"errors"
)

func (c *Client) Chat(ctx context.Context, messages []Message) (string, error) {
	req := ChatCompletionRequest{
		Model:    c.model,
		Messages: messages,
	}

	var resp ChatCompletionResponse

	_, err := c.resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/chat/completions")

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no response received")
	}

	return resp.Choices[0].Message.Content, nil
}
