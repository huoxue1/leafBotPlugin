package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/fan/utils/sql"
)

type Sign struct {
	QQ           int64 `json:"qq" db:"qq"`                       // qq号
	LastSign     int64 `json:"last_sign" db:"last_sign"`         // 上次签到时间戳
	ContinueSign int   `json:"continue_sign" db:"continue_sign"` // 连续签到
	Fraction     int64 `json:"fraction" db:"fraction"`           // 积分
}

var (
	db *sql.Sqlite
)

func UpdateFraction(qq int64, fraction int64) {
	haveContent(qq)
	s := new(Sign)
	s.QQ = qq
	_ = query(s)
	s.Fraction += fraction
	_ = update(*s)
}

func init() {
	Db := new(sql.Sqlite)
	Db.DBPath = "./config/sign.db"
	err := Db.Open()
	if err != nil {
		log.Errorln("打开数据库失败" + err.Error())
		return
	}
	db = Db
	createTable()
}

func createTable() {
	err := db.Create("sign", &Sign{})
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}

func add(sign Sign) error {
	err := db.Insert("sign", &sign)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return err
}

// 查询表里面是否存在用户信息，若不存在则插入
func haveContent(qq int64) {
	defer func() {
		recover()
	}()
	find := db.CanFind("sign", fmt.Sprintf("where qq=%d", qq))
	if !find {
		err := db.Insert("sign", &Sign{QQ: qq})
		if err != nil {
			return
		}
	}
}

func HavaContent(qq int64) {
	haveContent(qq)
}

func Query(sign2 *Sign) error {
	return query(sign2)
}

func query(sign *Sign) error {
	haveContent(sign.QQ)
	err := db.Find("sign", sign, fmt.Sprintf("where qq=%v", sign.QQ))
	if err != nil {
		return err
	}
	return err
}
func Update(sign2 Sign) error {
	return update(sign2)
}

func update(sign Sign) error {
	_, err := db.DB.Exec("update sign set continue_sign=?,last_sign=?,fraction=? where qq=?", sign.ContinueSign, sign.LastSign, sign.Fraction, sign.QQ)
	if err != nil {
		return err
	}
	return err
}

func delete(qq int64) {
	err := db.Del("sign", fmt.Sprintf("where qq=%d", qq))
	if err != nil {
		return
	}
}
