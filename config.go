package usermgmt

import (
	"github.com/sirupsen/logrus"
	"upper.io/db.v3/lib/sqlbuilder"
)

//Configuration ...
type Configuration struct {
	userDb    userDB
	appLogger appLogger
}

// NewConfiguration ...
func NewConfiguration(db sqlbuilder.Database, logger *logrus.Logger) *Configuration {

	usrDb := userDB{DB: db}
	appLog := appLogger{logger: logger}
	conf := Configuration{userDb: usrDb, appLogger: appLog}
	return &conf
}
