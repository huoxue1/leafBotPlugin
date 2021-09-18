package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/huoxue1/leafBot"
	"github.com/jmoiron/sqlx"
	"log"
)

type Connect struct {
	Db *sqlx.DB
}

func DbInit() *Connect {
	var err error
	db, err := sqlx.Open("mysql", leafBot.DefaultConfig.Datas["data_souce"].(string))
	if err != nil {
		log.Println("连接数据库失败")
	}
	return &Connect{Db: db}
}

func (con *Connect) Close() {
	err := con.Db.Close()
	if err != nil {
		log.Println("数据库关闭失败")
	}
}
