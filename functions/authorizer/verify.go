package main

import (
	"fmt"
	"strings"

	"github.com/shinbunbun/jikanwari-api-v2/tools"
)

func getTokenFromAuthZHeader(authZHeader string) string {
	return strings.Split(authZHeader, "Bearer ")[1]
}

func verify(authZHeader string) error {

	idToken := getTokenFromAuthZHeader(authZHeader)

	_, err := tools.GetRequest("https://api.line.me/oauth2/v2.1/verify?access_token="+idToken, nil)
	if err != nil {
		return fmt.Errorf("Invalid token: %v", err.Error())
	}

	return nil
}
