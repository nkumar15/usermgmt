package usermgmt

import (
	"log"
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

// Our ServeHTTP method is mostly the same, and also has the ability to
// access our *appContext's fields (templates, loggers, etc.) as well.
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Updated to pass ah.appContext as a parameter to our handler type.
	status, err := ah.Handler(ah.Conf, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			httpStatusNotFound(w, r, err)

			// And if we wanted a friendlier error page, we can
			// now leverage our context instance - e.g.
			// err := ah.renderTemplate(w, "http_404.tmpl", nil)
		case http.StatusInternalServerError:
			httpStatusInternalServerError(w, err)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}
