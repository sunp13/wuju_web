package models

import "github.com/sunp13/dbtool"

type asianHandModel struct{}

func (m *asianHandModel) GetListByCommID(commID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.asian_handicap
	where comm_id = ?
	order by add_time desc
	`
	params := []interface{}{
		commID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}
