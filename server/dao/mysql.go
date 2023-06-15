package dao

import (
	"fmt"

	entity "server/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type conf struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname'`
	Port     string `yaml:"port"`
}

var SqlConnection *gorm.DB

func InitMysql() (db *gorm.DB) {
	viper.SetConfigFile("resource/mysql.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	username := viper.GetString("username")
	password := viper.GetString("password")
	url := viper.GetString("url")
	port := viper.GetInt("port")
	dbname := viper.GetString("dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		url,
		port,
		dbname,
	)

	sqlConnection, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}

	sqlConnection.AutoMigrate(entity.User{})
	SqlConnection = sqlConnection
	return sqlConnection
}

// 关闭数据库连接
func Close() {
	SqlConnection.Close()
}
