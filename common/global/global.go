package global

import (
	"shopping/common/config"

	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
)
