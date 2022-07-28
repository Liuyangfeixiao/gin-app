package global

/*用来存放一些项目启动时的变量*/
import (
	"gin-demo/config"
	"github.com/spf13/viper"
)

// Application attributes TODO
type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

var App = new(Application)
