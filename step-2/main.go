package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Nameee   string
	LastName string
	NickName string `gorm:"default:'dummy'"`
}

// `File` belongs to `User`, `UserID` is the foreign key
type File struct {
	gorm.Model
	UserID int
	User   User
	Name   string
	Md5    string
	format string `gorm:"default:'mp4'"`
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &File{})
	x := &User{Nameee: "jhonn_i"}
	db.FirstOrInit(x)
	x.LastName = "lastin"
	db.Save(x)
	// db.Create(x)

	db.Create(&User{Nameee: "new_med1n_c", NickName: "Custom Nicky1"})
	db.Create(&User{Nameee: "new_med2n_c"})
	db.Create(&User{Nameee: "new_med3n_c"})

	var users []User
	db.Find(&users)
	fmt.Println("================================")
	for _, e := range users {
		fmt.Printf("ID: %d  Name:%s %s (%s) CreatedAt %s Updated At %s", e.ID, e.Nameee, e.LastName, e.NickName, e.CreatedAt, e.UpdatedAt)
		fmt.Println("---")
	}
	fmt.Println("================================")
	// Delete - delete product
}
