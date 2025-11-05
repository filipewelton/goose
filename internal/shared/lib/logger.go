package lib

import (
	"os"
	"regexp"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var re = regexp.MustCompile(`(retail_workflow).*`)

func Logger() *zap.Logger {
	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path = re.ReplaceAllString(path, "retail_workflow/app.log")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)

	if err != nil {
		panic(err)
	}

	writeSyncer := zapcore.AddSync(file)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)
	logger := zap.New(core)

	return logger
}
