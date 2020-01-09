package repositories

import (
	"GraduationProject/models"
	"GraduationProject/utils"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"strings"
	"time"
	"xorm.io/core"
)

var (
	// mysql数据库-驱动
	engine *xorm.Engine
	// redis-连接客户端
	rgo *redis.Client
	// 配置文件句柄
	conf = utils.GetConfig()
)

func init() {
	//初始化Redis
	if err := InitRedis(); err != nil {
		panic("failed to connect redis:" + err.Error())
	}
	//连接MySQL出错
	if err := InitMySQL(); err != nil {
		panic("connect mysql errror:" + err.Error())
	}
}

//初始化连接MySQL
func InitMySQL() error {
	var err error
	engine, err = xorm.NewEngine("mysql", conf.Mysql.DSN)
	if err != nil {
		return err
	}
	//数据库连通性检测
	if err = engine.Ping(); err != nil {
		panic("failed to ping mysql database:" + err.Error())
	}
	//同步数据库结构
	if err = engine.Sync2(new(models.User), new(models.Issue), new(models.Message), new(models.Subject)); err != nil {
		return err
	}
	//用于设置最大打开的连接数，默认值为0表示不限制
	engine.SetMaxOpenConns(32)
	//SetMaxIdleConns用于设置闲置的连接数
	engine.SetMaxIdleConns(16)
	//设置本地时区
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	//是否开启调试
	if conf.Mysql.Debug {
		engine.Logger().SetLevel(core.LOG_DEBUG)
	}
	return nil
}

// 初始化连接Redis
func InitRedis() error {
	var err error
	var r string
	c := strings.Split(conf.Redis.DSN, "@")
	rgo = redis.NewClient(&redis.Options{
		Addr:     c[1],
		Password: c[0], // no password set
		DB:       0,    // use default DB
	})

	if r, err = rgo.Ping().Result(); err != nil {
		return err
	} else {
		log.Println("ping redis:", r)
	}
	return nil
}
