package models

import "github.com/sunp13/dbtool"

type fulltimeModel struct{}

// GetListByCommID
func (m *fulltimeModel) GetListByCommID(commID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.full_time
	where comm_id = ?
	order by add_time desc
	`
	params := []interface{}{
		commID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}
