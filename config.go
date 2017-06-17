package usermgmt

import "upper.io/db.v3/lib/sqlbuilder"

//Env ... Global environment
type Env struct {
	userDb UserDB
	//logger *logrus.Logger
}

// NewEnvironment ...
func NewEnvironment(db sqlbuilder.Database) *Env {
	usrDb := UserDB{DB: db}
	env := Env{userDb: usrDb}
	return &env
}
