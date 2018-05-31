package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// DBConfigStruct DBConfig Struct
type DBConfigStruct struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}

// FabricConfigStruct FabricConfigStruct Struct
type FabricConfigStruct struct {
	APIURL       string
	APIAnchor    string
	APIChannel   string
	APIChaincode string
	APIAssetCC   string
}

// ServerConfigStruct ServerConfigStruct Struct
type ServerConfigStruct struct {
	Port        []string
	Mode        string
	GormLogMode string
}

// ObjConfigStruct ObjConfigStruct
type ObjConfigStruct struct {
	MsisdnType    string
	MsisdnVersion string
}

// DBConfig 数据库相关配置
var DBConfig DBConfigStruct

// FabricConfig Fabric Config
var FabricConfig FabricConfigStruct

// ServerConfig Server Config
var ServerConfig ServerConfigStruct

// ObjConfig ObjConfig
var ObjConfig ObjConfigStruct

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	DBConfig.Dialect = viper.GetString("database.dialect")
	DBConfig.Database = viper.GetString("database.database")
	DBConfig.User = viper.GetString("database.user")
	DBConfig.Password = viper.GetString("database.password")
	DBConfig.Host = viper.GetString("database.host")
	DBConfig.Port = viper.GetInt("database.port")
	DBConfig.Charset = viper.GetString("database.charset")
	DBConfig.MaxIdleConns = viper.GetInt("database.maxIdleConns")
	DBConfig.MaxOpenConns = viper.GetInt("database.maxOpenConns")

	DBConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)

	FabricConfig.APIURL = viper.GetString("fabric.api.url")
	FabricConfig.APIAnchor = viper.GetString("fabric.api.anchor")
	FabricConfig.APIChannel = viper.GetString("fabric.api.channel")
	FabricConfig.APIChaincode = viper.GetString("fabric.api.chaincode")
	FabricConfig.APIAssetCC = viper.GetString("fabric.api.asset_cc")

	ServerConfig.Port = strings.Split(viper.GetString("server.port"), ",")
	ServerConfig.Mode = viper.GetString("server.mode")
	ServerConfig.GormLogMode = viper.GetString("server.gorm.LogMode")

	ObjConfig.MsisdnType = viper.GetString("obj.msisdn.type")
	ObjConfig.MsisdnVersion = viper.GetString("obj.msisdn.version")
}
