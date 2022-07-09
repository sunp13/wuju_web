package models

import "github.com/sunp13/dbtool"

type comingModel struct{}

// 获取最近开赛的500场比赛
func (m *comingModel) GetList(timeMax string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	order by comm_time desc
	limit 1000
	`
	res, err = dbtool.D.QuerySQL(sql, nil)
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
