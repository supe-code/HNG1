package helper2

import (
	"context"
	"os"
	"strconv"
	"strings"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type Data struct{
	Operation string `json:"operation_type"`
	X int `json:"x"`
	Y int `json:"y"`
}

var Req Data

type Result struct{
	SlackUsername string `json:"slackUsername"`
	Result int `json:"result"`
	OperationType string `json:"operation_type"`
	Key string `json:"key,omitempty"`
}

var Res Result = Result{SlackUsername: "Arigbede Jacob"}

func Addition(x,y int)int{
	return x+y
}
func Subtraction(x,y int)int{
	return x-y
}
func Multiplication(x,y int)int{
	return x*y
}
func GPTVal(prompt string) (int,error){
	c := gogpt.NewClient(os.Getenv("open-key"))
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model: gogpt.GPT3TextDavinci001,
		MaxTokens: 5,
		Prompt:   prompt + " \"result only\"",
		Temperature: 0,
		TopP: 1,
		BestOf: 1,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return 0,err
	}
	p,_:= strconv.Atoi(strings.Trim(resp.Choices[0].Text,"\n"))
	return p,nil
}
func GPTOpr(prompt string) (string,error){
	c := gogpt.NewClient(os.Getenv("open-key"))
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model: gogpt.GPT3TextDavinci001,
		MaxTokens: 5,
		Prompt:   "in one word what operation should i use to solve \""+prompt+"\"",
		Temperature: 0,
		TopP: 1,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "",err
	}
	return strings.Trim(resp.Choices[0].Text,"\n"),err
}