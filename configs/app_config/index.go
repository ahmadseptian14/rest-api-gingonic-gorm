package appconfig

import "os"

var PORT = ":8001"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"


func InitAppConfig()  {
	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		PORT = portEnv
	}

	staticRoute := os.Getenv("STATIC_ROUTE")
	if staticRoute != "" {
		STATIC_ROUTE = staticRoute
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir != "" {
		STATIC_DIR = staticDir
	}
}
