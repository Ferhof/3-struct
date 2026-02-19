package api

import "3-struct/config"

type apiClient struct {
	key config.Config
}

func (a apiClient) SetKey() {
	newConfig := config.NewConfig()
	a.key = config.Config{Key: newConfig.Load()}
}

func (a apiClient) GetKey() string {
	return a.key.Key
}
