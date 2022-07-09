package initlize

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLog() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	writer := &lumberjack.Logger{
		Filename:   "./logs/web.log",
		MaxSize:    1000,
		MaxBackups: 5,
		LocalTime:  true,
	}
	log.Logger = zerolog.New(writer).With().Timestamp().Logger()
	log.Info().Str("module", "init").Msg("init_log succ")
}
