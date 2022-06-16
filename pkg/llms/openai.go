
package llms

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	*openai.Client

	Retries int
	Backoff time.Duration
}

// NewOpenAIClient returns an OpenAI client.
func NewOpenAIClient() (*OpenAIClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY is not set")
	}

	config := openai.DefaultConfig(apiKey)
	baseURL := os.Getenv("OPENAI_API_BASE")
	if baseURL != "" {
		config.BaseURL = baseURL

		if strings.Contains(baseURL, "azure") {
			config.APIType = openai.APITypeAzure
			config.APIVersion = "2024-02-01"
			config.AzureModelMapperFunc = func(model string) string {
				return regexp.MustCompile(`[.:]`).ReplaceAllString(model, "")
			}
		}
	}

	return &OpenAIClient{
		Retries: 5,
		Backoff: time.Second,
		Client:  openai.NewClientWithConfig(config),
	}, nil
}
