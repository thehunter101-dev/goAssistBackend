package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil{
		panic("Todo se fue a la vrg!")
	}
	fmt.Println("Thas Great!")
	return db
}
