package handler

import (
	"github.com/humbertodlacerda/gopportunities/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	Db = config.GetSQLite()
}
