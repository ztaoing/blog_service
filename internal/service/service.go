/**
* @Author:zhoutao
* @Date:2020/7/31 下午5:30
 */

package service

import (
	"blog_service/global"
	"blog_service/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	return &Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}
