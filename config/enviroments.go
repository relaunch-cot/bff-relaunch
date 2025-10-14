package config

import "os"

var (
	PORT        = os.Getenv("PORT")
	IS_INSECURE = os.Getenv("IS_INSECURE")

	USER_MICROSERVICE_CONN = os.Getenv("USER_MICROSERVICE_CONN")
	CHAT_MICROSSEVICE_CONN = os.Getenv("CHAT_MICROSERVICE_CONN")
)
