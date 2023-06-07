package config

import (
	"log"
	"os"
	

	"github.com/spf13/viper"
)


var (
	JWTKey                   = ""
	AWS_REGION               = ""
	ACCESS_KEY_ID            = ""
	ACCESS_KEY_SECRET        = ""
)

type Configuration struct {
	Port    string 
	Database struct {
		Driver                   string 
		Host                     string
		Name                     string
		Address                  string 
		Port                     string 
		Username                 string 
		Password                 string 
		jwtKey   				 string
	    AWS_REGION               string
	    ACCESS_KEY_ID            string
	    ACCESS_KEY_SECRET        string
	}
} 


var appConfig *Configuration 

func GetConfiguration() *Configuration {
	return InitConfiguration()
} 

func InitConfiguration() *Configuration {
	app := Configuration{}
	isRead := true 

	if val, found := os.LookupEnv("PORT"); found {
		app.Port = val 
		isRead = false
	}
	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.Database.jwtKey = val
		isRead = false
	} 
	if val, found := os.LookupEnv("DBUSER"); found {
		app.Database.Username = val 
		isRead = false
	}
	if val , found := os.LookupEnv("DBPASS"); found {
		app.Database.Password = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST") ; found {
		app.Database.Host = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		app.Database.Port = val 
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.Database.Name = val 
		isRead = false 
	}
	// untuk mencari env gambar
	if val, found := os.LookupEnv("AWS_REGION") ; found {
		app.Database.AWS_REGION = val 
		isRead = false 
	}
	if val, found := os.LookupEnv("ACCESS_KEY_ID"); found {
		app.Database.ACCESS_KEY_ID = val 
		isRead = false
	}
	if val, found := os.LookupEnv("ACCESS_KEY_SECRET"); found {
		app.Database.ACCESS_KEY_SECRET = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
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
	}
		JWTKey                   = app.Database.jwtKey
		AWS_REGION               = app.Database.AWS_REGION
		ACCESS_KEY_ID            = app.Database.ACCESS_KEY_ID
		ACCESS_KEY_SECRET        = app.Database.ACCESS_KEY_SECRET

		return &app
}