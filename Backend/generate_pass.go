package main

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Length int `json:"length"`
}

type Response struct {
	Password string `json:"password"`
}

func generatePassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	var password strings.Builder
	for i := 0; i < length; i++ {
		password.WriteByte(charset[r.Intn(len(charset))])
	}
	return password.String()
}

func handler(ctx context.Context, req Request) (Response, error) {
	if req.Length <= 0 {
		req.Length = 12
	}
	password := generatePassword(req.Length)
	return Response{Password: password}, nil
}

func main() {
	lambda.Start(handler)
}
