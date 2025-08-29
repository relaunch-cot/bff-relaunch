package config

import "os"

var (
	PORT        = os.Getenv("PORT")
	IS_INSECURE = os.Getenv("IS_INSECURE")

	USER_MICROSERVICE_CONN = os.Getenv("USER_MICROSERVICE_CONN")
)
