/**
* @Author:zhoutao
* @Date:2020/7/30 下午7:55
 */

package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	//配置的相对路径
	//viper允许设置多个配置路径
	//vp.AddConfigPath("configs/")
	//vp.SetConfigType("yaml")
	//指定多配置文件路径
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	//热更新的监听和重新加载
	s.WatchSettingChange()
	return s, nil
}

//热更新的监听
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		//OnConfigChange的回调形参其实就是fsnotify
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			//更新后重新加载所有配置
			_ = s.ReloadAllSection()
		})
	}()
}
