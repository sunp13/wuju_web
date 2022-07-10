package models

import (
	"fmt"

	"github.com/sunp13/dbtool"
)

type comingModel struct{}

// 获取最近开赛的500场比赛
func (m *comingModel) GetList(begin, timeMax string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where comm_time < ?
	  and comm_time > ?
	order by comm_time
	limit 500
	`

	params := []interface{}{
		timeMax,
		begin,
	}

	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// GetListByID
func (m *comingModel) GetListByID(id string) (res map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where comm_id = ?
	`
	params := []interface{}{
		id,
	}
	var result []map[string]interface{}
	result, err = dbtool.D.QuerySQL(sql, params)
	if len(result) > 0 {
		res = result[0]
	}
	return
}

// GetListByLeagueID 获取锦标赛内
func (m *comingModel) GetListByLeagueID(leagueID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where league_id = ?
	order by comm_time desc
	limit 500
	`
	params := []interface{}{
		leagueID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// GetListByHomeID
func (m *comingModel) GetListByHomeID(homeID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where home_id = ?
	order by comm_time desc
	limit 500
	`
	params := []interface{}{
		homeID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// GetListByHomeID
func (m *comingModel) GetListByAwayID(awayID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where away_id = ?
	order by comm_time desc
	limit 500
	`
	params := []interface{}{
		awayID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// GetListByHomeID
func (m *comingModel) GetListByTeamID(homeID, awayID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where home_id = ?
	or away_id = ?
	order by comm_time desc
	limit 500
	`
	params := []interface{}{
		awayID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// 获取主队信息
func (m *comingModel) GetHomeList() (res []map[string]interface{}, err error) {
	sql := `
	SELECT distinct home_id,home_name FROM b365api.up_coming
	`
	res, err = dbtool.D.QuerySQL(sql, nil)
	return
}

// 获取客队信息
func (m *comingModel) GetAwayList() (res []map[string]interface{}, err error) {
	sql := `
	SELECT distinct away_id,away_name FROM b365api.up_coming
	`
	res, err = dbtool.D.QuerySQL(sql, nil)
	return
}

//
func (m *comingModel) GetFilterList(leagueID, beginTime, endTime, homeID, awayID string) (res []map[string]interface{}, err error) {

	sql := `
	select * from b365api.up_coming
	where 1=1
	`
	if leagueID != "" {
		sql += fmt.Sprintf("and league_id = '%s' ", leagueID)
	}

	if homeID != "" {
		sql += fmt.Sprintf("and home_id = '%s' ", homeID)
	}

	if awayID != "" {
		sql += fmt.Sprintf("and away_id = '%s' ", awayID)
	}

	if beginTime != "" {
		sql += fmt.Sprintf("and comm_time > '%s' ", beginTime)
	}

	if endTime != "" {
		sql += fmt.Sprintf("and comm_time < '%s' ", endTime)
	}

	sql += "order by comm_time"

	res, err = dbtool.D.QuerySQL(sql, nil)
	return
}
