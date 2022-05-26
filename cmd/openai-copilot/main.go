package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/feiskyer/openai-copilot/pkg/assistants"
	"github.com/feiskyer/openai-copilot/pkg/consts"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	// global flags
	model, prompt string
	maxTokens     int
	maxIterations int
	countTokens   bool
	verbose       bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "openai-copilot",
		Short: "OpenAI Copilot",
		Run: func(cmd *cobra.Command, args []string) {
			chat()
		},
	}
)

func chat() {
	var err error
	var response string
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: consts.DefaultPrompt,
		},
	}

	// Non-interactive mode
	if prompt != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		})
		response, _, err = assistants.Assistant(model, messages, maxTokens, countTokens, verbose, maxIterations)
		if err != nil {
			color.Red(err.Error())
			return
		}

		fmt.Printf("%s\n\n", response)
		return
	}

	// Interactive mode
	color.New(color.FgYellow).Printf("You: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		})
		response, messages, err = assistants.Assistant(model, messages, maxTokens, countTokens, verbose, maxIterations)
		if err != nil {
			color.Red(err.Error())
			continue
		}

		color.New(color.FgYellow).Printf("AI: ")
		fmt.Printf("%s\n\n", response)
		color.New