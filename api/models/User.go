package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int `gorm:"primary_key;auto_increament" json:"id"`
	Name string `gorm:"not null; unique" json:"name"`
	Email string `gorm:"not null; unique" json:"email"`
	Password string `gorm:"not null; unique" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error{
	hashedPassword, err:= Hash(u.Password)
	if err != nil{
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare(){
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error{
	switch strings.ToLower(action){
	case "update":
		if u.Name == ""{
			return errors.New("Required Name")
		}
		if u.Password == ""{
			return errors.New("Required Password")
		}
		if u.Email == ""{
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil{
			return errors.New("Invalid Email")
		}
		return nil

	case "login":
		if u.Password == ""{
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err !=nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == ""{
			return errors.New("Required Name")
		}
		if u.Password == ""{
			return errors.New("Required Password")
		}
		if u.Email == ""{
			return errors.New("Required Email")
		}
		if err :=checkmail.ValidateFormat(u.Email); err != nil{
			return errors.New("Invalid Email")
		}
		return nil
	}

}

func