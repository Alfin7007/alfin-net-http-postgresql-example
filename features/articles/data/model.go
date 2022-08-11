package data

import (
	"http/example/features/articles"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title  string `json:"name"`
	UserID uint   `json:"user_id"`
	Detail string `json:"detail"`
	User   User
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (art Article) toCore() articles.Core {
	return articles.Core{
		ArticleID: int(art.ID),
		Title:     art.Title,
		Detail:    art.Detail,
		User: articles.User{
			UserID: int(art.User.ID),
			Name:   art.User.Name,
		},
	}
}

func toCoreList(art []Article) []articles.Core {
	var coreList []articles.Core
	for _, val := range art {
		coreList = append(coreList, val.toCore())
	}
	return coreList
}

func fromCore(artCore articles.Core) Article {
	artModel := Article{
		Title:  artCore.Title,
		UserID: uint(artCore.UserID),
		Detail: artCore.Detail,
	}
	return artModel
}
