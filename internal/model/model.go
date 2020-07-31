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
	"time"
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

//通过回调方法对公共字段进行处理
//新增行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		//scope.FieldByName获取当前是否包含所需的字段
		if createTimeField, ok := scope.FieldByName("CreateOn"); ok {
			//值是否为空
			if createTimeField.IsBlank {
				//若为空则设置
				//set 内部通过反射进行赋值
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeFiled, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeFiled.IsBlank {
				_ = modifyTimeFiled.Set(nowTime)
			}
		}
	}
}

//更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除行为的回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnfield := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnfield && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()

		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}