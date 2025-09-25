package models

import "cfbapi/util"

func Migrate_Model() {
	util.DB.AutoMigrate(&Recruit{})
}
