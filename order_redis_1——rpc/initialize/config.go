package initialize

import (
	"OrderTakeaway/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config-debug.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}


	if err := v.Unmarshal(&global.ServerConfig); err != nil { //将配置项转为结构体 //需要事先定义结构体，且每项名称需要一致
		panic(err)
	}

	zap.S().Infof("配置信息：%v",global.ServerConfig)

	//动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		//fmt.Println("配置文件被更改",e.Name)
		//被修改后重新转换结构体
		_=v.ReadInConfig()
		_=v.Unmarshal(&global.ServerConfig)
		zap.S().Infof("修改后的配置信息：%v",global.ServerConfig)
	})


}