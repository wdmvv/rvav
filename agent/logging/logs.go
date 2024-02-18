package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoggerSetup() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "01/02 - 03:04:05"})
}

func ReportErr(msg string, err error) {
	log.Err(err).Msg(msg)
}

func ReportAction(msg string) {
	log.Info().Msg(msg)
}
