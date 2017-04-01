package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emonuh/zap-error-dispatcher/config"
	"go.uber.org/zap"
)

func main() {
	configJson, err := ioutil.ReadFile("example/error_dispatcher/config.json")
	if err != nil {
		panic(err)
	}
	var myConfig config.ErrorDispatcherConfig
	if err := json.Unmarshal(configJson, &myConfig); err != nil {
		panic(err)
	}
	logger, err := myConfig.Build()
	if err != nil {
		panic(err)
	}

	logger.Debug("debug log")
	logger.Error("error log")
	logger = logger.With(zap.String("with", "with_value"))
	logger.Info("debug log")
	logger.DPanic("dpanic log")
	logger = logger.Named("named")
	logger.Warn("warn log")
	logger.Fatal("fatal log")
}
