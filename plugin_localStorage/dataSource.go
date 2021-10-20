package plugin_localStorage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB
)

type Data struct {
	Id     int    `json:"id" db:"id"`
	Key    string `json:"key" db:"key"`
	Value  string `json:"value" db:"value"`
	UserId string `json:"userId" db:"user_id"`
	Time   string `json:"time" db:"time"`
}

func init() {
}

func tableInit() {
	_, err := db.Exec(`
		create table if not exists data(
	id INTEGER
		constraint data_pk
			primary key autoincrement,
	key text not null,
	value text,
	userId text not null,
	time text not null)
	`)
	if err != nil {
		log.Errorln("创建数据库表失败")
		return
	}
}

func DbInit() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(dir)
	db, err = sql.Open("sqlite3", dir+"/tmp/data.db")
	if err != nil {
		return
	}
	tableInit()
}

func checkDb() {
	if db == nil {
		DbInit()
		return
	}
	err := db.Ping()
	if err != nil {
		DbInit()
		return
	}
}

func Insert(data Data) error {
	checkDb()

	_, err := db.Exec("insert into data (key, value, userId, time) VALUES (?,?,?,?)", data.Key, data.Value, data.UserId, data.Time)
	if err != nil {
		return err
	}
	return err
}

func Delete(id int) error {
	checkDb()

	_, err := db.Exec("delete from data where id=?", id)
	if err != nil {
		return err
	}
	return err
}

func Get(id int) (Data, error) {
	checkDb()
	data := Data{}
	row := db.QueryRow(`select * from data where id=?`, id)
	err := row.Scan(&data.Id, &data.Key, &data.Value, &data.UserId, &data.Time)
	if err != nil {
		return data, err
	}
	return data, err
}

func Query(userId int) ([]Data, error) {
	checkDb()

	rows, err := db.Query(`select * from data`)
	if err != nil {
		log.Errorln("数据库查询失败" + err.Error())
		return nil, err
	}
	var datas []Data
	for rows.Next() {
		data := Data{}
		err := rows.Scan(&data.Id, &data.Key, &data.Value, &data.UserId, &data.Time)
		if err != nil {
			log.Errorln("数据库结果解析失败" + err.Error())
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, err
}
