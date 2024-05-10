package util

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	logrus_logstash "github.com/sima-land/logrus-logstash-hook"
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

	// Create a new viper instance to read the .env file
	viper2 := viper.New()
	viper2.SetConfigName(".env")
	viper2.SetConfigType("env")
	viper2.AddConfigPath(".")

	errOnReadCfg := viper2.ReadInConfig()
	if errOnReadCfg != nil {
		log.Fatalf("Error while reading config file %s", errOnReadCfg)
	}

	// Merge the .env file with the existing configuration
	for _, key := range viper2.AllKeys() {
		viper.Set(key, viper2.Get(key))
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

	hook, err := logrus_logstash.NewHook("tcp", viper.GetString("LOGSTASH"), "simple-noti")
	if err != nil {
		log.Error(err)
	} else {
		log.Info("-= Add Log Stash =-")
		log.AddHook(hook)
	}
}
