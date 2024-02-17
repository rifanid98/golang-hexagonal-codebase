package entity

import (
	"time"

	"codebase/core"
	"codebase/pkg/helper"
)

type Account struct {
	Id               string     `json:"id"`
	Otp              string     `json:"otp,omitempty"`
	Username         string     `json:"username"`
	Email            string     `json:"email"`
	Password         string     `json:"password"`
	VerifiedDate     time.Time  `json:"verified_date"`
	ResetToken       string     `json:"reset_token,omitempty"`
	ResetTokenExpire time.Time  `json:"reset_token_expire,omitempty"`
	SentAccess       *time.Time `json:"sent_access,omitempty"`
	Created          time.Time  `json:"created"`
	Modified         time.Time  `json:"modified"`
}

func (a *Account) SetPassword() *core.CustomError {
	if a.Password == "" {
		return nil
	}

	hashed, err := helper.HashPassword(a.Password)
	if err != nil {
		return &core.CustomError{
			Code:    core.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
	}

	a.Password = hashed
	return nil
}

func (a *Account) SetNewPassword(password string) *core.CustomError {
	hashed, err := helper.HashPassword(password)
	if err != nil {
		return &core.CustomError{
			Code:    core.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
	}

	a.Password = hashed
	return nil
}

func (a *Account) SetResetToken() *Account {
	hashed, _ := helper.HashPassword(a.Id)
	a.ResetToken = hashed
	a.ResetTokenExpire = time.Now().Add(time.Minute * 10)
	return a
}

func (a *Account) CheckPasword(password string) *core.CustomError {
	valid := helper.CheckPasswordHash(password, a.Password)
	if !valid {
		return &core.CustomError{
			Code:    core.WRONG_PASSWORD,
			Message: "wrong password",
		}
	}
	return nil
}

func (a *Account) RemoveResetToken() *Account {
	a.ResetToken = ""
	a.ResetTokenExpire = time.Time{}
	return a
}
