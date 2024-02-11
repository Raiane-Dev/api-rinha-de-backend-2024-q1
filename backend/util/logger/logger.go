package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger
)

func Init() {
	log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func Error(message string, err error) {
	log.Error().Err(err).Msg(message)
}

func Info(message string) {
	log.Info().Msg(message)
}
