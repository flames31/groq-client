package groq

import "github.com/go-resty/resty/v2"

type Client struct {
	apiKey string
	resty  resty.Client
	model  string
}

func NewClient(apiKey, model string) *Client {
	client := resty.New().SetBaseURL("https://api.groq.com/openai/v1").
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetRetryCount(3)

	return &Client{
		apiKey: apiKey,
		resty:  *client,
		model:  model,
	}
}
