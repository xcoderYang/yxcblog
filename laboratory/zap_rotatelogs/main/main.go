package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"time"
)

/**
presets
 */
func main(){
	namespace()
}
func presets(){
	logger := zap.NewExample()
	defer logger.Sync()

	const url = "http://example.com"

	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		"url", url,
		"attempt", 3,
		"backoff", time.Second)
	sugar.Infof("Failed to fetch URL: %s", url)
	logger.Info("Failed to fetchURL.",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}
func basicCfg(){
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "../tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"foo":"bar"},
		"encoderConfig":{
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "colored"
		}
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil{
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil{
		panic(err)
	}
	defer logger.Sync()
	logger.Info("logger construction succeeded")
}
func advancedCfg(){
	highPriorty := zap.LevelEnablerFunc(func(lvl zapcore.Level)bool{
		return lvl >= zapcore.ErrorLevel
	})
	lowPriorty := zap.LevelEnablerFunc(func(lvl zapcore.Level)bool{
		return lvl < zapcore.ErrorLevel
	})

	topicDebugging := zapcore.AddSync(ioutil.Discard)
	topicErrors := zapcore.AddSync(ioutil.Discard)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriorty),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriorty),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriorty),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriorty),
		)

	logger := zap.New(core)
	defer logger.Sync()

	logger.Info("constructed a logger")
	logger.Fatal("error")
}
func atomicLevel(){
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
		))
	defer logger.Sync()
	logger.Info("info logging enabled")
	atom.SetLevel(zap.ErrorLevel)
	logger.Info("info logging disabled")
	logger.Error("error")
}
func atomicLevelCfg(){
	rawJSON := []byte(`{
		"level": "info",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoding": "json",
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil{
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil{
		panic(err)
	}
	defer logger.Sync()

	logger.Info("info logging enabled")

	cfg.Level.SetLevel(zap.ErrorLevel)
	logger.Info("info logging disabled")
}
func loggerCheck(){
	logger := zap.NewExample()
	//logger, _ := zap.NewProduction()
	defer logger.Sync()
	// Check方法会检查第一个参数 level是否 >= logger本身的 level，如果 >=，则返回checkEntry，否则返回 nil
	// 上面的例子 logger是从 NewExample生成，其默认的 level是 debug，所以所有的 Check都会返回 checkEntry
	// 如果改成 NewProduction生成，其默认的 level是 info，则 debugLevel的 Check会返回 nil，其余 level的 Check返回 checkEntry
	if ce := logger.Check(zap.DebugLevel, "debugging"); ce != nil{
		ce.Write(
			zap.String("foo", "bar"),
			zap.String("baz", "quux"),
			)
	}
}
func loggerName(){
	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("no name")
	main := logger.Named("main")
	main.Info("main logger")

	main.Named("subpackage").Info("sub-logger")
}
func namespace(){
	logger := zap.NewExample()
	defer logger.Sync()

	logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
		).Info("tracked some metrics")
}