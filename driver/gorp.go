package driver

import (
	"database/sql"
	"fmt"

	gorp "gopkg.in/gorp.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nekottyo/test-go-orm/model"
)

func Gorp() (err error) {
	db, err := sql.Open("mysql", "root:hogehoge@tcp/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF-8"}}

	// Model と table 名を紐付ける必要がある
	dbmap.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
	defer dbmap.Db.Close()

	users := []model.User{}

	fmt.Println("### gorp")
	fmt.Println("-- Find all users")
	if _, err = dbmap.Select(&users, "select * from users"); err != nil {
		return err
	}
	if err != nil {
		return err
	}
	fmt.Println(users)

	fmt.Println("-- Find name = 'gorp'")
	err = dbmap.Insert(&model.User{Name: "gorp", Age: 20})
	if err != nil {
		return err
	}

	var list []model.User
	_, err = dbmap.Select(&list, "select * from users where name=?", "gorp")
	if err != nil {
		return err
	}

	fmt.Println(list)

	dbmap.Exec("delete from users where name=?", "gorp")
	return nil
}
