package model

import (
	"errors"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	_ "github.com/glebarez/sqlite"
	_ "gorm.io/driver/mysql"
)

type Sign struct {
	QQ           int64 `json:"qq" db:"qq" gorm:"primaryKey"`     // qq号
	LastSign     int64 `json:"last_sign" db:"last_sign"`         // 上次签到时间戳
	ContinueSign int   `json:"continue_sign" db:"continue_sign"` // 连续签到
	Fraction     int64 `json:"fraction" db:"fraction"`           // 积分
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var (
	db *gorm.DB
)

func UpdateFraction(qq int64, fraction int64) {
	haveContent(qq)
	s := new(Sign)
	s.QQ = qq
	_ = query(s)
	s.Fraction += fraction
	_ = update(*s)
}

func InitDb(driver string, dsl string) (err error) {

	switch driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsl), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsl), &gorm.Config{})
	default:
		err = errors.New("不支持的数据库类型")
	}
	_ = db.AutoMigrate(&Sign{})
	return
}

func add(sign Sign) error {
	return db.Save(&sign).Error
}

// 查询表里面是否存在用户信息，若不存在则插入
func haveContent(qq int64) {
	db.Where(&Sign{QQ: qq}).FirstOrCreate(&Sign{QQ: qq, Fraction: 10, LastSign: time.Now().AddDate(0, 0, -1).Unix(), ContinueSign: 0})
}

func HavaContent(qq int64) {
	haveContent(qq)
}

func Query(sign2 *Sign) error {
	return query(sign2)
}

func query(sign *Sign) error {
	haveContent(sign.QQ)
	return db.Find(sign).Error
}
func Update(sign2 Sign) error {
	return update(sign2)
}

func update(sign Sign) error {
	return db.Save(&sign).Error
}

func delete(qq int64) {
	db.Table("sign").Where("qq = ?", qq).Delete(&Sign{})
}
