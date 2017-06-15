package usermgmt

import (
	"upper.io/db.v3/lib/sqlbuilder"
)

type DataStore struct {
	DB sqlbuilder.Database
}

//Env ... Global environment
type Env struct {
	DataStore DataStore
}
