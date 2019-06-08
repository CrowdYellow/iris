package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
)

func DBConnect() *xorm.Engine {
	engine, err := xorm.NewEngine(DB_CONNECTION, DB_USERNAME+":"+DB_PASSWORD+"@/"+DB_DATABASE+"?charset=utf8")
	if err != nil {
		log.Print("open mysql failed,err = ", err)
	}

	// 名称映射 SnakeMapper支持struct为驼峰式命名，表结构为下划线命名之间的转换
	engine.SetTableMapper(core.SnakeMapper{})

	return engine
}
