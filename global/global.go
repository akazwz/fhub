package global

import (
	"github.com/akazwz/fhub/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GDB  *gorm.DB
	VP   *viper.Viper
	CONF config.Config
)
