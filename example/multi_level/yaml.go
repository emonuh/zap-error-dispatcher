package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/emonuh/zap-error-dispatcher/config"
	"go.uber.org/zap"
)

func main() {
	configYaml, err := ioutil.ReadFile("example/multi_level/config.yaml")
	if err != nil {
		panic(err)
	}
	var myConfig config.MultiLevel
	if err := yaml.Unmarshal(configYaml, &myConfig); err != nil {
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
