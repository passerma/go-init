package model

import (
	"fmt"
	"go-init/src/conf"
	"go-init/src/log"
	"go-init/src/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库链接实例
var DdClient *gorm.DB

// {功能} 初始化表结构
// {参数} 无
// {返回} 无
func initTable() {
	// DdClient.AutoMigrate(&Blog{}, &User{}, &BlogComment{}, &UserInfo{}, &Message{})
}

// {功能} 初始化mysql链接
// {参数} 无
// {返回} 无
func initSQL() {
	sqlHost := conf.GetConf("sql_host")
	sqlPort := conf.GetConf("sql_port")
	sqlUser := conf.GetConf("sql_user")
	sqlPass := conf.GetConf("sql_pass")
	sqlDb := conf.GetConf("sql_db")
	sqlData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s", sqlUser, sqlPass, sqlHost, sqlPort, sqlDb)
	// 设置mysql日志等级
	loggerLevel := logger.Silent
	if util.IsDEV() {
		loggerLevel = logger.Info
	}
	// 连接mysql
	db, err := gorm.Open(mysql.Open(sqlData), &gorm.Config{
		Logger: logger.Default.LogMode(loggerLevel),
	})
	if err != nil {
		panic("[mysql] init error, open: " + sqlData + ", info: " + err.Error())
	} else {
		log.ComLoggerFmt("[mysql] init success: ", sqlData)
	}
	DdClient = db
}

// {功能} model 初始化
// {参数} 无
// {返回} 无
func Init() {
	initSQL()
	initTable()
}

// {功能} 关闭mysql链接
// {参数} 无
// {返回} 无
func Close() {
	log.ComLoggerClient.Info("[mysql] close")
	sqlDB, _ := DdClient.DB()
	sqlDB.Close()
}
