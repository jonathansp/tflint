package config

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/k0kubun/pp"
)

func TestNewAwsSession(t *testing.T) {
	type Result struct {
		Credentials *credentials.Credentials
		Region      *string
	}

	cases := []struct {
		Name   string
		Input  *Config
		Result Result
	}{
		{
			Name: "set credentials",
			Input: &Config{
				AwsCredentials: map[string]string{
					"access_key": "AWS_ACCESS_KEY",
					"secret_key": "AWS_SECRET_KEY",
					"region":     "us-east-1",
				},
			},
			Result: Result{
				Credentials: credentials.NewStaticCredentials("AWS_ACCESS_KEY", "AWS_SECRET_KEY", ""),
				Region:      aws.String("us-east-1"),
			},
		},
	}

	for _, tc := range cases {
		s := tc.Input.NewAwsSession()
		if !reflect.DeepEqual(s.Config.Credentials, tc.Result.Credentials) {
			t.Fatalf("\nBad: %s\nExpected: %s\n\ntestcase: %s", pp.Sprint(s.Config.Credentials), pp.Sprint(tc.Result.Credentials), tc.Name)
		}
		if !reflect.DeepEqual(s.Config.Region, tc.Result.Region) {
			t.Fatalf("\nBad: %s\nExpected: %s\n\ntestcase: %s", pp.Sprint(s.Config.Region), pp.Sprint(tc.Result.Region), tc.Name)
		}
	}
}
