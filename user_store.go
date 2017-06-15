package usermgmt

import "errors"

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

func (lcdb *lcDatabase) addUser(user *User) error {

	id, err := lcdb.DB.Collection("Users").Insert(user)
	if err != nil {
		return err
	}

	if i, ok := id.(int64); ok {
		user.ID = i
	} else {
		err := errors.New("Id not ok")
		return err
	}

	return nil
}

func (lcdb *lcDatabase) getUser(user User) (User, error) {
	return user, nil
}

func (lcdb *lcDatabase) deleteUser(user User) (User, error) {
	return user, nil
}

func (lcdb *lcDatabase) updateUser(user User) (User, error) {
	return user, nil
}
