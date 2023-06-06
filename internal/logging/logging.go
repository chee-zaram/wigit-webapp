package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

const (
	// logsDir is the directory where the log files are contained
	logsDir = "logs"
	// logFileName is the name of the production log file
	logFileName = "wigit_webapp_prod.log"
)

// logFilePath if the full path to the production log file
var logFilePath = filepath.Join(logsDir, logFileName)

// ConfigureLogger sets up a global logger using `zerolog` package for the program.
//
// If in `dev` mode, it logs to `stderr`, if in `prod` it logs to the wigit log file.
func ConfigureLogger(env string) *os.File {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	switch env {
	case "dev":
		stdOutWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger := zerolog.New(stdOutWriter).With().Timestamp().Logger()
		log.Logger = logger
	case "prod":
		createLogDir()
		backupLastLog()
		logFile := openLogFile()
		logFileWriter := zerolog.ConsoleWriter{Out: logFile, NoColor: true, TimeFormat: time.RFC3339}
		logger := zerolog.New(logFileWriter).With().Timestamp().Logger()
		log.Logger = logger
		return logFile
	default:
		fmt.Printf("env not valid: %s\n", env)
		os.Exit(2)
	}

	return nil
}

// createLogDir creates a directory for the program logs if not exists.
func createLogDir() {
	if err := os.Mkdir(logsDir, 0744); err != nil && !os.IsExist(err) {
		log.Fatal().Err(err).Msg("unable to create logs directory")
	}
}

// backupLastLog backs up the previous log file if exists using current date time.
func backupLastLog() {
	timeStamp := time.Now().UTC().Format("20060201_15_04_05")
	base := strings.TrimSuffix(logFileName, filepath.Ext(logFileName))
	bkpLogName := fmt.Sprintf("%s_%s%s", base, timeStamp, filepath.Ext(logFilePath))
	bkpLogPath := filepath.Join(logsDir, bkpLogName)

	logFile, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Panic().Err(err).Msg("failed to read file for backup")
	}

	if err = ioutil.WriteFile(bkpLogPath, logFile, 0644); err != nil {
		log.Panic().Err(err).Msg("failed to write to backup log file")
	}
}

// openLogFile opens the configured log file for writing.
func openLogFile() *os.File {
	logFile, err := os.OpenFile(logFilePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Panic().Err(err).Msg("failed to open log file for writing")
	}

	return logFile
}

// SetGinLogToFile sets gin's logger to be the custom log file.
func SetGinLogToFile(logFile *os.File) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(logFile)
}

// SetGORMLogToFile sets gorm to log to log file
func SetGORMLogToFile() logger.Interface {
	return logger.New(&log.Logger, logger.Config{
		SlowThreshold:             0,
		Colorful:                  false,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Error,
	})
}
