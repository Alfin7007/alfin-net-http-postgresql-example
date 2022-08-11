package factory

import (
	"database/sql"
	articleBussiness "http/example/features/articles/bussiness"
	articleData "http/example/features/articles/data"
	articlePresenter "http/example/features/articles/presentation"
	userBussiness "http/example/features/users/bussiness"
	userData "http/example/features/users/data"
	userPresenter "http/example/features/users/presentation"
)

type Presenter struct {
	UserPresenter    *userPresenter.UserHandler
	ArticlePresenter *articlePresenter.ArticleHandler
}

func InitFactory(db *sql.DB) Presenter {
	varUserData := userData.NewUserRepo(db)
	varUserBussiness := userBussiness.NewUserBussiness(varUserData)
	varUserPresentation := userPresenter.NewUserHandler(varUserBussiness)

	varArticleData := articleData.NewArticleRepo(db)
	varArticleBussiness := articleBussiness.NewArticleBussiness(varArticleData)
	varArticlePresentation := articlePresenter.NewArticleHandler(varArticleBussiness)

	return Presenter{
		UserPresenter:    varUserPresentation,
		ArticlePresenter: varArticlePresentation,
	}
}
