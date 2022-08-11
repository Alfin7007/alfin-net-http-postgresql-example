package data

import (
	"database/sql"
	"http/example/features/articles"
)

type psqlArticleRepo struct {
	db *sql.DB
}

func NewArticleRepo(conn *sql.DB) articles.Data {
	return &psqlArticleRepo{
		db: conn,
	}
}

func (repo psqlArticleRepo) InsertData(artCore articles.Core) error {
	artModel := fromCore(artCore)
	query := `insert into articles (title, detail, user_id) values ($1, $2, $3)`
	sqlStatement, sqlErr := repo.db.Prepare(query)
	if sqlErr != nil {
		return sqlErr
	}

	exeResult, exeErr := sqlStatement.Exec(artModel.Title, artModel.Detail, artModel.UserID)
	if row, _ := exeResult.RowsAffected(); row == 0 {
		return exeErr
	}
	return nil
}

func (repo psqlArticleRepo) SelectAll() ([]articles.Core, error) {
	var artModel []Article
	queryResult, queryErr := repo.db.Query(`select id, title, detail from articles`)
	if queryErr != nil {
		return []articles.Core{}, queryErr
	}
	for queryResult.Next() {
		data := Article{}
		errScan := queryResult.Scan(&data.ID, &data.Title, &data.Detail)
		if errScan != nil {
			return []articles.Core{}, errScan
		}
		artModel = append(artModel, data)
	}
	return toCoreList(artModel), nil
}

func (repo psqlArticleRepo) SelectData(id int) (articles.Core, error) {
	var artModel Article
	query := `select a.id, a.title, a.detail, u.id, u.name from articles a inner join users u on u.id = a.user_id where a.id = $1`
	sqlStatement, sqlErr := repo.db.Prepare(query)
	if sqlErr != nil {
		return articles.Core{}, sqlErr
	}
	queryResult, queryErr := sqlStatement.Query(id)
	if queryErr != nil {
		return articles.Core{}, queryErr
	}

	for queryResult.Next() {
		scanErr := queryResult.Scan(&artModel.ID, &artModel.Title, &artModel.Detail, &artModel.User.ID, &artModel.User.Name)
		if scanErr != nil {
			return articles.Core{}, scanErr
		}
	}

	return artModel.toCore(), nil
}
