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

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	//新增数据库连接实例的上下文信息注册
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
