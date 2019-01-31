package main

import (
	"fmt"

	"github.com/nekottyo/test-go-orm/driver"
)

func main() {
	if err := driver.Gorm(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println()
	// if err := driver.Gorp(); err != nil {
	// 	fmt.Println(err)
	// }
}
