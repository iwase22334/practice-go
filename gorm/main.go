package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Message struct {
	ID      int64
	UserID  string
	Message string
}

type MessageOmit struct {
	UserID  string
	Message string
}

func main() {
	dsn := "user:password@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to db", err)
	}

	{
		result := db.Table("table1").Create(&MessageOmit{UserID: "user4", Message: "and you?"})
		if result.Error != nil {
			fmt.Println("failed to fetch message", result.Error)
		}
	}

	{
		var table = []Message{}
		result := db.Table("table1").Find(&table)
		if result.Error != nil {
			fmt.Println("failed to fetch message", result.Error)
		}
		for _, row := range table {
			fmt.Println(row.ID, row.UserID, row.Message)
		}

	}

}
