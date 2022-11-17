# AI-Powered Copilot CLI

Your life Copilot powered by OpenAI (CLI interface for OpenAI with searching).

Features:

* Web access and Google search support without leaving the terminal.
* Automatically execute any steps predicted from prompt instructions.
* Human interactions on uncertain instructions to avoid inappropriate operations.

## Install

Install the copilot with the commands below:

```sh
go install github.com/rusthigh/ai-powered-copilot-cli/cmd/ai-powered-copilot-cli@latest
```

## Setup

* OpenAI API key should be set to `OPENAI_API_KEY` environment variable to enable the ChatGPT feature.
  * `OPENAI_API_BASE` should be set as well for Azure OpenAI service and other self-hosted OpenAI services.
* Google Search API key and CSE ID should be set to `GOOGLE_API_KEY` and `GOOGLE_CSE_ID`.

## How to use

```sh
Usage:
  ai-powered-copilot-cli [flags]

Flags:
  -c, --count-tokens     Print tokens count
  -h, --help             help for ai-powered-copilot-cli
  -t, --max-tokens int   Max tokens for the GPT model (default 1024)
  -m, --model string     OpenAI model to use (default "gpt-4")
  -p, --prompt string    Prompts sent to GPT model for non-interactive mode. If not set, interactive mode is used
  -v, --verbose          Enable verbose output (default true)
```

### Interactive mode

Here is a conversation sample (user inputs are after `You:`):

```sh
$ ai-powered-copilot-cli --verbose=false
You: What is OpenAI?
AI: OpenAI is an artificial intelligence r