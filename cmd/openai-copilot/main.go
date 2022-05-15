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
	countTokens   b