package mysql

import (
	"strconv"
	"time"

	"go-port-and-adapter/systems/config"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	//driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.Load()
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MysqlConnect(configName string) *gorm.DB {
	mysql := viper.Sub("database." + configName)
	connection, err := gorm.Open("mysql", mysql.GetString("user")+":"+mysql.GetString("password")+"@tcp("+mysql.GetString("host")+":"+strconv.Itoa(mysql.GetInt("port"))+")/"+mysql.GetString("db_name")+"?charset=utf8&parseTime=True&loc=Local")
	// enable debug
	if err != nil {
		panic(err)
	}

	if mysql.GetBool("debug") {
		return connection.Debug()
	}

	pool := viper.Sub("database." + configName + ".pool")
	connection.DB().SetMaxOpenConns(pool.GetInt("maxOpenConns"))
	connection.DB().SetMaxIdleConns(pool.GetInt("maxIdleConns"))
	connection.DB().SetConnMaxLifetime(pool.GetDuration("maxLifetime") * time.Second)

	return connection
}
