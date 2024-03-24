package configs

import (
	appconfig "gin-gonic-gorm/configs/app_config"
	dbconfig "gin-gonic-gorm/configs/db_config"
)

func InitConfig()  {
	appconfig.InitAppConfig()
	dbconfig.InitDbConfig()
}
