package constants

import "os"

type App struct {
	AwsDefaultRegion string
	CognitoClientId  string
	RedisHost        string
	RedisPort        string
	SecretKey        string
}

func NewApp() *App {
	return &App{
		AwsDefaultRegion: os.Getenv("AWS_DEFAULT_REGION"),
		CognitoClientId:  os.Getenv("COGNITO_CLIENT_ID"),
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		SecretKey:        os.Getenv("SECRET_KEY"),
	}
}
