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

	fmt.Println("### gorm")
	fmt.Println("-- Find all users")

	if err := db.Find(&users).Error; err != nil {
		return err
	}

	fmt.Println(users)

	fmt.Println("-- Find name =  'gorm'")
	if err := db.Create(&model.User{Name: "gorm", Age: 18}).Error; err != nil {
		return err
	}

	if err := db.Where("name = ?", "gorm").Find(&users).Error; err != nil {
		return err
	}

	fmt.Println(users)

	if err := db.Delete(model.User{}, "name = ?", "gorm").Error; err != nil {
		return err
	}

	return nil
}
