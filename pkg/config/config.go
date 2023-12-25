package config

import (
	"errors"
	"os"
	"reflect"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var lock = &sync.RWMutex{}

var appConfig *viper.Viper

func Init(path string, name string) error {
	lock.Lock()
	defer lock.Unlock()
	// 如果文件夹不存在，则创建
	appConfig = viper.New()
	err := mkdirAll(path)
	if err != nil {
		return err
	}
	appConfig.AddConfigPath(path)
	appConfig.SetConfigName(name)
	appConfig.SetConfigType("yaml")
	err = appConfig.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			appConfig.Set("proxy", make(map[string]string))
			err := appConfig.SafeWriteConfig()
			if err != nil {
				log.Warnf("配置文件初始化错误，%v", err)
				return err
			}
		} else {
			log.Warnf("配置文件读取错误，%v", err)
			return err
		}
	}
	return nil
}

func mkdirAll(p string) error {
	if !isExist(p) {
		err := os.MkdirAll(p, os.ModePerm)
		return err
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
func Unmarshal(key string, v interface{}) error {
	pv := reflect.ValueOf(v)
	if pv.Kind() != reflect.Ptr {
		return errors.New("must be a ptr")
	}
	lock.RLock()
	defer lock.RUnlock()
	return appConfig.UnmarshalKey(key, v)
}

func GetStringMap(key string) map[string]interface{} {
	lock.RLock()
	defer lock.RUnlock()
	if ok := appConfig.IsSet(key); !ok {
		return nil
	}
	return appConfig.GetStringMap(key)
}
