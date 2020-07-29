/**
* @Author:zhoutao
* @Date:2020/7/29 下午8:42
 */

package model

type Model struct {
	ID         uint32 `json:"id" gorm:"primary_key"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  string `json:"created_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
