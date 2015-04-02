package db

import (
	"dublog/config"
	"log"
)

func init() {
    log.SetFlags(log.Lshortfile)
	readConf()
}

func readConf() {
	var mysqlSpecs map[interface{}]interface{}
	err := config.Get("mysql", &mysqlSpecs)
	if nil != err {
		log.Fatalln(err)
	}
	log.Println(mysqlSpecs)

    err = startDB()
    if nil != err {
		log.Fatalln(err)
	}
}

func startDB()error{
    return nil
}
