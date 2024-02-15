package logs

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoggerSetup() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
}

func ReportErr(msg string, err error){
	log.Err(err).Msg(msg)
}

func ReportAction(msg string){
	log.Info().Msg(msg)
}