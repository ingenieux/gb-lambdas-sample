// +build !lambda
package main

import (
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"testing"
)

func TestHandle(t *testing.T) {
	testContext := &runtime.Context{
		ClientContext: &runtime.ClientContext{
			Client:      &runtime.Client{},
			Environment: map[string]string{},
			Custom:      map[string]string{},
		},
		Identity: &runtime.CognitoIdentity{},
	}

	Handle([]byte(`null`), testContext)
}
