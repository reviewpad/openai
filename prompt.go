// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file

package openai

import (
	"context"
	"errors"
	"fmt"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// Prompt prompts the OpenAI API with a list of messages using 0.000001 temperature and a retry mechanism.
func (c *OpenAIClient) Prompt(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	var err error
	var resp openai.ChatCompletionResponse

	for i := 0; i < 3; i++ {
		resp, err = c.Client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model:       c.Model,
				Messages:    messages,
				Temperature: 0.000001,
			},
		)

		if err == nil {
			break
		}

		fmt.Printf("[debug] repeating request\n")
		e := &openai.APIError{}
		if errors.As(err, &e) {
			switch e.HTTPStatusCode {
			case 429:
				delta := 10 * (i + 1)
				time.Sleep(time.Duration(delta) * time.Second)
				continue
			case 500:
				continue
			default:
				break
			}
		}
	}

	if err != nil {
		return "", err
	}

	reply := resp.Choices[0].Message.Content
	return reply, nil
}
