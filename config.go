package usermgmt

import (
	"net/http"

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

// AppHandler ...
type AppHandler struct {
	Conf    *Configuration
	Handler func(*Configuration, http.ResponseWriter, *http.Request) (int, error)
}

// ServeHTTP ...
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	status, err := ah.Handler(ah.Conf, w, r)
	requestID := r.Header.Get("X-Request-Id")
	if err != nil {
		ah.Conf.appLogger.logger.Debugf("HTTP %s %s %s %s %d %q", requestID, r.Host, r.URL.Path, r.Method, status, err)
		switch status {
		case http.StatusNotFound:
			httpStatusNotFound(w, r, err)
		case http.StatusInternalServerError:
			httpStatusInternalServerError(w, err)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
	ah.Conf.appLogger.logger.Debugf("HTTP %s %s %s %s %d", requestID, r.Host, r.URL.Path, r.Method, status)
}
