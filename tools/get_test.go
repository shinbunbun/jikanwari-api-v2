package tools

import (
	"encoding/json"
	"testing"
)

func TestGetRequest(t *testing.T) {
	testEndpointBase := "https://httpbin.org/"
	type args struct {
		url    string
		header map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "200",
			args: args{
				url: testEndpointBase + "get",
				header: map[string]string{
					"Authorization": "Bearer testToken",
				},
			},
			wantErr: false,
		},
		{
			name: "400",
			args: args{
				url:    testEndpointBase + "status/400",
				header: map[string]string{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := GetRequest(tt.args.url, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			var resMap map[string]interface{}
			_ = json.Unmarshal([]byte(resp), &resMap)
			headers := resMap["headers"].(map[string]interface{})
			for k, v := range tt.args.header {
				if headers[k] != v {
					t.Errorf("GetRequest() header = %v, %v, want %v", k, headers[k], tt.args.header)
				}
			}
		})
	}
}
