package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name     string
	NickName string `gorm:"default:'dummy'"`
	LastName string
	Color    string `gorm:"default:'white'"`
	Files    []File `gorm:"many2many:user_file;"`
}

// `File` belongs to `User`, `UserID` is the foreign key
type File struct {
	gorm.Model
	Name   string
	Md5    string
	Format string `gorm:"default:'gif'"`
	UserId uint
	Users  []User `gorm:"many2many:user_file;"`
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &File{})

	createUser(db, "mak11", "meddy1", "red")
	createUser(db, "mak12", "", "yellow")
	createUser(db, "mak13", "meddy3", "green")

	printAllUsers(db)

	create3File(db)
	printFilesInDb(db)

	var u1 User
	db.First(&u1, "Name = ?", "med1")
	var u2 User
	db.First(&u2, "Name = ?", "med2")
	var u3 User
	db.First(&u3, "Name = ?", "med3")

	var f1 File
	db.First(&f1, "Name = ?", "Fname1")
	var f2 File
	db.First(&f2, "Name = ?", "Fname2")
	var fx File
	db.First(&fx, "Name = ?", "fx")

	// u2.Files = append(u2.Files, f1, f2)
	// u3.Files = append(u3.Files, fx, f1)

	// db.Save(&u1)
	// db.Save(&u2)
	// db.Save(&u3)

	f1.Users = append(f1.Users, u1, u3)
	f2.Users = append(f2.Users, u1)
	fx.Users = append(f2.Users, u1, u2, u3, u3, u3)

	db.Save(&f1)
	db.Save(&f2)
	db.Save(&fx)

	// fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	// getUserRelatedToFile(db, "Filename2")
	getUserRelatedToFile(db, "Fname1")
	getUserRelatedToFile(db, "Fname2")
	getUserRelatedToFile(db, "fx")

	filesReslatedToUser(db, "med1")

}

func getUserRelatedToFile(db *gorm.DB, filename string) {
	fmt.Println("Getting Users Related File   ", filename)
	var f File
	var us []*User
	db.First(&f, "Name = ?", filename)
	db.Model(&f).Related(&us, "Users")
	for _, e := range us {
		fmt.Println(e.ID, e.Name)
	}
}

func filesReslatedToUser(db *gorm.DB, name string) {
	fmt.Println("Getting File Related To User  ", name)

	var fs []*File
	var u User
	db.First(&u, "Name = ?", name)

	db.Model(&u).Related(&fs, "Files")
	for _, e := range fs {
		fmt.Println(e.ID, e.Name)
	}
	fmt.Println("/////////////////////")
}

func printAllUsers(db *gorm.DB) {
	var users []User
	db.Find(&users)
	fmt.Println("============USERS====================")
	for _, u := range users {
		prinUser(u)
	}
	fmt.Println("================================")

}

func prinUser(u User) {
	fmt.Printf("ID: %d  Name:%s %s (%s) \n", u.ID, u.Name, u.LastName, u.NickName)
	fmt.Printf("User Files: \n")
	for _, m := range u.Files {
		fmt.Printf("\t Name %s format %s FID %d \n ", m.Name, m.Format, m.ID)
	}

}

func create3File(db *gorm.DB) {
	f1 := File{
		Name:   "Fname1",
		Md5:    "1fassdfsfdsfdsfdsfsd21313fsfds",
		Format: "flv",
	}

	f2 := File{
		Name:   "Fname2",
		Md5:    "1121212121s",
		Format: "mp4",
	}

	fx := File{
		Name: "fx",
		Md5:  "dfsgdsgdfgdf",
	}

	if err := db.Save(&f1).Error; err != nil {
		fmt.Printf("Got errors when save post %v", err)
	}

	if err := db.Save(&f2).Error; err != nil {
		fmt.Printf("Got errors when save post %v", err)
	}

	if err := db.Save(&fx).Error; err != nil {
		fmt.Printf("Got errors when save post %v", err)
	}
}

func createUser(db *gorm.DB, n string, nick string, c string) {
	db.Create(&User{Name: n, NickName: nick, Color: c})
}

func enterToContinue() {
	fmt.Println("Press the Enter ")
	var input string
	fmt.Scanln(&input)
}

func printFilesInDb(db *gorm.DB) {
	var files []File
	db.Find(&files)
	fmt.Println("============== ALL FILES==================")

	for _, f := range files {
		fmt.Printf("ID %d Name %s Format %s \n", f.ID, f.Name, f.Format)
	}
	fmt.Println("===============End All Files=================")

}
