package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/huoxue1/leafBot"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

type Connect struct {
}

func DbInit() *Connect {
	dataSource, ok := leafBot.DefaultConfig.Datas["data_souce"]
	if !ok {
		return nil
	}
	var err error
	Db, err = sqlx.Open("mysql", dataSource.(string))
	if err != nil {
		log.Println("连接数据库失败")
	}
	return &Connect{}
}

func (con *Connect) Close() {
	err := Db.Close()
	if err != nil {
		log.Println("数据库关闭失败")
	}
}
