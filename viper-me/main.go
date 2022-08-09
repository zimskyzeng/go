package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type User struct {
	Name 		string 				`json:"name"`
	Hobbies 	[]string			`json:"hobbies"`
	Clothing 	map[string]string 	`json:"clothing"`
	Age 		int 				`json:"age"`
	Eyes 		string				`json:"eyes"`
}

var (
	UserConfig User
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	err = viper.Unmarshal(&UserConfig)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", UserConfig)
}
