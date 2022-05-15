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
		Short: "Open