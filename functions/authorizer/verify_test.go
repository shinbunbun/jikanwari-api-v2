package main

import (
	"reflect"
	"testing"

	"github.com/shinbunbun/jikanwari-api-v2/tools"
)

func Test_getTokenFromAuthZHeader(t *testing.T) {
	type args struct {
		authZHeader string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				authZHeader: "Bearer testToken",
			},
			want: "testToken",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTokenFromAuthZHeader(tt.args.authZHeader); got != tt.want {
				t.Errorf("getTokenFromAuthZHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseVerifyApiResponse(t *testing.T) {
	type args struct {
		resp string
	}
	tests := []struct {
		name    string
		args    args
		want    *VerifyApiResponse
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				resp: `{"client_id":"testClientId","scope":"testScope","expires_in":3600}`,
			},
			want: &VerifyApiResponse{
				ClientId:  "testClientId",
				Scope:     "testScope",
				ExpiresIn: 3600,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseVerifyApiResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseVerifyApiResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseVerifyApiResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyClientId(t *testing.T) {
	type args struct {
		parsedResp *VerifyApiResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				parsedResp: &VerifyApiResponse{
					ClientId:  "testEnvVar",
					Scope:     "testScope",
					ExpiresIn: 3600,
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				parsedResp: &VerifyApiResponse{
					ClientId:  "testDifferentId",
					Scope:     "testScope",
					ExpiresIn: 3600,
				},
			},
			wantErr: true,
		},
	}
	testEnv := tools.TestEnv{}
	var config tools.Config = &testEnv
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyClientId(config, tt.args.parsedResp); (err != nil) != tt.wantErr {
				t.Errorf("verifyClientId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
