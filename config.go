package usermgmt

import (
	"github.com/sirupsen/logrus"
	"upper.io/db.v3/lib/sqlbuilder"
)

//Configuration ...
type Configuration struct {
	db        UserStore
	appLogger appLogger
}

// NewConfiguration ...
func NewConfiguration(db sqlbuilder.Database, logger *logrus.Logger) *Configuration {

	usrDb := &userDB{DB: db}
	appLog := appLogger{logger: logger}
	conf := Configuration{db: usrDb, appLogger: appLog}
	return &conf
}
