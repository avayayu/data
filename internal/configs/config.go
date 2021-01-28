package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/avayayu/quant_data/internal/utils"
	"gopkg.in/yaml.v2"
)

var configFilePath string
var config *configs
var configOnce sync.Once

func init() {
	configFilePath = "F://github.com/quant_data/configs/config.yaml"
}

const MYSQL = "mysql"
const MYSSQL = "mssql"

func SetConfigPath(path string) {
	configFilePath = path
}

//Configs 获取配置
type Configs interface {
	Get(string) string
	GetBool(string) bool
	GetUint(key string) uint
	GetInt(key string) int
}

type configs struct {
	LogPath             string              `yaml:"logPath" configs:"default:./sync.log;env"`
	DeviceServerConfigs DeviceServerConfigs `yaml:"deviceServerConfigs"`
	Server              ServerConfigs       `yaml:"server"`
	DB                  DBConfigs           `yaml:"db"`
	Redis               RedisConfigs        `yaml:"redis"`
	Sync                SyncConfigs         `yaml:"sync"`
}

type ServerConfigs struct {
	ServerPort string `yaml:"serverPort" configs:"default:4002;env"`
	LogLevel   string `yaml:"logLevel" configs:"default:Info;env"`
}

type DeviceServerConfigs struct {
	BrokerURL   string `yaml:"mosquittoURL" configs:"default:127.0.0.1;env"`
	BrokerPort  string `yaml:"brokerPort" configs:"default:1883;env"`
	FileRoot    string `yaml:"fileRoot" configs:"default:/mnt;env"`
	MaxFileSize string `yaml:"fileMaxSize" configs:"default:10485760;env"`
}

type DBConfigs struct {
	MySQLURL        string `yaml:"mysqlURL" configs:"default:127.0.0.1;env:mysqlURL"`
	MySQLPORT       string `yaml:"mysqlPORT" configs:"default:3306;env:mysqlPORT"`
	MysqlUserName   string `yaml:"mysqlUserName" configs:"default:root;env:mysqlURL"`
	MysqlPassword   string `yaml:"mysqlPassword" configs:"env:mysqlPassword"`
	MysqlDBName     string `yaml:"mysqlDBName" configs:"default:BRIS;env"`
	MongoURL        string `yaml:"mongoURL" configs:"default:192.168.100.128;env"`
	MongoPORT       string `yaml:"mongoPORT" configs:"default:27017;env"`
	MongoDBName     string `yaml:"mongoDBName" configs:"default:BRIS;env"`
	MongoDBUserName string `yaml:"mongodbUserName" configs:"default:root;env"`
	MongoDBPassword string `yaml:"mongoDBPassword" configs:"default:bfr123123;env"`
	HisURL          string `yaml:"hisURL" configs:"default:127.0.0.1;env:hisURL"`
	HisPORT         string `yaml:"hisPORT" configs:"default:3306;env:hisPORT"`
	HisUserName     string `yaml:"hisUserName" configs:"default:sa;env"`
	HisPassword     string `yaml:"hisPassword" configs:"env"`
	HisDBName       string `yaml:"HisDBName" configs:"env"`
	HisDBType       string `yaml:"dbType" configs:"default:mysql;env"`
}

type SyncConfigs struct {
	Working               bool        `yaml:"working" configs:"default:false"`
	TimeFormat            string      `yaml:"timeFormat" configs:"default:2006010215:04:05"`
	StopLongAdvice        bool        `yaml:"stopLongAdvice" configs:"default:true"`
	SplitTreatment        bool        `yaml:"splitTreatment" configs:"default:true"`
	IgnoreTemporaryAdvice bool        `yaml:"ignoreTemporaryAdvice" configs:"default:true"`
	AdviceChangeOutFlag   bool        `yaml:"adviceChangeOutFlag" configs:"default:true"`
	RemoteVisitMode       string      `yaml:"remoteVisitMode" configs:"default:view"`
	AutoRecord            bool        `yaml:"autoMaticRecord" configs:"default:false"` //床旁补记账 直接写入执行表
	View                  ViewsConfig `yaml:"view"`
}

type RedisConfigs struct {
	URL  string `yaml:"url" configs:"default:127.0.0.1;env"`
	Port string `yaml:"port" configs:"default:6379;env"`
	DB   string `yaml:"db" configs:"default:0"`
}

type ViewsConfig struct {
	OutPatientView string `yaml:"outPatientView" configs:"default:V_OUTPATIENT"`
	InPatientView  string `yaml:"inPatientView" configs:"default:V_INPATIENT"`
	InAdviceView   string `yaml:"inAdviceView" configs:"default:V_ADVICE"`
	OutAdviceView  string `yaml:"outAdviceView" configs:"default:V_ADVICE"`
	UserView       string `yaml:"userView" configs:"default:V_USER"`
	DepartmentID   string `yaml:"departmentID" configs:"default:23"`
	TreatmentView  string `yaml:"treatmentView" configs:"default:V_TREATMENT"`
}

