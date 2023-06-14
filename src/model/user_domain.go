package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"time"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string
	GetNationality() string
	GetCreatedAt() time.Time

	SetID(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(email, password, name string, age int8, nationality string, createdAt time.Time) UserDomainInterface {
	return &userDomain{
		email:       email,
		password:    password,
		name:        name,
		age:         age,
		nationality: nationality,
		createdAt:   createdAt,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(name string, age int8, nationality string) UserDomainInterface {
	return &userDomain{
		name:        name,
		age:         age,
		nationality: nationality,
	}
}

type userDomain struct {
	id          string
	email       string
	password    string
	name        string
	age         int8
	nationality string
	createdAt   time.Time
}

func (ud *userDomain) GetCreatedAt() time.Time {
	createdAt := time.Now()
	return createdAt
}

func (ud *userDomain) GetNationality() string {
	return ud.nationality
}

func (ud *userDomain) SetID(id string) {
	ud.id = id

}

func (ud *userDomain) GetID() string {
	return ud.id

}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))

}
