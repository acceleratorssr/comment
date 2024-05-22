package flag

import (
	"comment/global"
	"comment/models"
)

func Makemigrations() {
	// 生成表结构
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.CommentSubjectModels{},
			&models.CommentIndexModels{},
			&models.CommentContentModels{},
			&models.UserModels{},
		)
	if err != nil {
		global.Log.Errorf("Makemigrations fail:%s", err)
		return
	}
	global.Log.Info("Makemigrations success")
}
