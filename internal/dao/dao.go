/**
* @Author:zhoutao
* @Date:2020/7/31 下午5:04
 */

package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{
		engine: engine,
	}
}
