package driver

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/nekottyo/test-go-orm/model"
)

func Gorm() (err error) {
	db, err := gorm.Open("mysql", "root:hogehoge@tcp/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	users := []model.User{}

	db.Find(&users)

	fmt.Println("### gorm")
	fmt.Println("Find users")
	fmt.Println(users)

	db.Create(&model.User{Name: "Jinzhu", Age: 18})
	jinzhu := db.Where("name = ?", "Jinzhu").Find(&model.User{})

	fmt.Println(jinzhu)

	return nil
}
