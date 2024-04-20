package conf

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// 是否开发环境
var dev = os.Getenv("env") == "dev"

// 解析完成的配置文件
var allConf map[string]interface{}

// {功能} conf 初始化
// {参数} 无
// {返回} 无
func init() {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/app.yml")
	if err != nil {
		panic("配置文件 app.yml 不存在")
	}
	bytes, _ := io.ReadAll(file)
	err = yaml.Unmarshal(bytes, &allConf)
	if err != nil {
		panic("配置文件解析失败: " + err.Error())
	}
}

// {功能} 获取配置文件
// {参数} key 配置 key, defaultValue 默认值
// {返回} 配置信息
func GetConf(key string, defaultValue ...string) string {
	if dev {
		devConf := allConf["dev"]
		if devConf != nil {
			if value, ok := devConf.(map[interface{}]interface{})[key]; ok {
				return fmt.Sprint(value)
			}
		}
	} else {
		prodConf := allConf["prod"]
		if prodConf != nil {
			if value, ok := prodConf.(map[interface{}]interface{})[key]; ok {
				return fmt.Sprint(value)
			}
		}
	}
	if allConf[key] == nil {
		if defaultValue != nil {
			return defaultValue[0]
		}
		return ""
	}
	return fmt.Sprint(allConf[key])
}
