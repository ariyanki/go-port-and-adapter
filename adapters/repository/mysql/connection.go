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

//Database Database
type Database struct {
	Host      string
	User      string
	Password  string
	DBName    string
	DBNumber  int
	Port      int
	Pool 	  DbPool
	DebugMode bool
}

type DbPool struct {
	MaxOpenConns  int
	MaxIdleConns  int
	MaxLifetime  time.Duration
}

func init() {
	config.Load()
}

// LoadDBConfig load database configuration
func LoadDBConfig(name string) Database {
	db := viper.Sub("database." + name)
	conf := Database{
		Host:      db.GetString("host"),
		User:      db.GetString("user"),
		Password:  db.GetString("password"),
		DBName:    db.GetString("db_name"),
		Port:      db.GetInt("port"),
		Pool: 	   DbPool{
			MaxOpenConns:  db.GetInt("maxOpenConns"),
			MaxIdleConns:  db.GetInt("maxIdleConns"),
			MaxLifetime:  db.GetDuration("maxLifetime"),
		},
		DebugMode: db.GetBool("debug"),
	}
	return conf
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MysqlConnect(configName string) *gorm.DB {
	mysql := LoadDBConfig(configName)
	connection, err := gorm.Open("mysql", mysql.User+":"+mysql.Password+"@tcp("+mysql.Host+":"+strconv.Itoa(mysql.Port)+")/"+mysql.DBName+"?charset=utf8&parseTime=True&loc=Local")
	// enable debug
	if err != nil {
		panic(err)
	}

	if mysql.DebugMode {
		return connection.Debug()
	}

	connection.DB().SetMaxOpenConns(mysql.Pool.MaxOpenConns)
	connection.DB().SetMaxIdleConns(mysql.Pool.MaxIdleConns)
	connection.DB().SetConnMaxLifetime(mysql.Pool.MaxLifetime * time.Second)

	return connection
}
