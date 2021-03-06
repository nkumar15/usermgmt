package usermgmt

import (
	"errors"

	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// User ...
type User struct {
	ID         int64  `db:"Id,omitempty"`
	Name       string `db:"Name"`
	Email      string `db:"Email"`
	Password   string `db:"Password"`
}

// UserStore ...
type UserStore interface {
	AddUser(*User) error
	GetUserByID(id int64) (*User, error)
	GetUsers() (*[]User, error)
	DeleteUserByID(id int64) error
	UpdateUser(user *User) error
}

// UserDB ...
type userDB struct {
	DB sqlbuilder.Database
}

// AddUser ...
func (userDb *userDB) AddUser(user *User) error {

	dbs := userDb.DB
	id, err := dbs.Collection("Users").Insert(user)
	if err != nil {
		return err
	}

	// Todo and check:
	//what if record inserted but interface type is not determined correctly
	// For eg here in sqlite3 it is int64, what if returned int for pg?
	// this code won't work in pg
	// How to deal with this situation?
	if i, ok := id.(int64); ok {
		user.ID = i
	} else {
		err := errors.New("Id not ok")
		return err
	}

	return nil
}

// GetUserByID ...
func (userDb *userDB) GetUserByID(id int64) (*User, error) {

	dbs := userDb.DB
	col := dbs.Collection("Users")
	res := col.Find(db.Cond{"Id": id})
	defer res.Close()

	user := new(User)
	err := res.One(user)

	if err == db.ErrNoMoreRows {
		return nil, errors.New("ErrNoMoreRows")
	}
	return user, nil
}

// GetUsers ...
func (userDb *userDB) GetUsers() (*[]User, error) {

	var users []User
	dbs := userDb.DB
	col := dbs.Collection("Users")
	res := col.Find()
	defer res.Close()

	err := res.All(&users)
	return &users, err
}

// DeleteUserByID ...
func (userDb *userDB) DeleteUserByID(id int64) error {

	dbs := userDb.DB
	col := dbs.Collection("Users")
	res := col.Find(db.Cond{"Id": id})
	defer res.Close()
	return res.Delete()
}

// UpdateUser ...
func (userDb *userDB) UpdateUser(user *User) error {

	dbs := userDb.DB
	col := dbs.Collection("Users")
	res := col.Find(db.Cond{"Id": user.ID})
	defer res.Close()
	presentUser := new(User)
	err := res.One(presentUser)
	if err != nil {
		return err
	}
	presentUser = user
	err = res.Update(presentUser)
	if err != nil {
		return err
	}
	return nil
}
