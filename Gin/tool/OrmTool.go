package tool

import (
	"Go-Learning/Gin/model"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type OrmTool struct {
	*xorm.Engine
}

var DBEngine *OrmTool

func InitOrmEngine(cfg *Config) (*OrmTool, error) {
	dataBase := cfg.DataBase
	connStr := dataBase.User + ":" + dataBase.Password + "@tcp(" + dataBase.Host + ":" + dataBase.Port + ")/" + dataBase.DBName + "?charset=" + dataBase.Charset
	// 创建数据库操作引擎
	engine, err := xorm.NewEngine(dataBase.Driver, connStr)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(dataBase.ShowSql)
	// sync2 优化后的,更普遍, 这里是将结构体转化为数据库中的表
	err = engine.Sync2(new(model.SmsCode), new(model.Member), new(model.FoodCategory))
	if err != nil {
		return nil, err
	}
	ormTool := new(OrmTool)
	// 将数据库操作引擎赋值给Orm对象中的变量
	ormTool.Engine = engine
	// 将创建好的ORM工具类赋值给全局变量
	DBEngine = ormTool
	return ormTool, nil
}
