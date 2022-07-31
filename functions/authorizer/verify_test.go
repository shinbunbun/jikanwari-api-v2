package main

import "testing"

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
