package config

import(
	"fmt"
	"github.com/spf13/viper"
)

func ParseConfig() error {
	viper.SetConfigName("config") // 配置文件的名字（无后缀名）
	viper.AddConfigPath("./conf") // 多次调用以添加许多搜索路径
	viper.AddConfigPath(".") // 可选择在工作目录中查找配置
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	viper.WatchConfig()
	return nil
}