package api

import "3-struct/config"

var key string

func SetKey() {
	newConfig := config.NewConfig()
	newConfig.Load()
	key = newConfig.Key
}
