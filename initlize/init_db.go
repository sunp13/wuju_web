package initlize

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	"github.com/rs/zerolog/log"
	"github.com/sunp13/dbtool"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() error {
	runMode := web.AppConfig.DefaultString("runmode", "dev")
	err := dbtool.Init(fmt.Sprintf("./conf/datasource_%s.yml", runMode))
	if err != nil {
		return err
	}
	log.Info().Msg("init_db succ")
	return nil
}
