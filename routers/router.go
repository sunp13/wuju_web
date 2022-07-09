package routers

import (
	"wuju_web/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// League
	web.Include(&controllers.LeagueController{})
	// upComing
	web.Include(&controllers.UpComingController{})
	// asia
	web.Include(&controllers.AsiaController{})
	// full
	web.Include(&controllers.FulltimeController{})
	// goal
	web.Include(&controllers.GoalController{})
	// event
	web.Include(&controllers.EventController{})
	// main
	web.Include(&controllers.MainController{})
}
