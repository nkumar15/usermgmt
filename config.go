package usermgmt

import "upper.io/db.v3/lib/sqlbuilder"

//Configuration ... Global Configurationironment
type Configuration struct {
	userDb UserDB
	//logger *logrus.Logger
}

// NewConfigurationironment ...
func NewConfiguration(db sqlbuilder.Database) *Configuration {
	usrDb := UserDB{DB: db}
	conf := Configuration{userDb: usrDb}
	return &conf
}
