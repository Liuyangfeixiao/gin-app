package bootstrap

import (
	"fmt"
	"gin-demo/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// InitConfig : 初始化配置
func InitConfig() *viper.Viper {
	// 设置配置环境路径
	config := "config.yaml"
	// 生产环境可以通过设置环境变量来改变文件配置路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 监听配置文件, 一旦配置文件改变则提醒并重载
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: ", in.Name)
		// 重载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
