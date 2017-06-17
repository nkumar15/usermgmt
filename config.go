package usermgmt

import (
	"upper.io/db.v3/lib/sqlbuilder"
)

//Env ... Global environment
type Env struct {
	DB sqlbuilder.Database
}
