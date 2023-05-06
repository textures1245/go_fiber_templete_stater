package util

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logFile = "/tmp/simple-log-go.log"
var logLevel = "DEBUG"

func Init() {
	fmt.Println("-= Init Data =-")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	logLevel = viper.GetString("LOG_LEVEL")
	println("log:" + logLevel)

	file, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", logFile, "%Y-%m-%d"),
		rotatelogs.WithLinkName(logFile+".link"),
		rotatelogs.WithMaxAge(time.Second*60*24*5),
		rotatelogs.WithRotationTime(time.Second*60*24),
	)

	mw := io.MultiWriter(os.Stdout, file)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	log.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        true,
		NoColors:        false,
		FieldsOrder:     []string{"component", "function"},
	})

	log.SetOutput(mw)

	switch logLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}
}
