package factory

import (
	"database/sql"
	userBussiness "http/example/features/users/bussiness"
	userData "http/example/features/users/data"
	userPresenter "http/example/features/users/presentation"
	// articleBussiness "myexample/go-gin/features/articles/bussiness"
	// articleData "myexample/go-gin/features/articles/data"
	// articlePresenter "myexample/go-gin/features/articles/presentation"
)

type Presenter struct {
	UserPresenter *userPresenter.UserHandler
	// ArticlePresenter *articlePresenter.ArticleHandler
}

func InitFactory(db *sql.DB) Presenter {
	varUserData := userData.NewUserRepo(db)
	varUserBussiness := userBussiness.NewUserBussiness(varUserData)
	varUserPresentation := userPresenter.NewUserHandler(varUserBussiness)

	// varArticleData := articleData.NewArticleRepo(db)
	// varArticleBussiness := articleBussiness.NewArticleBussiness(varArticleData)
	// varArticlePresentation := articlePresenter.NewArticleHandler(varArticleBussiness)

	return Presenter{
		UserPresenter: varUserPresentation,
		// ArticlePresenter: varArticlePresentation,
	}
}
