package driver

import (
	"database/sql"
	"fmt"

	gorp "gopkg.in/gorp.v2"

	lorem "github.com/drhodes/golorem"
	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	"github.com/nekottyo/test-go-orm/model"
)

func Gorp() (err error) {
	db, err := sql.Open("mysql", "root:hogehoge@tcp/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF-8"}}

	fmt.Println("\n### gorp")

	// Model と table 名を紐付ける必要がある
	dbmap.AddTableWithName(model.User{}, "users").SetKeys(true, "ID")
	dbmap.AddTableWithName(model.Item{}, "items").SetKeys(true, "ID")
	defer dbmap.Db.Close()

	for i := 0; i < 10; i++ {
		item := new(model.Item)
		item.Name = lorem.Sentence(5, 20)
		dbmap.Insert(&item)

		user := new(model.User)
		user.Name = lorem.Word(3, 20)
		user.Age = 10
		pp.Print(user)

		dbmap.Insert(&user)
	}

	var users []User
	fmt.Println("-- Find all users")
	if _, err = dbmap.Select(&users, "select * from users"); err != nil {
		return err
	}
	if err != nil {
		return err
	}
	fmt.Println(users)

	var user User
	fmt.Println("-- Find first user")
	err = dbmap.SelectOne(&user, "SELECT * FROM users", 1)
	if err != nil {
		return err
	}
	pp.Print(user)

	dbmap.Exec("delete from users where name=?", "gorp")
	return nil
}
