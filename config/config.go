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

type MidtransConfig struct {
	ServerKey      string     `mapstructure:"SERVERKEY"`
	ClientKey      string     `mapstructure:"CLIENTKEY"` 
	Env            int        `mapstructure:"ENV"`
	URLHandler     string     `mapstructure:"EXP"` 
	Unit           string     `mapstructure:"UNIT"`
} 

type NSQConfig struct {
	Host      string    `mapstructure:"HOST"`
	Port      string    `mapstructure:"PORT"` 
	Topic     string    `mapstructure:"TOPIC"`
	Topic2    string    `mapstructure:"TOPIC2"` 
	Topic3    string    `mapstructure:"TOPIC3"` 
	Topic4    string    `mapstructure:"TOPIC4"` 
	Topic5    string    `mapstructure:"TOPIC5"`
	Topic6    string    `mapstructure:"TOPIC6"`
	Topic7    string    `mapstructure:"TOPIC7"`
	Topic8    string    `mapstructure:"TOPIC8"` 
	Topic9    string    `mapstructure:"TOPIC9"`
	Topic10   string    `mapstructure:"TOPIC10"`
	Topic11   string    `mapstructure:"TOPIC11"`
	Topic12   string    `mapstructure:"TOPIC12"`
	Topic13   string    `mapstructure:"TOPIC13"`
	Topic14   string    `mapstructure:"TOPIC14"`
}

type Configuration struct {
	Port    string                    `mapstructure:"Port"`
	Database struct {  
			Driver             string `mapstructure:"Driver"`
			Host               string `mapstructure:"Host"`
			Name               string `mapstructure:"Name"`
			Port               string `mapstructure:"Port"`
			Username           string `mapstructure:"Username"`
			Password           string `mapstructure:"Password"`
			JwtKey             string `mapstructure:"JwtKey"`
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
	
		JWTKey                   = app.Database.JwtKey
		AWS_REGION               = app.Database.AWS_REGION
		ACCESS_KEY_ID            = app.Database.ACCESS_KEY_ID
		ACCESS_KEY_SECRET        = app.Database.ACCESS_KEY_SECRET

		return &app
}
