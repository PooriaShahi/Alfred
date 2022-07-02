package operations

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetConfig(config map[string]string) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	for k, v := range config {
		viper.Set(k, v)
	}

	viper.WriteConfig()
	fmt.Println("Config set successfuly")
}
