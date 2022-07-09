package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["wuju_web/controllers:AsiaController"] = append(beego.GlobalControllerRouter["wuju_web/controllers:AsiaController"],
        beego.ControllerComments{
            Method: "JSONList",
            Router: "/asia/list.json",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wuju_web/controllers:EventController"] = append(beego.GlobalControllerRouter["wuju_web/controllers:EventController"],
        beego.ControllerComments{
            Method: "JSONList",
            Router: "/event/list.json",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wuju_web/controllers:FulltimeController"] = append(beego.GlobalControllerRouter["wuju_web/controllers:FulltimeController"],
        beego.ControllerComments{
            Method: "JSONList",
            Router: "/fulltime/list.json",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wuju_web/controllers:GoalController"] = append(beego.GlobalControllerRouter["wuju_web/controllers:GoalController"],
        beego.ControllerComments{
            Method: "JSONList",
            Router: "/goal/list.json",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wuju_web/controllers:LeagueController"] = append(beego.GlobalControllerRouter["wuju_web/controllers:LeagueController"],
        beego.ControllerComments{
            Method: "JSONGetList",
            Router: "/league/list.json",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wuju_web/controllers:UpComingController"] = append(beego.GlobalControllerRouter["wuju_web/controllers:UpComingController"],
        beego.ControllerComments{
            Method: "JSONList",
            Router: "/upcoming/list.json",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
