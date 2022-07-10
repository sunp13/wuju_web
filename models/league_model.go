package models

import "github.com/sunp13/dbtool"

type leagueModel struct {
}

// GetList
func (m *leagueModel) GetList() (res []map[string]interface{}, err error) {
	sql := `
	select t1.*,
		 t2.dic_name as league_name_cn
	from b365api.league_info t1
	left join b365_dics t2 on (t2.dic_type = 'LEAGUE_NAME' and t1.league_name = t2.dic_value)
	order by t1.league_name
	`
	res, err = dbtool.D.QuerySQL(sql, nil)
	return
}
