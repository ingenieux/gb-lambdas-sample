// build +lambda
package main

import (
	"encoding/json"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{DisableColors: true, DisableTimestamp: true})
}

func Handle(req json.RawMessage, ctx *runtime.Context) (string, error) {
	log.Info("Context: %+v", *ctx)
	log.Info("Request received: %s", string(req))

	return string(req), nil
}
