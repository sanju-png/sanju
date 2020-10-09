package main

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type vechicles struct {
	Name         string
	Code         string
	Price        uint
	Manufactured time.Time
	gorm.Model
}

func main() {
	db, err := gorm.Open(sqlite.Open("grooom.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&vechicles{})

	// db.Create(&vechicles{Name: "Maruti", Code: "535", Price: 100})

	// CREATED IN BATCH
	var users = []vechicles{{Name: "Suzuki"}, {Name: "Audi"}, {Name: "Maruti"}, {Code: "12w"}, {Code: "32"}}
	db.Create(&users)

	user := vechicles{Code: "Bca", Price: 18, Manufactured: time.Now()}
	db.Create(&user)

	// find product with code "12w"
	var car vechicles
	// db.First(&car, "code = ?", "12w")

	// SELECT * FROM car WHERE id = 1 AND UPDATE NAME AND CODE.
	db.First(&car)
	car.Name = "MarutichangedtoI10"
	car.Code = "100FG"
	db.Save(&car)

	//SELECT * FROM CAR WHERE id = 5 and update its name
	db.Find(&car, []int{3, 5}).Update("name", "newupdate")

	//

}

// Note: I went through groom document. I find that groom is just a model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt and we can change permission like read only,rw,etc
// I normally get the idea of groom and tried to create and update.I find comfortable to create records and batch record selected but (creating from sql expression :Is it necessary ?,if yes i will go through it.)
// IN updates i tried to update single column updates will try tomorrow more.so my question is should i follow creating from sql expression too.
