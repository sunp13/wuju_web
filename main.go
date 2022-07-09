package main

import (
	"wuju_web/initlize"
	_ "wuju_web/routers"

	"github.com/beego/beego/v2/server/web"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := initlize.InitALL(); err != nil {
		log.Fatal().Str("err", err.Error()).Send()
	}
	web.Run()
}
