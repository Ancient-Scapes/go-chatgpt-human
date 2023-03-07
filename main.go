package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Printf("%s\n", args[0])
	}

	client := openai.NewClient(os.Getenv("API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: args[0],
				},
			},
		},
	)

	if err != nil {
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
