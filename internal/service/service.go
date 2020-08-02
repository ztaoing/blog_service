/**
* @Author:zhoutao
* @Date:2020/7/31 下午5:30
 */

package service

import (
	"blog_service/global"
	"blog_service/internal/dao"
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	svs := Service{ctx: ctx}
	//新增数据库连接实例的上下文信息注册
	svs.dao = dao.New(otgorm.WithContext(svs.ctx, global.DBEngine))
}
