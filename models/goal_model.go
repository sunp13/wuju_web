package models

import "github.com/sunp13/dbtool"

type goalModel struct{}

// GetListByCommID
func (m *goalModel) GetListByCommID(commID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.goal_line
	where comm_id = ?
	order by add_time desc
	`
	params := []interface{}{
		commID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}
