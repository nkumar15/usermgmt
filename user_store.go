package usermgmt

import (
	"errors"

	db "upper.io/db.v3"
)

// User ...
type User struct {
	ID         int64  `db:"Id,omitempty"`
	Name       string `db:"Name"`
	GUID       string `db:"Guid"`
	Password   string `db:"Password"`
	Email      string `db:"Email"`
	Salt       string `db:"Salt"`
	JoinedDate string `db:"JoinedDate"`
}

func (ds *DataStore) addUser(user *User) error {

	id, err := ds.DB.Collection("Users").Insert(user)
	if err != nil {
		return err
	}

	// what if record inserted but interface type is not determined correctly
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

func (ds *DataStore) getUser(user *User) error {

	col := ds.DB.Collection("Users")
	res := col.Find(db.Cond{"Id": user.ID})
	defer res.Close()

	err := res.One(user)

	if err == db.ErrNoMoreRows {
		return errors.New("ErrNoMoreRows")
	}
	return nil
}

func (ds *DataStore) deleteUser(user *User) error {

	col := ds.DB.Collection("Users")
	res := col.Find(db.Cond{"Id": user.ID})
	defer res.Close()
	return res.Delete()
}

func (ds *DataStore) getUsers(users *[]User) error {

	col := ds.DB.Collection("Users")
	res := col.Find()
	defer res.Close()

	err := res.All(users)
	return err
}
