package driver

import (
	"fmt"
	"time"

	lorem "github.com/drhodes/golorem"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp"
	"github.com/nekottyo/test-go-orm/model"
)

type Model struct {
	ID uuid.UUID `gorm:"primary_key;type:char(36)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	model.User
	Model

	Item   Item `gorm:"foreignkey:ItemID;association_foreignkey:Refer"`
	ItemID uuid.UUID
}

type Item struct {
	Model
	model.Item
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	scope.SetColumn("ID", uuid)
	return nil
}

func (u *Item) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	scope.SetColumn("ID", uuid)
	return nil
}

func setupDB() (db *gorm.DB, err error) {
	conn, err := gorm.Open("mysql", "root:hogehoge@tcp/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return conn, err
	}
	return conn, nil
}

func Gorm() (err error) {
	db, err := setupDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}, &Item{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Item{})

	fmt.Println("### gorm")
	for i := 0; i < 10; i++ {

		item := new(Item)
		item.Name = lorem.Sentence(5, 20)
		db.Create(&item)
		pp.Print(item)

		user := new(User)
		user.Name = lorem.Word(3, 20)
		user.Age = 10
		user.ItemID = item.ID

		db.Create(&user)
		db.Save(&user)
		db.Save(&item)
	}

	fmt.Println("\n\n-- Find all users with item")

	var users []User
	if err := db.First(&users).Error; err != nil {
		return err
	}
	pp.Print(users)
	fmt.Println("\n\n-- Find relation")

	db.LogMode(true)
	var user User
	if err := db.First(&user).Related(&user.Item).Error; err != nil {
		return err
	}
	db.LogMode(false)

	pp.Print(user)

	if err := db.Delete(model.User{}, "name = ?", "gorm").Error; err != nil {
		return err
	}

	return nil
}
