package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shinbunbun/jikanwari-api-v2/tools"
)

func handler(_ context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authZHeader := request.AuthorizationToken
	if authZHeader == "" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Authorization header is empty")
	}

	prodEnv := tools.Prod{}
	var config tools.Config = &prodEnv

	err := verify(config, authZHeader)
	if err != nil {
		return generatePolicy("AllowdUser", "Deny", request.MethodArn), nil
	}

	return generatePolicy("DeniedUser", "Allow", request.MethodArn), nil
}

func main() {
	lambda.Start(handler)
}
