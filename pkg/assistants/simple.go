
package assistants

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/llms"
	"github.com/feiskyer/openai-copilot/pkg/tools"
	"github.com/sashabaranov/go-openai"
)

const (
	defaultMaxIterations = 10
)

// Assistant is the simplest AI assistant.
func Assistant(model string, prompts []openai.ChatCompletionMessage, maxTokens int, countTokens bool, verbose bool, maxIterations int) (result string, chatHistory []openai.ChatCompletionMessage, err error) {
	chatHistory = prompts
	if len(chatHistory) == 0 {
		return "", nil, fmt.Errorf("prompts cannot be empty")
	}

	client, err := llms.NewOpenAIClient()
	if err != nil {
		return "", nil, fmt.Errorf("unable to get OpenAI client: %v", err)
	}

	defer func() {
		if countTokens {
			count := llms.NumTokensFromMessages(chatHistory, model)
			color.Green("Total tokens: %d\n\n", count)
		}