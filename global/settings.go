/**
* @Author:zhoutao
* @Date:2020/7/30 下午8:11
 */

package global

import (
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
)

//将配置信息和应用程序关联起来
//全局变量的初始化是会随着应用程序的不断演进而不断变化的，这里展示的并不一定是最终结果
var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger //logger对象
	JwtSetting      *setting.JwtSettings
)
