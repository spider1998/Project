package app

import (
	"Project/doit/entity"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	User            entity.User
	Article         entity.Article
	Content         entity.Content
	ArticleVersion  entity.ArticleVersion
	ArticleForward  entity.ArticleForward
	Log             entity.Log
	Operator        entity.Operator
	Sort            entity.Sort
	Comment         entity.Comment
	Reply           entity.Reply
	SysCron         entity.SysCron
	GenderStatistic entity.GenderStatistic
	Message         entity.Message
	Friend          entity.Friend
)

func Migrate(dsn string) error {
	db, err := gorm.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		Logger.Error().Err(err).Msg("DB connection error.")
		panic(err)
	}
	err = db.AutoMigrate(&User, &Article, &Content, &ArticleVersion, &ArticleForward,
		&Log, &Operator, &Sort, &SysCron, &Comment, &Reply, &GenderStatistic,
			&Message, &Friend).Error
	return err
}
