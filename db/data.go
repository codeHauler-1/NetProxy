package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Sql *gorm.DB
var err error

func init() {
	if Sql, err = gorm.Open("sqlite3", "./edgeBox.sqlite"); err != nil {
		fmt.Println(err.Error())
	}
}

type BridgeInfo struct {
	ID   int64  `gorm:"column:id"`
	Vkey string `gorm:"column:vkey"`
}

func (b *BridgeInfo) IsExist() (bool, error) {
	var i []BridgeInfo
	ret := Sql.Where(b).Find(&i)
	if ret.Error != nil {
		return false, ret.Error
	}
	if len(i) < 1 {
		return false, nil
	}
	return true, nil
}

//db.DB.AutoMigrate(&db.Device{})