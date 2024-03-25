package dao

import (
	"Oracle/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func initMysql() (userName string, password string, address string, dbName string, maxIdleConns int, maxOpenConns int) {
	pwd, _ := os.Getwd()
	viper.SetConfigName("DBConfig")
	viper.SetConfigType("yml")
	viper.AddConfigPath(pwd + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("读取数据库配置文件失败：", err)
		return
	}

	userName = viper.GetString("mysql.username")
	password = viper.GetString("mysql.password")
	dbName = viper.GetString("mysql.dbname")
	address = viper.GetString("mysql.host")
	maxIdleConns = viper.GetInt("mysql.max_idle_conn")
	maxOpenConns = viper.GetInt("mysql.max_open_conn")
	return
}

var DB *gorm.DB

func Init() {
	log.Println("Connecting Mysql database......")
	userName, passWord, address, database, maxIdleConns, maxOpenConns := initMysql()
	dsn := userName + ":" + passWord + "@" + "(" + address + ")/" + database + "?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	configDB, err := db.DB()
	if err != nil {
		logrus.Error(err)
		return
	}
	configDB.SetMaxIdleConns(maxIdleConns)
	configDB.SetMaxOpenConns(maxOpenConns)
	configDB.SetConnMaxIdleTime(5 * time.Minute)
	configDB.SetConnMaxLifetime(5 * time.Minute)
	DB = db
	if err != nil {
		logrus.Error("数据库连接失败")
		return
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(model.Data{}, model.Users{})
	if err != nil {
		logrus.Fatal("failed to auto migrate")
	}

	log.Println("Connect Mysql success !")

}
