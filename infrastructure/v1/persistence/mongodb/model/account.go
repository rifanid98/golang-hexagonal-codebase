package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"codebase/core/v1/entity"
)

type Account struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	Email            string             `bson:"email,omitempty"`
	Username         string             `bson:"username,omitempty"`
	Password         string             `bson:"password,omitempty"`
	Otp              string             `bson:"otp,omitempty"`
	ResetToken       string             `bson:"reset_token"`
	ResetTokenExpire time.Time          `bson:"reset_token_expire"`
	SentAccess       *time.Time         `bson:"sent_access,omitempty"`
	VerifiedDate     time.Time          `bson:"verified_date"`
	Created          time.Time          `bson:"created,omitempty"`
	Modified         time.Time          `bson:"modified,omitempty"`
}

func (doc *Account) Bind(account *entity.Account) *Account {
	return &Account{
		Id:               GetObjectId(account.Id),
		Email:            account.Email,
		Username:         account.Username,
		Password:         account.Password,
		Otp:              account.Otp,
		ResetToken:       account.ResetToken,
		ResetTokenExpire: account.ResetTokenExpire,
		SentAccess:       account.SentAccess,
		VerifiedDate:     time.Time{},
		Created:          time.Time{},
		Modified:         account.Modified,
	}
}

func (doc *Account) Entity() *entity.Account {
	return &entity.Account{
		Id:               GetObjectIdHex(doc.Id),
		Otp:              doc.Otp,
		Username:         doc.Username,
		Email:            doc.Email,
		Password:         doc.Password,
		VerifiedDate:     doc.VerifiedDate,
		ResetToken:       doc.ResetToken,
		ResetTokenExpire: doc.ResetTokenExpire,
		SentAccess:       doc.SentAccess,
		Created:          doc.Created,
		Modified:         doc.Modified,
	}
}

type Accounts []Account

func (accs Accounts) Bind(accounts []entity.Account) Accounts {
	for i := range accounts {
		now := time.Now()
		acc := new(Account).Bind(&accounts[i])
		acc.Created = now
		acc.Modified = now
		if accounts[i].Id != "" {
			acc.Id = GetObjectId(accounts[i].Id)
		}
		accs = append(accs, *acc)
	}
	return accs
}

func (accs Accounts) Generics() []any {
	var data []any
	for i := range accs {
		data = append(data, accs[i])
	}
	return data
}

func (accs Accounts) Entities() []entity.Account {
	var es []entity.Account
	for i := range accs {
		es = append(es, *accs[i].Entity())
	}
	return es
}
