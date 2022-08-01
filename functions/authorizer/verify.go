package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shinbunbun/jikanwari-api-v2/tools/config"
	"github.com/shinbunbun/jikanwari-api-v2/tools/request"
)

type VerifyApiResponse struct {
	ClientId  string `json:"client_id"`
	Scope     string `json:"scope"`
	ExpiresIn int    `json:"expires_in"`
}

func getTokenFromAuthZHeader(authZHeader string) string {
	return strings.Split(authZHeader, "Bearer ")[1]
}

func parseVerifyApiResponse(resp string) (*VerifyApiResponse, error) {
	var parsedResp = new(VerifyApiResponse)
	err := json.Unmarshal([]byte(resp), &parsedResp)
	if err != nil {
		return &VerifyApiResponse{}, fmt.Errorf("JSON unmarshal error: %v", err.Error())
	}
	return parsedResp, nil
}

func verifyClientId(config config.Config, parsedResp *VerifyApiResponse) error {
	if parsedResp.ClientId != config.GetEnv("LINE_CHANNEL_ID") {
		return fmt.Errorf("Invalid client id: %v", parsedResp.ClientId)
	}
	return nil
}

func verify(config config.Config, authZHeader string) error {

	idToken := getTokenFromAuthZHeader(authZHeader)

	resp, err := request.GetRequest("https://api.line.me/oauth2/v2.1/verify?access_token="+idToken, nil)
	if err != nil {
		return fmt.Errorf("Invalid token: %v", err.Error())
	}

	parsedResp, err := parseVerifyApiResponse(resp)
	if err != nil {
		return err
	}

	return verifyClientId(config, parsedResp)
}
