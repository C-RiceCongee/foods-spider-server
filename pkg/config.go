package pkg

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
Config 结构体实例的指针类型 （new返回结构体实例的指针类型）
*/
var Config = new(WebConfig)

type WebConfig struct {
	*App `mapstructure:"app"`
	// *ConsulConfig       `mapstructure:"consul"`
	// *TargetService      `mapstructure:"targetService"`
	*LogConfig          `mapstructure:"log"`
	*MysqlConfig        `mapstructure:"mysql"`
	*RedisConfig        `mapstructure:"redis"`
	*JwtConfig          `mapstructure:"jwt"`
	*EmailServiceConfig `mapstructure:"emailService"`
	*WeChat             `mapstructure:"wechat"`
}
type App struct {
	Name              string `mapstructure:"name"`
	Mode              string `mapstructure:"mode"`
	Port              int    `mapstructure:"port"`
	MachineID         int64  `mapstructure:"machineID"`
	OnlineTime        string `mapstructure:"onlineTime"`
	StaticBaseUrl     string `mapstructure:"staticBaseUrl"`
	StaticFileDirPath string `mapstructure:"staticFileDirPath"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Wait string `mapstructure:"wait"`
}
type TargetService struct {
	ServiceName string `mapstructure:"service_name"`
}
type LogConfig struct {
	Level         string `mapstructure:"level"`
	Filename      string `mapstructure:"filename"`
	CallerSkip    int    `mapstructure:"call_skip"`
	ErrorFilename string `mapstructure:"err_filename"`
	MaxSize       int    `mapstructure:"max_size"`
	MaxAge        int    `mapstructure:"max_age"`
	MaxBackups    int    `mapstructure:"max_backups"`
	Compress      bool   `mapstructure:"compress"`
	LocalTime     bool   `mapstructure:"local_time"`
}
type MysqlConfig struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	DbName        string `mapstructure:"db_name"`
	Charset       string `mapstructure:"charset"`
	ParseTime     bool   `mapstructure:"parse_time"`
	MaxOpenConnes int    `mapstructure:"max_open_connes"`
	MaxIdleConnes int    `mapstructure:"max_idle_connes"`
	LocalTime     string `mapstructure:"local_time"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
}

type JwtConfig struct {
	TokenExpireDuration int    `mapstructure:"token_expire_duration"`
	Secret              string `mapstructure:"secret"`
	Issuer              string `mapstructure:"issuer"`
}

type EmailServiceConfig struct {
	From     string `mapstructure:"from"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type WeChat struct {
	AppId        string   `mapstructure:"appid"`
	Secret       string   `mapstructure:"secret"`
	SendTemplate []string `mapstructure:"send_template"`
}

func handleUnmarshal() {
	err := viper.Unmarshal(Config)
	if err != nil {
		panic("The configuration serialization failure ~ ಥ_ಥ : " + err.Error())
	}
}
func InitConfig() {
	var configDir string
	var env string
	flag.StringVar(&configDir, "dir", "", "-dir to specify the configuration file directory")
	flag.StringVar(&env, "env", "debug", "-env to specify the running env debug | release")
	flag.Parse()
	if len(configDir) == 0 {
		fmt.Println("Run -dir to specify the configuration file directory ")
		os.Exit(0)
	}
	if len(env) == 0 {
		fmt.Println("Run -env to specify the running env  debug | release")
		os.Exit(0)
	}
	/*
		TODO
		根据启动环境
		动态写配置文件路径！
		动态传入配置文件名称！
	*/
	fmt.Println(configDir)
	viper.AddConfigPath(configDir)
	if env == "debug" {
		viper.SetConfigName("webappDevConfig") // 配置文件名称(无扩展名)
	} else {
		viper.SetConfigName("webappProdConfig") // 配置文件名称(无扩展名)
	}
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	handleUnmarshal()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf(fmt.Sprintf("Config file changed:%s", in.Name))
		handleUnmarshal()
	})
}
