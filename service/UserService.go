package service

import (
	"log"
	"myapp/database"
	"myapp/graph/model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUser() ([]*model.User, error) {
	var user []*model.User

	db := database.Connect()
	defer db.Close()

	sql, err := db.Query("select uid, email, password from user")
	if err != nil {
		log.Print(err)
	}

	for sql.Next() {
		var users model.User
		sql.Scan(&users.ID, &users.Email, &users.Password)
		user = append(user, &users)
	}

	return user, nil
}
