/**
* @Author:zhoutao
* @Date:2020/7/29 下午8:42
 */

package model

import (
	"blog_service/global"
	"blog_service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 注意需要引入 github.com/jinzhu/gorm/dialects/mysql
type Model struct {
	ID         uint32 `json:"id" gorm:"primary_key"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  string `json:"created_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(database *setting.DatabaseSettings) (*gorm.DB, error) {
	db, err := gorm.Open(database.DBtype, fmt.Sprintf("%s:%s@tcp(%s)?charset=%s&parseTime=%t&loc=Local",
		database.UserName,
		database.Passsword,
		database.Host,
		database.Charset,
		database.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	//模式
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(database.MaxIdleConns)
	db.DB().SetMaxOpenConns(database.MaxOpenConns)
	return db, nil
}
