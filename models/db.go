package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	. "github.com/andyron/architchat/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
)

var Db *sql.DB // 数据库连接池

// 数据库连接初始化方法
func init() {
	var err error
	config := LoadConfig()
	driver := config.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", config.Db.User, config.Db.Password, config.Db.Address, config.Db.Database)
	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// 创建随机UUID，RFC 4122
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// 哈希加密
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
