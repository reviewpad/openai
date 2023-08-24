// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file

package openai

import (
	"errors"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

const (
	OPENAI_TOKEN  = "OPENAI_TOKEN"
	OPENAI_PREFIX = "openai-"
)

type OpenAIClient struct {
	Client *openai.Client
	Model  string
}

// NewClient creates a new OpenAI API client using an environment variable.
func NewOpenAIClient(model string) (*OpenAIClient, error) {
	token, val := os.LookupEnv(OPENAI_TOKEN)
	if !val {
		return nil, errors.New("missing openai token")
	}

	client := openai.NewClient(token)

	if model == "" {
		model = openai.GPT3Dot5Turbo
	} else {
		if strings.HasPrefix(model, OPENAI_PREFIX) {
			model = strings.TrimPrefix(model, OPENAI_PREFIX)
		} else {
			return nil, errors.New("invalid model")
		}
	}

	return &OpenAIClient{
		Client: client,
		Model:  model,
	}, nil
}
