package initialize

import (
	"fmt"
	"github.com/akazwz/gin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// InitViper
// 初始化读取配置文件
func InitViper() (config *viper.Viper) {
	if gin.Mode() == "debug" {
		config = viper.New()
		config.AddConfigPath("./")
		config.SetConfigName("config")
		config.SetConfigType("yaml")
		if err := config.ReadInConfig(); err != nil {
			panic(err)
			return nil
		}

		/* 读取到全局变量 CONF中 */
		if err := config.Unmarshal(&global.CONF); err != nil {
			panic(err)
			return nil
		}

		config.WatchConfig()
		config.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file updated:", e.Name)
			if err := config.Unmarshal(&global.CONF); err != nil {
				panic(err)
			}
		})
	}
	return
}
