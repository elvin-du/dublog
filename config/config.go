package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	_conf interface{}
)

func init() {
	loadConf("config.yml")
}

func loadConf(path string) {
	f, err := os.Open(path)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	bin, err := ioutil.ReadAll(f)
	if nil != err {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(bin, &_conf)
	if nil != err {
		log.Fatalln(err)
	}
}

func Get(key string, result interface{}) error {
	keys := strings.Split(key, ":")

	c := _conf
	for _, key := range keys {
		if cm, ok := c.(map[interface{}]interface{}); ok {
			if c, ok = cm[key]; ok {
				continue
			}
			return errors.New("key: " + key + " not found")
		}

		return errors.New("not a map")
	}

	v := reflect.ValueOf(result)

	if v.Kind() != reflect.Ptr {
		return errors.New("Need Ptr")
	}

	v.Elem().Set(reflect.ValueOf(c))
	return nil
}
