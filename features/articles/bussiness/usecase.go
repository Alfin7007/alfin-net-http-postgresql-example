package bussiness

import (
	"http/example/features/articles"
)

type articleData struct {
	artData articles.Data
}

func NewArticleBussiness(artData articles.Data) articles.Bussiness {
	return articleData{
		artData: artData,
	}
}
func (uc articleData) InsertArticle(artCore articles.Core) error {
	err := uc.artData.InsertData(artCore)
	return err
}

func (uc articleData) GetAll() ([]articles.Core, error) {
	result, err := uc.artData.SelectAll()
	return result, err
}

func (uc articleData) GetArticle(id int) (articles.Core, error) {
	result, err := uc.artData.SelectData(id)
	return result, err
}
