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
	Files    []*File `gorm:"many2many:user_file;"`
}

// `File` belongs to `User`, `UserID` is the foreign key
type File struct {
	gorm.Model
	Name   string
	Md5    string
	Format string `gorm:"default:'mp4'"`
	UserId uint
	Users  []*User `gorm:"many2many:user_file;"`
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &File{})

	db.Create(&User{Name: "med1", NickName: "meddy1"})
	db.Create(&User{Name: "med2"})
	db.Create(&User{Name: "med3"})

	var users []User
	db.Find(&users)
	fmt.Println("============USERS====================")
	for _, e := range users {
		prinUser(e)
	}
	fmt.Println("================================")

	var u1 User
	db.First(&u1, "Name = ?", "med1")

	var u2 User
	db.First(&u2, "Name = ?", "med2")

	var u3 User
	db.First(&u3, "Name = ?", "med3")

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
		Name:   "Very Important File",
		Md5:    "dfsgdsgdfgdf",
		Format: "mp4",
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

	var files []File
	db.Find(&files)
	fmt.Println("==============FILES==================")

	for _, e := range files {
		fmt.Println("---")
		fmt.Printf("Name %v format %s id %v userid %v UserName %s UserName %d", e.Name, e.Format, e.ID)
		fmt.Println("---")
	}
	fmt.Println("================================")

	db.First(&u1, "Name = ?", "med1")

	db.First(&u2, "Name = ?", "med2")

	db.First(&u3, "Name = ?", "med3")

	u2.Files = append(u2.Files, &f1, &f2)

	db.Save(&u1)
	db.Save(&u2)

	db.First(&u2, "Name = ?", "med2")
	prinUser(u2)

	getUserRelatedToFile(db, "Filename2")

	filesReslatedToUser(db, "med2")

}

func getUserRelatedToFile(db *gorm.DB, filename string) {
	fmt.Println("Getting Users Related File  ", filename)
	var f File
	var us []*User
	db.First(&f, "id = ?", 1)
	fmt.Println("/////////////////////")
	fmt.Printf("f %v", f.Name)
	fmt.Println("/////////////////////")
	db.Model(&f).Related(&us, "Files")
	fmt.Println(len(us))
	for _, e := range us {
		fmt.Println(e.ID, e.Name)
	}

}

func filesReslatedToUser(db *gorm.DB, name string) {
	fmt.Println("Getting File Related To User  ", name)

	var fs []*File
	var u User
	db.First(&u, "Name = ?", name)
	fmt.Println("/////////////////////")
	db.Model(&u).Related(&fs, "Files")
	fmt.Println(len(fs))
	for _, e := range fs {
		fmt.Println(e.ID, e.Name)
	}

}

func prinUser(u User) {
	fmt.Println("********************")
	fmt.Printf("ID: %d  Name:%s %s (%s) \n", u.ID, u.Name, u.LastName, u.NickName)
	for _, m := range u.Files {
		fmt.Printf("\t Name %s format %s fid %d md5 %s\n ", m.Name, m.Format, m.ID, m.Md5)
	}
	fmt.Println("********************")
}
