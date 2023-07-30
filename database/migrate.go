package database

import "github.com/yahkerobertkertasnya/preweb/graph/model"

func MigrateTable() {
	db := GetInstance()

	err := db.AutoMigrate(&model.User{}, &model.Tweet{}, &model.Follow{}, &model.TweetLike{})
	if err != nil {
		panic(err)
	}
}
