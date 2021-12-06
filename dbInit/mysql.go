package dbInit

import (
	"bytes"
	"dog3pack/gos/env"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlInit mysql连接初始化
func MysqlInit(option env.MysqlOption) (*gorm.DB, error) {
	//连接
	var connStr bytes.Buffer
	connStr.WriteString(option.Account)
	connStr.WriteString(":")
	connStr.WriteString(option.Password)
	connStr.WriteString("@(")
	connStr.WriteString(option.Host)
	connStr.WriteString(":")
	connStr.WriteString(strconv.Itoa(option.Port))
	connStr.WriteString(")/")
	connStr.WriteString(option.Database)
	connStr.WriteString("?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local")

	var err error
	db, err := gorm.Open(mysql.Open(connStr.String()), &gorm.Config{})

	if db == nil {
		panic("默认[mysql.go]连接失败")
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(option.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(option.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。

	sqlDB.SetConnMaxLifetime(time.Duration(option.ConnMaxLifetime) * time.Second)

	//connections are not closed due to a connection's idle time.

	//sqlDB.SetConnMaxIdleTime(time.Duration(option.ConnMaxIdleTime) * time.Second)

	if err != nil {
		err := sqlDB.Close()
		if err != nil {
			return nil, err
		}
	}
	return db, err
}

// MysqlClose 断开Mysql连接
func MysqlClose(db *gorm.DB) {
	if db == nil {
		return
	}
	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	if err != nil {
		return
	}
	db = nil
}

// CheckMysqlRst Mysql执行结果校验
func CheckMysqlRst(rst *gorm.DB) {
	if rst.Error != nil {
		panic(rst.Error)
	}
}
