package models

import "github.com/sunp13/dbtool"

type eventModel struct{}

// GetListByCommID
func (m *eventModel) GetListByCommID(commID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.inplay_event
	where comm_id = ?
	order by add_time desc
	`
	params := []interface{}{
		commID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}
