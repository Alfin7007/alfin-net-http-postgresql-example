package response

import "http/example/features/articles"

type Article struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	User   User
}

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

type Articles struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func FromSingleCore(artCore articles.Core) Article {
	return Article{
		ID:     artCore.ArticleID,
		Title:  artCore.Title,
		Detail: artCore.Detail,
		User: User{
			UserID: artCore.User.UserID,
			Name:   artCore.User.Name,
		},
	}
}

func fromSliceCore(artCore articles.Core) Articles {
	return Articles{
		ID:     artCore.ArticleID,
		Title:  artCore.Title,
		Detail: artCore.Detail,
	}
}

func FromCoreList(artCore []articles.Core) []Articles {
	var articlesResponse []Articles
	for _, val := range artCore {
		articlesResponse = append(articlesResponse, fromSliceCore(val))
	}
	return articlesResponse
}
