package models

import "time"

type User struct {
	ID 		int
	Name 		string          `gorm:"not null"`
	Email   	string		`gorm:"not null;unique"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time
}
func (u User) GetUsers() *[]User {
   db := GormDB()
  	if !db.HasTable(&User{}){
		db.CreateTable(&User{})
	}
   users := &[]User{}
   db.Find(users)
   return users
}

func (u User) PostUser(name string, email string){
	db := GormDB()

	db.Create(&User{Name:name, Email:email})

	return
}

func(u User) DeleteUser(id int){
	db := GormDB()
	user := User{}

	db.First(&user,id)
	db.Delete(&user)

	return
}