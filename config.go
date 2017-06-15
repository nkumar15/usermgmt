package usermgmt

import (
	"upper.io/db.v3/lib/sqlbuilder"
)

type lcDatabase struct {
	DB sqlbuilder.Database
}

//Env ... Global environment
type Env struct {
	Database lcDatabase
}
