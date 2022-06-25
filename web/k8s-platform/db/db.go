package db

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"k8s-platform/config"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-10-28 13:07
 * @description
 */

var (
	isInit bool
	GORM   *gorm.DB
	err    error
)

//db的初始化函数，与数据库建立连接
func Init() {
	//判断是否已经初始化了
	if isInit {
		return
	}

	//组装连接配置
	//parseTime是查询结果是否自动解析为时间
	//loc是Mysql的时区设置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPwd,
		config.DbHost,
		config.DbPort,
		config.DbName)

	//与数据库建立连接，生成一个*gorm.DB类型的对象
	var DbType gorm.Dialector
	//DbType := *new(gorm.Dialector)
	if config.DbType == "mysql" {
		DbType = mysql.Open(dsn)
	} else {
		panic("config.DbType!=\"mysql\"")
	}
	GORM, err = gorm.Open(DbType, &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	fmt.Println("db", GORM)

	//开启连接池
	sqlDb, err := GORM.DB()
	if err != nil {
		logger.Error(errors.New("获取Deploymen列表失败？" + err.Error()))
		return
	}

	//连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDb.SetMaxIdleConns(config.MaxIdleConns)
	//设置了连接池最大数量
	sqlDb.SetMaxOpenConns(config.MaxOpenConns)
	//设置了连接可复用的最大时间
	sqlDb.SetConnMaxIdleTime(time.Duration(config.MaxLifeTime))
	isInit = true
	logger.Info("连接数据库成功！")
}

//db的关闭函数
func Close() error {
	sqlDB, err := GORM.DB()
	if err != nil {
		logger.Error(errors.New("获取Deploment列表失败？" + err.Error()))
		panic(err.Error())
	}
	return sqlDB.Close()
}
