package initlize

func InitALL() error {
	// 初始化日志包
	initLog()
	// 初始化数据库
	err := initDB()
	if err != nil {
		return err
	}
	return nil
}
