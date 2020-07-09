package main

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)
import _ "github.com/go-playground/validator/v10"

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/30 9:05
 */

type RegisterReq struct {
	Username       string `json:"username" validate:"gt=0"`
	PasswordNew    string `json:"password_new" validate:"gt=0"`
	PasswordRepeat string `json:"password_repeat" validate:"eqfield=PasswordNew"`
	Email          string `json:"email" validate:"email"`
}

func register(req RegisterReq) error {

	if len(req.Username) > 0 {
		if len(req.PasswordNew) > 0 && len(req.PasswordRepeat) > 0 {
			if req.PasswordNew == req.PasswordRepeat {
				if emailFormatValida(req.Email) {
					createUser()
					return nil
				} else {
					return errors.New("invalid email")
				}
			} else {
				return errors.New("password and reinput must be the same")
			}
		} else {
			return errors.New("password and password reinput must be longer than 0")
		}
	} else {
		return errors.New("length of username cannot be 0")
	}
}

func createUser() {
	fmt.Println("invoked in createUser method")
}

func emailFormatValida(s string) bool {

	return false
}
func registerRefactor(req RegisterReq) error {
	if len(req.Username) == 0 {
		return errors.New("length of username connot be 0")
	}
	if len(req.PasswordRepeat) == 0 || len(req.PasswordNew) == 0 {
		return errors.New("password and password reinput must be longer than 0")
	}
	if req.PasswordNew != req.PasswordRepeat {
		return errors.New("password and reinput must be the same")
	}
	if emailFormatValida(req.Email) {
		return errors.New("invalid email")
	}
	createUser()
	return nil
}
func main() {
	//create registerReq object
	var req = RegisterReq{
		Username:       "gisonwin",
		PasswordNew:    "abcd",
		PasswordRepeat: "abcd",
		Email:          "gisonwin@qq.com",
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err)
	}
	err = registerRefactor(req)
	if err != nil {
		fmt.Println(err)
	}
}
