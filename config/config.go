package config

import (
	"log"
	
	

	"github.com/spf13/viper"	
)


var (
	JWTKey                   = ""
	AWS_REGION               = ""
	ACCESS_KEY_ID            = ""
	ACCESS_KEY_SECRET        = ""
)

type Configuration struct {
	Port    string                    `mapstructure:"Port"`
	Database struct {  
			Driver             string `mapstructure:"Driver"`
			Host               string `mapstructure:"Host"`
			Name               string `mapstructure:"Name"`
			Port               string `mapstructure:"Port"`
			Username           string `mapstructure:"Username"`
			Password           string `mapstructure:"Password"`
			jwtKey             string `mapstructure:"jwtKey"`
			AWS_REGION         string `mapstructure:"AWS_REGION"`
			ACCESS_KEY_ID      string `mapstructure:"ACCESS_KEY_ID"`
			ACCESS_KEY_SECRET  string `mapstructure:"ACCESS_KEY_SECRET"`
	}       `mapstructure:"database"`
} 


var appConfig *Configuration 

func GetConfiguration() *Configuration {
	return InitConfiguration()
} 

func InitConfiguration() *Configuration {
	app := Configuration{}
	 
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
		viper.SetConfigType("env") 
		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil 
		} 
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
		}
	
		JWTKey                   = app.Database.jwtKey
		AWS_REGION               = app.Database.AWS_REGION
		ACCESS_KEY_ID            = app.Database.ACCESS_KEY_ID
		ACCESS_KEY_SECRET        = app.Database.ACCESS_KEY_SECRET

		return &app
}
