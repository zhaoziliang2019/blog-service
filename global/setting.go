package global

import (
	"github.com/jinzhu/gorm"
	"github.com/zhaoziliang2019/blog-service/pkg/logger"
	"github.com/zhaoziliang2019/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	Appsetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSetting
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
)
