package account

import (
	sql2 "database/sql"
	"main/model"
	"os"
	"time"
)

type Account struct {
	Id int

	Username string
	Email    string
	Password string

	CreateOn  time.Time
	LastLogin sql2.NullTime
}

func GetByEmail(email string) *Account {
	database := model.Database{Driver: "postgres", Source: os.Getenv("POSTGRESQL_URL")}
	database.Connection()
	sql := "SELECT * FROM account WHERE email = $1"
	row := database.Db.QueryRow(sql, email)
	defer database.Close()

	var account Account
	err := row.Scan(&account.Id, &account.Username, &account.Email, &account.Password, &account.CreateOn, &account.LastLogin)
	if err == sql2.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}

	return &account
}

func New(account Account) bool {

	database := model.Database{Driver: "postgres", Source: os.Getenv("POSTGRESQL_URL")}
	database.Connection()
	sql := "INSERT INTO account(email, username, password, create_on) VALUES($1, $2, $3, $4)"
	_, err := database.Db.Exec(sql, account.Username, account.Email, account.Password, account.CreateOn)
	defer database.Close()

	if err != nil {
		panic(err)
		return false
	}

	return true
}

func GetAll() *[]Account {
	database := model.Database{Driver: "postgres", Source: os.Getenv("POSTGRESQL_URL")}
	database.Connection()
	sql := "SELECT * FROM account"
	rows, err := database.Db.Query(sql)
	defer database.Close()

	if err != nil {
		panic(err)
		return nil
	}

	var accounts []Account

	for rows.Next() {
		var account Account

		err := rows.Scan(&account.Id, &account.Username, &account.Email, &account.Password, &account.CreateOn, &account.LastLogin)
		if err != nil {
			panic(err)
			return nil
		}

		accounts = append(accounts, account)
	}

	return &accounts
}