func getDefault(v interface{}) {

	tp := reflect.TypeOf(v)

	if tp.Kind() != reflect.Ptr {
		panic("v must be ptr")
	}

	tp = tp.Elem()
	rVal := reflect.ValueOf(v).Elem()
	for i := 0; i < tp.NumField(); i++ {
		t := tp.Field(i)
		f := rVal.Field(i)
		if t.Type.Kind() == reflect.Struct {
			getDefault(f.Addr().Interface())
		}
		// 得到tag中的字段名
		configsSettings := t.Tag.Get("configs")

		settingsArr := strings.Split(configsSettings, ";")

		for _, setting := range settingsArr {
			if strings.Contains(setting, "default") {
				defaultValue := strings.ReplaceAll(setting, "default:", "")

				switch t.Type.Kind() {
				case reflect.Uint:
					data, _ := strconv.Atoi(defaultValue)
					f.SetUint(uint64(data))
				case reflect.String:
					f.Set(reflect.ValueOf(defaultValue))
				case reflect.Bool:
					if defaultValue == "true" {
						f.SetBool(true)
					} else {
						f.SetBool(false)
					}
					break
				}
			}
		}
	}
}

//getFromEnv 从环境变量中获取配置 覆盖默认值 覆盖配置文件的值
func getFromEnv(v interface{}) {
	tp := reflect.TypeOf(v)
	rVal := reflect.ValueOf(v)

	if tp.Kind() != reflect.Ptr {
		panic("v must be ptr")
	}

	tp = tp.Elem()
	rVal = rVal.Elem()

	for i := 0; i < tp.NumField(); i++ {
		t := tp.Field(i)
		f := rVal.Field(i)

		configsSettings := t.Tag.Get("configs")

		if strings.Contains(configsSettings, "env") {
			//该变量需要读取环境变量的值
			arrs := strings.Split(configsSettings, ";")

			for _, arr := range arrs {
				if strings.Contains(arr, "env") {
					envSettings := strings.Split(arr, ":")
					var setting string
					if len(envSettings) > 1 {
						setting = os.Getenv(envSettings[1])
					} else {
						setting = utils.LcFirst(t.Name)
						setting = os.Getenv(setting)
					}
					if setting != "" {
						f.Set(reflect.ValueOf(setting))
					}

				}
			}
		}
	}

}

func getSubStruct(v interface{}, fieldName string) interface{} {
	rtp := reflect.TypeOf(v)
	rval := reflect.ValueOf(v)

	if rtp.Kind() == reflect.Ptr {
		rval = rval.Elem()
	}
	fieldNameTitle := strings.Title(fieldName)

	cdata := rval.FieldByName(fieldNameTitle)

	return cdata.Interface()

}

func (c *configs) GetUint(key string) uint {

	data, err := strconv.Atoi(c.Get(key))
	if err != nil {
		panic(err)
	}

	return uint(data)

}

func (c *configs) GetInt(key string) int {
	data, err := strconv.Atoi(c.Get(key))
	if err != nil {
		panic(err)
	}
	return data
}

func (c *configs) Get(key string) string {

	keyArr := strings.Split(key, ".")
	sub := interface{}(c)
	for index, key := range keyArr {
		sub = getSubStruct(sub, key)
		if sub == nil {
			panic("retreive settings Error")
		}
		if index == len(keyArr)-1 {
			//最后一个配置节点
			return sub.(string)
		}
	}
	return ""
}

func (c *configs) GetBool(key string) bool {

	keyArr := strings.Split(key, ".")
	sub := interface{}(c)
	for index, key := range keyArr {
		sub = getSubStruct(sub, key)
		if sub == nil {
			panic("retreive settings Error")
		}
		if index == len(keyArr)-1 {
			//最后一个配置节点
			return sub.(bool)
		}
	}
	return false
}

//ReadConfigs 从配置文件读取配置 同时合并环境变量
func readConfigs() Configs {

	file, err := os.Open(configFilePath)
	fmt.Printf("configPath is: %s ", configFilePath)
	defer file.Close()
	if os.IsNotExist(err) {
		config := configs{}
		getDefault(&config)

		data, err := yaml.Marshal(&config)

		if err != nil {
			panic(err)
		}
		path, _ := os.Getwd()
		path = filepath.Join(path, configFilePath)
		file, err = os.Create(path)
		if err != nil {
			panic(err)
		}
		file.Write(data)
		file.Close()
	}

	if err == nil {
		data, _ := ioutil.ReadAll(file)
		err = yaml.Unmarshal([]byte(data), &config)
		getFromEnv(&config.DB)
	}

	return config
}

func GetConfigs() Configs {
	configOnce.Do(func() {
		readConfigs()
	})
	return config
}
