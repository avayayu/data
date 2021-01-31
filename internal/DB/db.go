package DB

import (
	"sync"

	"github.com/avayayu/micro/dao"
	"github.com/avayayu/micro/dao/drivers/mysql"
	"github.com/avayayu/quant_data/internal/configs"
)


var db dao.DAO
var once sync.Once


func newDatabase() {
	config := configs.GetConfigs()

	mysqlURL := config.Get("DB.mySQLURL")
	mysqlPORT := config.Get("DB.mySQLPORT")

	mysqlConfigs := &mysql.MysqlConfigs{}
	mysqlConfigs.URL = mysqlURL
	mysqlConfigs.Port = mysqlPORT
	mysqlConfigs.UserName = config.Get("DB.mysqlUserName")
	mysqlConfigs.Password = config.Get("DB.mysqlPassword")
	mysqlConfigs.DBName = config.Get("DB.mysqlDBName")

	mysqlDrivers := &mysql.MysqlDriver{
		Configs: mysqlConfigs,
	}
	db = dao.NewDatabase(mysqlDrivers)

	
}


