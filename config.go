package usermgmt

import (
	"github.com/sirupsen/logrus"
	"upper.io/db.v3/lib/sqlbuilder"
)

//Configuration ...
type Configuration struct {
	userDb    UserDB
	appLogger AppLogger
}

// NewConfiguration ...
func NewConfiguration(db sqlbuilder.Database, logger *logrus.Logger) *Configuration {

	usrDb := UserDB{DB: db}
	appLog := AppLogger{logger: logger}
	conf := Configuration{userDb: usrDb, appLogger: appLog}
	return &conf
}
