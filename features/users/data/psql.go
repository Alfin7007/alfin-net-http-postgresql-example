package data

import (
	"database/sql"
	"errors"
	"fmt"
	"http/example/features/users"
)

type psqlUserRepo struct {
	db *sql.DB
}

func NewUserRepo(conn *sql.DB) users.Data {
	return &psqlUserRepo{
		db: conn,
	}
}

func (repo psqlUserRepo) InsertUser(userCore users.Core) (err error) {
	repo.db.Begin()
	query := (`INSERT INTO users (name, email, password) VALUES ($1,$2,$3)`)
	statement, errPrepare := repo.db.Prepare(query)
	if errPrepare != nil {
		return errPrepare
	}
	result, err := statement.Exec(userCore.Name, userCore.Email, userCore.Password)
	fmt.Println(err)
	if err != nil {
		return errors.New("error query" + err.Error())
	}
	if row, _ := result.RowsAffected(); row == 0 {
		return errors.New("register fail")
	}
	defer repo.db.Close()

	return nil
}

func (repo psqlUserRepo) FindUser(email string) (userCore users.Core, err error) {
	repo.db.Begin()
	userModel := User{}

	query := `select id, email, password from users where email = $1`
	sqlStatement, errPrepare := repo.db.Prepare(query)
	if errPrepare != nil {
		fmt.Println(errPrepare.Error() + " prepare")
		return userCore, errPrepare
	}
	sqlResult, errSQL := sqlStatement.Query(email)
	if errSQL != nil {
		fmt.Println(errSQL.Error() + " sql")
		return userCore, errSQL
	}

	for sqlResult.Next() {
		errScan := sqlResult.Scan(&userModel.ID, &userModel.Email, &userModel.Password)
		if errScan != nil {
			return userCore, errScan
		}
	}
	if userModel.ID == 0 {
		return userCore, errors.New("not found")
	}

	userCore = userModel.ToCore()
	return userCore, nil
}

func (repo psqlUserRepo) SelectUser(id int) (userCore users.Core, err error) {
	userModel := User{}

	query := fmt.Sprintf(`select id, name, email from users where id = %d`, id)
	sqlResult, errSQL := repo.db.Query(query)
	if errSQL != nil {
		return userCore, errSQL
	}

	for sqlResult.Next() {
		errScan := sqlResult.Scan(&userModel.ID, &userModel.Name, &userModel.Email)
		if errScan != nil {
			return userCore, errScan
		}
	}
	userCore = userModel.ToCore()
	return userCore, nil
}
