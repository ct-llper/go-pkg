package mysql

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"sync"
	"time"
)

var (
	Instance Connect
	once     sync.Once
	dbs      *Driver
	drivers  *Drivers
)

type Connect struct {
	config  map[string][]interface{}
	configs map[string]map[string][]interface{}
}

// Driver 连接池
type Driver struct {
	MasterDb []*gorm.DB
	SlaveDb  []*gorm.DB
}

// Drivers 多库
type Drivers struct {
	MasterDb map[string][]*gorm.DB
	SlaveDb  map[string][]*gorm.DB
}

// Config 单库初始化配置
func Config(config map[string][]interface{}) (c Connect) {
	once.Do(func() {
		// 初始化数据库配置
		Instance = Connect{config: config}
	})
	return Instance
}

// Init 单库初始化连接
func (c Connect) Init() {
	if len(c.config) > 0 {
		dbs = new(Driver)
		for driver, conf := range c.config {
			switch driver {
			case "master":
				for _, v := range conf {
					db, err := gorm.Open("mysql", v.(map[string]interface{})["dialect"])
					if err == nil {
						db.DB().SetMaxOpenConns(v.(map[string]interface{})["MaxOpenConnects"].(int))
						db.DB().SetMaxIdleConns(v.(map[string]interface{})["MaxIdleConnects"].(int))
						db.DB().SetConnMaxLifetime(v.(map[string]interface{})["ConnMaxLifetime"].(time.Duration))
						db.SingularTable(true)
						db.LogMode(true)
						dbs.MasterDb = append(dbs.MasterDb, db)
					} else {
						panic(err.Error())
					}
				}
			case "slave":
				for _, v := range conf {
					db, err := gorm.Open("mysql", v.(map[string]interface{})["dialect"])
					if err == nil {
						db.DB().SetMaxOpenConns(v.(map[string]interface{})["MaxOpenConnects"].(int))
						db.DB().SetMaxIdleConns(v.(map[string]interface{})["MaxIdleConnects"].(int))
						db.DB().SetConnMaxLifetime(v.(map[string]interface{})["ConnMaxLifetime"].(time.Duration))
						db.SingularTable(true)
						db.LogMode(true)
						dbs.SlaveDb = append(dbs.SlaveDb, db)
					} else {
						panic(err.Error())
					}
				}
			}
		}
	}
}

// master 主库
func (c *Connect) master(ctx context.Context, dbName string) (dbConnect *gorm.DB, err error) {
	var (
		master []*gorm.DB
	)
	if dbName == "default" {
		master = dbs.MasterDb
	} else {
		master = drivers.MasterDb[dbName]
	}
	var num int
	switch len(master) {
	case 1:
		num = 0
		dbConnect = master[num]
	default:
		if len(master) > 0 {
			rand.Seed(time.Now().UnixNano())
			num = rand.Intn(len(master))
			dbConnect = master[num]
		} else {
			return dbConnect, errors.New("Database does not exist Master")
		}
	}
	dbConnect.SetLogger(&GormLogger{ctx: ctx})
	if dbName == "default" {
		//默认表名应用任何规则
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return c.config["master"][num].(map[string]interface{})["TablePrefix"].(string) + defaultTableName
		}
	} else {
		// 默认表名应用任何规则
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return c.configs[dbName]["master"][num].(map[string]interface{})["TablePrefix"].(string) + defaultTableName
		}
	}
	return dbConnect, nil
}

// slave 从库
func (c *Connect) slave(ctx context.Context, dbName string) (dbConnect *gorm.DB, err error) {
	var (
		slave []*gorm.DB
		num   int
	)
	if dbName == "default" {
		slave = dbs.SlaveDb
	} else {
		slave = drivers.SlaveDb[dbName]
	}
	switch len(slave) {
	case 1:
		num = 0
		dbConnect = slave[num]
	default:
		if len(slave) > 0 {
			rand.Seed(time.Now().UnixNano())
			num = rand.Intn(len(slave))
			dbConnect = slave[num]
		} else {
			return dbConnect, errors.New("Database does not exist slave")
		}
	}
	dbConnect.SetLogger(&GormLogger{ctx: ctx})
	if dbName == "default" {
		//默认表名应用任何规则
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return c.config["master"][num].(map[string]interface{})["TablePrefix"].(string) + defaultTableName
		}
	} else {
		//默认表名应用任何规则
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return c.configs[dbName]["master"][num].(map[string]interface{})["TablePrefix"].(string) + defaultTableName
		}
	}
	return dbConnect, nil
}

// Configs 多库初始化配置
func Configs(config map[string]map[string][]interface{}) (c Connect) {
	once.Do(func() {
		//初始化数据库配置
		Instance = Connect{configs: config}
	})
	return Instance
}

// Inits 多库初始化
func (c Connect) Inits() {
	//MasterDb map[string][]*gorm.DB
	if len(c.configs) < 1 {
		panic("No configuration database")
	}
	drivers = new(Drivers)
	drivers.MasterDb = make(map[string][]*gorm.DB, 0)
	drivers.SlaveDb = make(map[string][]*gorm.DB, 0)
	for dbName, conf := range c.configs {
		for ms, confDbs := range conf {
			switch ms {
			case "master":
				for _, v := range confDbs {
					db, err := gorm.Open("mysql", v.(map[string]interface{})["dialect"])
					if err == nil {
						db.DB().SetMaxOpenConns(v.(map[string]interface{})["MaxOpenConnects"].(int))
						db.DB().SetMaxIdleConns(v.(map[string]interface{})["MaxIdleConnects"].(int))
						db.DB().SetConnMaxLifetime(v.(map[string]interface{})["ConnMaxLifetime"].(time.Duration))
						db.SingularTable(true)
						db.LogMode(true)
						drivers.MasterDb[dbName] = append(drivers.MasterDb[dbName], db)
					} else {
						panic(err.Error())
					}
				}
			case "slave":
				for _, v := range confDbs {
					db, err := gorm.Open("mysql", v.(map[string]interface{})["dialect"])
					if err == nil {
						db.DB().SetMaxOpenConns(v.(map[string]interface{})["MaxOpenConnects"].(int))
						db.DB().SetMaxIdleConns(v.(map[string]interface{})["MaxIdleConnects"].(int))
						db.DB().SetConnMaxLifetime(v.(map[string]interface{})["ConnMaxLifetime"].(time.Duration))
						db.SingularTable(true)
						db.LogMode(true)
						drivers.SlaveDb[dbName] = append(drivers.SlaveDb[dbName], db)
					} else {
						panic(err.Error())
					}
				}
			}
		}
	}
}

// GetDB 获取数据库
func (c *Connect) GetDB(isMaster bool, ctx context.Context) (db *gorm.DB, err error) {
	if isMaster {
		db, err = c.master(ctx, "default")
	} else {
		db, err = c.slave(ctx, "default")
	}
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

// GetDBs 获取数据库
func (c *Connect) GetDBs(isMaster bool, dbName string, ctx context.Context) (db *gorm.DB, err error) {
	if dbName == "" {
		panic("No database selected")
	}
	if isMaster {
		db, err = c.master(ctx, dbName)
	} else {
		db, err = c.slave(ctx, dbName)
	}
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

// GormLogger gorm 日志
type GormLogger struct {
	ctx context.Context
}

func (g *GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logx.WithContext(g.ctx).Info("sql:", v[1:])
	case "log":
		logx.WithContext(g.ctx).Info("log:", v)
	}
}
