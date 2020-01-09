package utils

//该处定义了如何从根目录下获取并解析配置信息

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

//读取配置文件信息
type E struct {
	Environments `yaml:"environments"`
}

//请手动添加结构来实现yaml的解析
type Environments struct {
	Debug  bool   `yaml:"debug"`  //是否开启debug模式
	Server string `yaml:"server"` //服务运行的host:port
	Jwt    Jwt    `yaml:"jwt"`    //jwt签名配置
	User   User   `yaml:"user"`   //用户相关配置
	Email  Email  `yaml:"email"`  //邮件配置相关
	Mysql  Driver `yaml:"mysql"`  //mysql数据库配置
	Redis  Driver `yaml:"redis"`  //redis数据库配置
}

// Driver 数据库驱动链接
type Driver struct {
	DSN   string `yaml:"dsn"`   //数据库连接的源名称
	Debug bool   `yaml:"debug"` //开启数据库debug模式
}

// User 用户配置项
type User struct {
	ResetExpire int64 `yaml:"reset_expires"`	//重置密码过期时间
}

// Jwt 认证配置
type Jwt struct {
	SignKey     string `yaml:"sign_key"`
	SignMethod  string `yaml:"sign_method"`
	SignIssuer  string `yaml:"sign_issuer"`
	SignSubject string `yaml:"sign_subject"`
	SignExpires int64  `yaml:"sign_expires"`
}

// Email 邮件配置相关
type Email struct {
	ServerHost string `yaml:"server_host"`   //邮件SMTP服务地址
	ServerPort int    `yaml:"server_port"`   //邮件SMTP端口地址
	UserEmail  string `yaml:"from_email"`    //发送者邮件地址
	Username   string `yaml:"from_user"`     //发送者昵称别名
	Password   string `yaml:"from_password"` //发送者密码
}

var conf *E
var once sync.Once

// GetConfig 返回一个全局配置文件的单例
func GetConfig() *E {
	once.Do(func() {
		conf = new(E)
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			//读取配置文件失败,停止执行
			panic("read config file error:" + err.Error())
		}
	})
	return conf
}
